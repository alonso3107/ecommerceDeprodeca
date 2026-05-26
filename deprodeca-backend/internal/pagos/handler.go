package pagos

import (
	"strconv"

	"deprodeca-backend/internal/gamificacion"
	"deprodeca-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// Handler maneja las peticiones HTTP de pagos.
type Handler struct {
	service          *Service
	gamificacionSvc  *gamificacion.Service
}

func NewHandler(service *Service, gamificacionSvc *gamificacion.Service) *Handler {
	return &Handler{service: service, gamificacionSvc: gamificacionSvc}
}

// RegistrarPago maneja POST /api/v1/pagos
func (h *Handler) RegistrarPago(c *fiber.Ctx) error {
	var input RegistrarPagoInput

	if err := c.BodyParser(&input); err != nil {
		return response.BadRequest(c, "Cuerpo de la petición inválido")
	}

	pago, err := h.service.RegistrarPago(c.Context(), input)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, pago, "Comprobante de pago registrado")
}

// ListarMisPagos maneja GET /api/v1/pagos
func (h *Handler) ListarMisPagos(c *fiber.Ctx) error {
	usuarioID := c.Locals("usuarioID").(int64)
	pagina, _ := strconv.Atoi(c.Query("pagina", "1"))
	limite, _ := strconv.Atoi(c.Query("limite", "20"))

	pagos, err := h.service.ListarMisPagos(c.Context(), usuarioID, pagina, limite)
	if err != nil {
		return response.Internal(c, "Error al listar pagos")
	}

	return response.Ok(c, pagos, "Pagos obtenidos")
}

// VerificarPago maneja PUT /api/v1/pagos/:id/verificar (solo admin)
// Confirma el pago, actualiza el pedido a "confirmado" y dispara gamificación.
func (h *Handler) VerificarPago(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "ID de pago inválido")
	}

	pago, err := h.service.VerificarPago(c.Context(), id, h.gamificacionSvc)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Ok(c, pago, "Pago verificado. Puntos de gamificación acreditados.")
}
