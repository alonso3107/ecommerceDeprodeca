package usuarios

import (
	"context"
	"errors"
	"strings"
)

// Service contiene la lógica de negocio para gestión de usuarios.
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// ObtenerPerfil devuelve el perfil del usuario por su ID.
func (s *Service) ObtenerPerfil(ctx context.Context, usuarioID int64) (Usuario, error) {
	return s.repo.ObtenerPorID(ctx, usuarioID)
}

// ActualizarPerfil actualiza los datos editables del perfil.
func (s *Service) ActualizarPerfil(ctx context.Context, usuarioID int64, input ActualizarPerfilInput) error {
	if strings.TrimSpace(input.Nombre) == "" {
		return errors.New("nombre es obligatorio")
	}
	if strings.TrimSpace(input.Empresa) == "" {
		return errors.New("empresa es obligatorio")
	}

	return s.repo.Actualizar(ctx, usuarioID, input)
}
