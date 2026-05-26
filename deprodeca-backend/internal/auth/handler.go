package auth

import (
	"deprodeca-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// Handler maneja las peticiones HTTP de autenticación.
// Solo recibe HTTP, valida input, llama al service, devuelve JSON (SRP de SOLID).
type Handler struct {
	service *Service
}

// NewHandler crea el handler de auth con su servicio inyectado.
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Registro maneja POST /api/v1/auth/registro
func (h *Handler) Registro(c *fiber.Ctx) error {
	var input RegistroInput

	if err := c.BodyParser(&input); err != nil {
		return response.BadRequest(c, "Cuerpo de la petición inválido")
	}

	resp, err := h.service.Registrar(c.Context(), input)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, resp, "Cuenta creada exitosamente")
}

// Login maneja POST /api/v1/auth/login
func (h *Handler) Login(c *fiber.Ctx) error {
	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return response.BadRequest(c, "Cuerpo de la petición inválido")
	}

	resp, err := h.service.Login(c.Context(), input)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Ok(c, resp, "Inicio de sesión exitoso")
}
