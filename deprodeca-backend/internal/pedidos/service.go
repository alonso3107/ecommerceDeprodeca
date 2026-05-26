package pedidos

import (
	"context"
	"errors"
	"fmt"

	"deprodeca-backend/internal/gamificacion"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Service contiene la lógica de negocio para pedidos.
type Service struct {
	repo Repository
	db   *pgxpool.Pool // Necesario para helpers de stock (transacciones cross-domain)
}

func NewService(repo Repository, db *pgxpool.Pool) *Service {
	return &Service{repo: repo, db: db}
}

// CrearPedido crea un pedido con sus ítems. Valida stock, descuenta, calcula total.
func (s *Service) CrearPedido(ctx context.Context, usuarioID int64, input CrearPedidoInput) (Pedido, error) {
	if len(input.Items) == 0 {
		return Pedido{}, errors.New("el pedido debe tener al menos un ítem")
	}

	// Validar stock y calcular total
	var total float64
	type itemValidado struct {
		productoID int64
		cantidad   int
		precio     float64
		nombre     string
	}
	var items []itemValidado

	for _, it := range input.Items {
		if it.Cantidad <= 0 {
			return Pedido{}, errors.New("la cantidad debe ser mayor a 0")
		}

		precio, stock, nombre, err := ObtenerProductoPrecioYStock(ctx, s.db, it.ProductoID)
		if err != nil {
			return Pedido{}, errors.New("producto no encontrado: " + nombre)
		}
		if stock < it.Cantidad {
			return Pedido{}, errors.New("stock insuficiente para: " + nombre)
		}

		items = append(items, itemValidado{
			productoID: it.ProductoID,
			cantidad:   it.Cantidad,
			precio:     precio,
			nombre:     nombre,
		})
		total += precio * float64(it.Cantidad)
	}

	// Crear pedido
	pedidoID, err := s.repo.CrearPedido(ctx, usuarioID, total)
	if err != nil {
		return Pedido{}, err
	}

	// Insertar detalles y descontar stock
	for _, it := range items {
		subtotal := it.precio * float64(it.cantidad)
		if err := s.repo.AgregarDetalle(ctx, pedidoID, it.productoID, it.cantidad, it.precio, subtotal); err != nil {
			return Pedido{}, err
		}
		if err := DescontarStock(ctx, s.db, it.productoID, it.cantidad); err != nil {
			return Pedido{}, err
		}
	}

	// Obtener pedido completo
	return s.ObtenerPedido(ctx, pedidoID)
}

// ObtenerPedido obtiene un pedido con sus detalles.
func (s *Service) ObtenerPedido(ctx context.Context, pedidoID int64) (Pedido, error) {
	pedido, err := s.repo.ObtenerPedidoPorID(ctx, pedidoID)
	if err != nil {
		return Pedido{}, err
	}

	detalles, err := s.repo.ListarDetallesPorPedido(ctx, pedidoID)
	if err != nil {
		return pedido, err
	}
	pedido.Detalles = detalles

	return pedido, nil
}

// ListarMisPedidos lista los pedidos del usuario autenticado.
func (s *Service) ListarMisPedidos(ctx context.Context, usuarioID int64, pagina, limite int) ([]Pedido, error) {
	if pagina < 1 {
		pagina = 1
	}
	if limite < 1 || limite > 50 {
		limite = 20
	}

	return s.repo.ListarPedidosPorUsuario(ctx, usuarioID, pagina, limite)
}

// ListarTodosPedidos lista todos los pedidos (solo admin).
func (s *Service) ListarTodosPedidos(ctx context.Context, pagina, limite int) ([]Pedido, error) {
	if pagina < 1 {
		pagina = 1
	}
	if limite < 1 || limite > 100 {
		limite = 50
	}
	return s.repo.ListarTodos(ctx, pagina, limite)
}

// ActualizarEstado cambia el estado de un pedido (solo admin).
func (s *Service) ActualizarEstado(ctx context.Context, pedidoID int64, estado string) error {
	estadosValidos := map[string]bool{
		"pendiente": true, "confirmado": true, "enviado": true,
		"entregado": true, "cancelado": true,
	}
	if !estadosValidos[estado] {
		return fmt.Errorf("estado inválido: %s", estado)
	}

	return s.repo.ActualizarEstado(ctx, pedidoID, estado)
}

// EnviarPedido marca un pedido como "enviado" (solo admin).
func (s *Service) EnviarPedido(ctx context.Context, pedidoID int64) (Pedido, error) {
	pedido, err := s.ObtenerPedido(ctx, pedidoID)
	if err != nil {
		return Pedido{}, err
	}
	if pedido.Estado != "confirmado" {
		return Pedido{}, fmt.Errorf("solo pedidos confirmados pueden enviarse (estado actual: %s)", pedido.Estado)
	}

	if err := s.repo.ActualizarEstado(ctx, pedidoID, "enviado"); err != nil {
		return Pedido{}, err
	}

	return s.ObtenerPedido(ctx, pedidoID)
}

// EntregarPedido marca un pedido como "entregado" y dispara la gamificación (solo admin).
func (s *Service) EntregarPedido(ctx context.Context, pedidoID int64, gs *gamificacion.Service) (Pedido, error) {
	pedido, err := s.ObtenerPedido(ctx, pedidoID)
	if err != nil {
		return Pedido{}, err
	}
	if pedido.Estado != "enviado" {
		return Pedido{}, fmt.Errorf("solo pedidos enviados pueden entregarse (estado actual: %s)", pedido.Estado)
	}

	if err := s.repo.ActualizarEstado(ctx, pedidoID, "entregado"); err != nil {
		return Pedido{}, err
	}

	// ─── Gamificación: acumular puntos al entregar ───
	if _, _, err := gs.AcumularPuntos(ctx, pedido.UsuarioID, pedido.Total); err != nil {
		return Pedido{}, fmt.Errorf("pedido entregado pero error al acumular puntos: %w", err)
	}

	return s.ObtenerPedido(ctx, pedidoID)
}
