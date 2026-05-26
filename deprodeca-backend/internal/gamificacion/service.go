package gamificacion

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

// Service contiene la lógica de negocio de gamificación.
type Service struct {
	repo Repository
}

// NewService crea el servicio con su repositorio inyectado.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// ObtenerRango devuelve el estado de gamificación del usuario.
// Si no existe, lo crea automáticamente (primer acceso).
func (s *Service) ObtenerRango(ctx context.Context, usuarioID int64) (Gamificacion, error) {
	g, err := s.repo.ObtenerPorUsuario(ctx, usuarioID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return s.repo.Crear(ctx, usuarioID)
		}
		return Gamificacion{}, err
	}
	return g, nil
}

// AcumularPuntos suma puntos y volumen al usuario tras una compra.
// Recalcula el rango automáticamente según el nuevo volumen acumulado.
// Devuelve el estado actualizado y si hubo cambio de rango.
func (s *Service) AcumularPuntos(ctx context.Context, usuarioID int64, montoCompra float64) (Gamificacion, bool, error) {
	puntosGanados := CalcularPuntos(montoCompra)

	// Obtener estado actual
	actual, err := s.ObtenerRango(ctx, usuarioID)
	if err != nil {
		return Gamificacion{}, false, err
	}

	nuevoVolumen := actual.VolumenTotal + montoCompra
	rangoRecalculado := CalcularRango(nuevoVolumen)

	if err := s.repo.ActualizarPuntosYRango(ctx, usuarioID, puntosGanados, montoCompra, rangoRecalculado); err != nil {
		return Gamificacion{}, false, err
	}

	subioRango := RangoSuperior(actual.Rango, rangoRecalculado)
	if subioRango {
		log.Printf("[GAMIFICACIÓN] ¡Usuario #%d subió de %s a %s! (volumen: S/ %.2f)",
			usuarioID, NombreRango(actual.Rango), NombreRango(rangoRecalculado), nuevoVolumen)
	}

	// Obtener estado actualizado
	nuevo, err := s.repo.ObtenerPorUsuario(ctx, usuarioID)
	if err != nil {
		return Gamificacion{}, false, err
	}

	return nuevo, subioRango, nil
}

// Leaderboard devuelve el top N del ranking.
func (s *Service) Leaderboard(ctx context.Context, limite int) ([]LeaderboardItem, error) {
	if limite < 1 || limite > 100 {
		limite = 20
	}
	return s.repo.ObtenerLeaderboard(ctx, limite)
}
