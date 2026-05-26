package gamificacion

import (
	"strconv"

	"deprodeca-backend/pkg/response"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// Handler maneja las peticiones HTTP de gamificación.
// Solo recibe HTTP, valida input, llama al service, devuelve JSON (SRP).
type Handler struct {
	service *Service
	wsHub   *HubWebSocket
}

// NewHandler crea el handler con sus dependencias inyectadas.
func NewHandler(service *Service, wsHub *HubWebSocket) *Handler {
	return &Handler{
		service: service,
		wsHub:   wsHub,
	}
}

// ObtenerRango maneja GET /api/v1/gamificacion/rango
func (h *Handler) ObtenerRango(c *fiber.Ctx) error {
	usuarioID := c.Locals("usuarioID").(int64)

	g, err := h.service.ObtenerRango(c.Context(), usuarioID)
	if err != nil {
		return response.Internal(c, "Error al obtener rango")
	}

	return response.Ok(c, g, "Rango obtenido")
}

// Leaderboard maneja GET /api/v1/gamificacion/leaderboard
func (h *Handler) Leaderboard(c *fiber.Ctx) error {
	limite, _ := strconv.Atoi(c.Query("limite", "20"))

	items, err := h.service.Leaderboard(c.Context(), limite)
	if err != nil {
		return response.Internal(c, "Error al obtener leaderboard")
	}

	return response.Ok(c, items, "Leaderboard obtenido")
}

// WebSocketRangos delega al HubWebSocket para manejar la conexión WS.
func (h *Handler) WebSocketRangos(c *websocket.Conn) {
	h.wsHub.ManejarConexion(c)
}

// NotificarCambioRango notifica a un usuario específico si está conectado.
func (h *Handler) NotificarCambioRango(usuarioID int64, g Gamificacion) {
	h.wsHub.NotificarCambioRango(usuarioID, g)
}

// NotificarLeaderboardActualizado notifica a todos los conectados.
func (h *Handler) NotificarLeaderboardActualizado(leaderboard any) {
	h.wsHub.NotificarLeaderboardActualizado(leaderboard)
}
