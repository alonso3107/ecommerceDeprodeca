package contacto

import (
	"deprodeca-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Crear(c *fiber.Ctx) error {
	var input ContactoInput

	if err := c.BodyParser(&input); err != nil {
		return response.BadRequest(c, "Cuerpo de la petición inválido")
	}

	contacto, err := h.service.Crear(c.Context(), input)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, contacto, "Consulta enviada correctamente")
}
