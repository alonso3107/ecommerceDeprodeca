package pagos

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Service contiene la lógica de negocio para pagos.
type Service struct {
	repo Repository
	db   *pgxpool.Pool
}

// NewService crea el servicio de pagos.
func NewService(repo Repository, db *pgxpool.Pool) *Service {
	return &Service{repo: repo, db: db}
}

// RegistrarPago valida y guarda un comprobante de pago.
// Permite comprobante_url vacío (checkout simplificado).
func (s *Service) RegistrarPago(ctx context.Context, input RegistrarPagoInput) (Pago, error) {
	if input.PedidoID <= 0 {
		return Pago{}, errors.New("pedido_id es obligatorio")
	}
	if input.Monto <= 0 {
		return Pago{}, errors.New("monto debe ser mayor a 0")
	}
	if input.Metodo == "" {
		return Pago{}, errors.New("método de pago es obligatorio")
	}

	metodosValidos := map[string]bool{
		"transferencia": true,
		"yape":          true,
		"plin":          true,
		"efectivo":      true,
	}
	if !metodosValidos[input.Metodo] {
		return Pago{}, errors.New("método de pago inválido")
	}

	return s.repo.Crear(ctx, input)
}

// VerificarPago confirma el pago y actualiza el pedido a "confirmado".
// Solo el admin puede llamar a este método.
// La gamificación (puntos/rango) se dispara al ENTREGAR el pedido, no al pagar.
func (s *Service) VerificarPago(ctx context.Context, pagoID int64) (Pago, error) {
	// 1. Obtener el pago
	pago, err := s.repo.ObtenerPorID(ctx, pagoID)
	if err != nil {
		return Pago{}, fmt.Errorf("pago no encontrado: %w", err)
	}

	if pago.Estado == "verificado" {
		return Pago{}, errors.New("el pago ya fue verificado")
	}

	// 2. Actualizar estado del pago
	if err := s.repo.ActualizarEstado(ctx, pagoID, "verificado"); err != nil {
		return Pago{}, fmt.Errorf("error al verificar pago: %w", err)
	}

	// 3. Actualizar estado del pedido a "confirmado"
	if _, err := s.db.Exec(ctx,
		`UPDATE pedidos SET estado = 'confirmado', updated_at = NOW() WHERE id = $1`,
		pago.PedidoID,
	); err != nil {
		return Pago{}, fmt.Errorf("error al confirmar pedido: %w", err)
	}

	pago.Estado = "verificado"
	return pago, nil
}

// ListarMisPagos lista los pagos del usuario autenticado.
func (s *Service) ListarMisPagos(ctx context.Context, usuarioID int64, pagina, limite int) ([]Pago, error) {
	if pagina < 1 {
		pagina = 1
	}
	if limite < 1 || limite > 50 {
		limite = 20
	}

	return s.repo.ListarPorUsuario(ctx, usuarioID, pagina, limite)
}

// ListarTodosPagos lista todos los pagos (solo admin).
func (s *Service) ListarTodosPagos(ctx context.Context, pagina, limite int) ([]Pago, error) {
	if pagina < 1 {
		pagina = 1
	}
	if limite < 1 || limite > 100 {
		limite = 50
	}
	return s.repo.ListarTodos(ctx, pagina, limite)
}
