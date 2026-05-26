package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"deprodeca-backend/config"
	"deprodeca-backend/internal/auth"
	"deprodeca-backend/internal/gamificacion"
	"deprodeca-backend/internal/pagos"
	"deprodeca-backend/internal/pedidos"
	"deprodeca-backend/internal/productos"
	"deprodeca-backend/internal/usuarios"
	"deprodeca-backend/middleware"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func main() {
	// ─── 1. Cargar configuración ──────────────────────────
	cfg := config.Cargar()

	// ─── 2. Conectar PostgreSQL ────────────────────────────
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, cfg.Database.URL)
	if err != nil {
		log.Fatalf("[DB] No se pudo conectar a PostgreSQL: %v", err)
	}
	defer dbPool.Close()

	if err := dbPool.Ping(ctx); err != nil {
		log.Fatalf("[DB] Ping fallido: %v", err)
	}
	log.Println("[DB] PostgreSQL conectado — deprodeca")

	// ─── 3. Conectar Redis ─────────────────────────────────
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Printf("[REDIS] Advertencia: Redis no disponible (%v) — continuando sin caché", err)
	} else {
		log.Println("[REDIS] Conectado")
	}

	// ─── 4. Inicializar Fiber ──────────────────────────────
	app := fiber.New(fiber.Config{
		AppName:      "DEPRODECA API",
		ErrorHandler: errorHandler,
	})

	// ─── 5. Middlewares globales ───────────────────────────
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

	// ─── 6. Inyección de dependencias (SOLID: interfaces → servicios → handlers) ──
	// Auth
	authRepo := auth.NewRepository(dbPool)
	authService := auth.NewService(authRepo, cfg.JWT.Secret, cfg.JWT.Expiration)
	authHandler := auth.NewHandler(authService)

	// Usuarios
	usuariosRepo := usuarios.NewRepository(dbPool)
	usuariosService := usuarios.NewService(usuariosRepo)
	usuariosHandler := usuarios.NewHandler(usuariosService)

	// Productos + Categorías
	productosRepo := productos.NewRepository(dbPool)
	productosService := productos.NewService(productosRepo)
	productosHandler := productos.NewHandler(productosService)

	// Pedidos
	pedidosRepo := pedidos.NewRepository(dbPool)
	pedidosService := pedidos.NewService(pedidosRepo, dbPool)
	pedidosHandler := pedidos.NewHandler(pedidosService)

	// Gamificación
	gamificacionRepo := gamificacion.NewRepository(dbPool)
	gamificacionService := gamificacion.NewService(gamificacionRepo)
	gamificacionWSHub := gamificacion.NewHubWebSocket(gamificacionService, cfg.JWT.Secret)
	gamificacionHandler := gamificacion.NewHandler(gamificacionService, gamificacionWSHub)

	// Pagos
	pagosRepo := pagos.NewRepository(dbPool)
	pagosService := pagos.NewService(pagosRepo, dbPool)
	pagosHandler := pagos.NewHandler(pagosService, gamificacionService)

	// ─── 7. Rutas ──────────────────────────────────────────
	api := app.Group("/api/v1")

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "DEPRODECA API v1 — operativa",
			"data": fiber.Map{
				"env": cfg.Server.Env,
			},
		})
	})

	// ── Auth (público) ─────────────────────────────────────
	authGroup := api.Group("/auth")
	authGroup.Post("/registro", authHandler.Registro)
	authGroup.Post("/login", authHandler.Login)

	// ── Usuarios (protegido) ───────────────────────────────
	usuariosGroup := api.Group("/usuarios", middleware.JWTAuth(cfg.JWT.Secret))
	usuariosGroup.Get("/perfil", usuariosHandler.Perfil)
	usuariosGroup.Put("/perfil", usuariosHandler.ActualizarPerfil)

	// ── Categorías (público) ───────────────────────────────
	api.Get("/categorias", productosHandler.ListarCategorias)
	api.Get("/categorias/:slug", productosHandler.ObtenerCategoria)

	// ── Productos (público) ────────────────────────────────
	api.Get("/productos", productosHandler.ListarProductos)
	api.Get("/productos/buscar", productosHandler.BuscarProductos)
	api.Get("/productos/:slug", productosHandler.ObtenerProducto)

	// ── Pedidos (protegido) ────────────────────────────────
	pedidosGroup := api.Group("/pedidos", middleware.JWTAuth(cfg.JWT.Secret))
	pedidosGroup.Post("/", pedidosHandler.Crear)
	pedidosGroup.Get("/", pedidosHandler.ListarMisPedidos)
	pedidosGroup.Get("/:id", pedidosHandler.Obtener)

	// ── Pagos (protegido) ──────────────────────────────────
	pagosGroup := api.Group("/pagos", middleware.JWTAuth(cfg.JWT.Secret))
	pagosGroup.Post("/", pagosHandler.RegistrarPago)
	pagosGroup.Get("/", pagosHandler.ListarMisPagos)

	// Admin: verificar pago → dispara gamificación
	adminPagos := api.Group("/pagos", middleware.JWTAuth(cfg.JWT.Secret), middleware.RequireRol("admin"))
	adminPagos.Put("/:id/verificar", pagosHandler.VerificarPago)

	// ── Gamificación (protegido) ───────────────────────────
	gamificacionGroup := api.Group("/gamificacion", middleware.JWTAuth(cfg.JWT.Secret))
	gamificacionGroup.Get("/rango", gamificacionHandler.ObtenerRango)
	gamificacionGroup.Get("/leaderboard", gamificacionHandler.Leaderboard)

	// ── WebSocket Gamificación ─────────────────────────────
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws/rangos", websocket.New(gamificacionHandler.WebSocketRangos))

	// ─── 7. Rutas no encontradas ──────────────────────────
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Ruta no encontrada",
			"data":    nil,
		})
	})

	// ─── 8. Iniciar servidor con graceful shutdown ────────
	go func() {
		addr := ":" + os.Getenv("SERVER_PORT")
		log.Printf("[SERVER] DEPRODECA API iniciada en http://localhost%s", addr)
		if err := app.Listen(addr); err != nil {
			log.Fatalf("[SERVER] Error al iniciar: %v", err)
		}
	}()

	// ─── 9. Esperar señal de apagado ──────────────────────
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[SERVER] Apagando servidor...")
	_ = rdb.Close()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Fatalf("[SERVER] Error al apagar: %v", err)
	}
	log.Println("[SERVER] Servidor detenido correctamente")
}

// errorHandler centralizado para Fiber.
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"data":    nil,
		"message": err.Error(),
	})
}
