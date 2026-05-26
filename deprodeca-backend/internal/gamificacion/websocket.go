package gamificacion

import (
	"context"
	"sync"

	"deprodeca-backend/pkg/jwt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// HubWebSocket gestiona las conexiones WebSocket activas por usuarioID.
// Separado del Handler HTTP para cumplir SRP (Single Responsibility Principle).
type HubWebSocket struct {
	mu       sync.RWMutex
	conns    map[int64]*websocket.Conn
	jwtSecret string
	service   *Service
}

// NewHubWebSocket crea un nuevo hub de WebSocket.
func NewHubWebSocket(service *Service, jwtSecret string) *HubWebSocket {
	return &HubWebSocket{
		conns:     make(map[int64]*websocket.Conn),
		jwtSecret: jwtSecret,
		service:   service,
	}
}

// ManejarConexion es el handler para WS /ws/rangos.
// El cliente se conecta con un token JWT como query param: ?token=xxx
func (h *HubWebSocket) ManejarConexion(c *websocket.Conn) {
	// Autenticar vía query param token
	tokenStr := c.Query("token")
	if tokenStr == "" {
		c.WriteJSON(fiber.Map{"error": "token requerido"})
		c.Close()
		return
	}

	claims, err := jwt.ValidarToken(tokenStr, h.jwtSecret)
	if err != nil {
		c.WriteJSON(fiber.Map{"error": "token inválido o expirado"})
		c.Close()
		return
	}

	usuarioID := claims.UsuarioID

	// Registrar conexión
	h.mu.Lock()
	h.conns[usuarioID] = c
	h.mu.Unlock()

	defer func() {
		h.mu.Lock()
		delete(h.conns, usuarioID)
		h.mu.Unlock()
		c.Close()
	}()

	// Enviar estado actual al conectar
	g, err := h.service.ObtenerRango(context.Background(), usuarioID)
	if err == nil {
		c.WriteJSON(fiber.Map{
			"type": "estado_actual",
			"data": g,
		})
	}

	// Mantener conexión viva (ping/pong)
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}
}

// NotificarCambioRango envía un mensaje WebSocket al usuario si está conectado.
func (h *HubWebSocket) NotificarCambioRango(usuarioID int64, g Gamificacion) {
	h.mu.RLock()
	conn, ok := h.conns[usuarioID]
	h.mu.RUnlock()

	if ok {
		conn.WriteJSON(fiber.Map{
			"type": "cambio_rango",
			"data": g,
		})
	}
}

// NotificarLeaderboardActualizado envía el leaderboard a todos los conectados.
func (h *HubWebSocket) NotificarLeaderboardActualizado(leaderboard any) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, conn := range h.conns {
		conn.WriteJSON(fiber.Map{
			"type": "leaderboard",
			"data": leaderboard,
		})
	}
}
