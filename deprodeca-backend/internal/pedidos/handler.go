package pedidos

import (
	"strconv"

	"deprodeca-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// Handler maneja las peticiones HTTP de pedidos.
type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Crear maneja POST /api/v1/pedidos
func (h *Handler) Crear(c *fiber.Ctx) error {
	usuarioID := c.Locals("usuarioID").(int64)

	var input CrearPedidoInput
	if err := c.BodyParser(&input); err != nil {
		return response.BadRequest(c, "Cuerpo de la petición inválido")
	}

	pedido, err := h.service.CrearPedido(c.Context(), usuarioID, input)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}

	return response.Created(c, pedido, "Pedido creado exitosamente")
}

// ListarMisPedidos maneja GET /api/v1/pedidos
func (h *Handler) ListarMisPedidos(c *fiber.Ctx) error {
	usuarioID := c.Locals("usuarioID").(int64)
	pagina, _ := strconv.Atoi(c.Query("pagina", "1"))
	limite, _ := strconv.Atoi(c.Query("limite", "20"))

	pedidos, err := h.service.ListarMisPedidos(c.Context(), usuarioID, pagina, limite)
	if err != nil {
		return response.Internal(c, "Error al listar pedidos")
	}

	return response.Ok(c, pedidos, "Pedidos obtenidos")
}

// Obtener maneja GET /api/v1/pedidos/:id
func (h *Handler) Obtener(c *fiber.Ctx) error {
	usuarioID := c.Locals("usuarioID").(int64)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "ID de pedido inválido")
	}

	pedido, err := h.service.ObtenerPedido(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Pedido no encontrado")
	}

	// Solo el dueño o admin pueden ver
	if pedido.UsuarioID != usuarioID {
		rol := c.Locals("rol").(string)
		if rol != "admin" {
			return response.Forbidden(c, "No tienes permiso para ver este pedido")
		}
	}

	return response.Ok(c, pedido, "Pedido obtenido")
}
