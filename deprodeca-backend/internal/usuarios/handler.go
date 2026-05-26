package usuarios

import (
	"deprodeca-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// Handler maneja las peticiones HTTP de usuarios.
type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Perfil maneja GET /api/v1/usuarios/perfil
func (h *Handler) Perfil(c *fiber.Ctx) error {
	usuarioID := c.Locals("usuarioID").(int64)

	usuario, err := h.service.ObtenerPerfil(c.Context(), usuarioID)
	if err != nil {
		return response.NotFound(c, "Usuario no encontrado")
	}

	return response.Ok(c, usuario, "Perfil obtenido")
}

// ActualizarPerfil maneja PUT /api/v1/usuarios/perfil
func (h *Handler) ActualizarPerfil(c *fiber.Ctx) error {
	usuarioID := c.Locals("usuarioID").(int64)

	var input ActualizarPerfilInput
	if err := c.BodyParser(&input); err != nil {
		return response.BadRequest(c, "Cuerpo de la petición inválido")
	}

	if err := h.service.ActualizarPerfil(c.Context(), usuarioID, input); err != nil {
		return response.BadRequest(c, err.Error())
	}

	// Devolver perfil actualizado
	usuario, err := h.service.ObtenerPerfil(c.Context(), usuarioID)
	if err != nil {
		return response.Ok(c, nil, "Perfil actualizado")
	}

	return response.Ok(c, usuario, "Perfil actualizado")
}
