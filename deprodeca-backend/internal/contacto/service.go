package contacto

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Crear(ctx context.Context, input ContactoInput) (Contacto, error) {
	input.Nombre = strings.TrimSpace(input.Nombre)
	input.Email = strings.ToLower(strings.TrimSpace(input.Email))
	input.Telefono = strings.TrimSpace(input.Telefono)
	input.Motivo = strings.TrimSpace(input.Motivo)
	input.Mensaje = strings.TrimSpace(input.Mensaje)

	if err := validarContacto(input); err != nil {
		return Contacto{}, err
	}

	if input.Motivo == "" {
		input.Motivo = "No especificado"
	}

	contacto, err := s.repo.CrearContacto(ctx, input)
	if err != nil {
		return Contacto{}, fmt.Errorf("error al guardar la consulta: %w", err)
	}

	return contacto, nil
}

func validarContacto(input ContactoInput) error {
	if input.Nombre == "" {
		return errors.New("nombre es obligatorio")
	}
	if input.Email == "" {
		return errors.New("email es obligatorio")
	}
	if _, err := mail.ParseAddress(input.Email); err != nil {
		return errors.New("email inválido")
	}
	if input.Mensaje == "" {
		return errors.New("mensaje es obligatorio")
	}
	return nil
}
