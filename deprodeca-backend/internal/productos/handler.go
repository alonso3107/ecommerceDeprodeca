package productos

import (
	"strconv"

	"deprodeca-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// Handler maneja las peticiones HTTP del catálogo.
type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// ListarProductos maneja GET /api/v1/productos
func (h *Handler) ListarProductos(c *fiber.Ctx) error {
	pagina, _ := strconv.Atoi(c.Query("pagina", "1"))
	limite, _ := strconv.Atoi(c.Query("limite", "20"))
	categoriaID, _ := strconv.ParseInt(c.Query("categoria_id", "0"), 10, 64)
	busqueda := c.Query("q", "")

	productos, err := h.service.ListarProductos(c.Context(), pagina, limite, categoriaID, busqueda)
	if err != nil {
		return response.Internal(c, "Error al listar productos")
	}

	return response.Ok(c, productos, "Productos obtenidos")
}

// ObtenerProducto maneja GET /api/v1/productos/:slug
func (h *Handler) ObtenerProducto(c *fiber.Ctx) error {
	slug := c.Params("slug")

	producto, err := h.service.ObtenerProductoPorSlug(c.Context(), slug)
	if err != nil {
		return response.NotFound(c, "Producto no encontrado")
	}

	return response.Ok(c, producto, "Producto obtenido")
}

// BuscarProductos maneja GET /api/v1/productos/buscar?q=
func (h *Handler) BuscarProductos(c *fiber.Ctx) error {
	query := c.Query("q", "")
	if query == "" {
		return response.BadRequest(c, "Parámetro de búsqueda 'q' requerido")
	}

	pagina, _ := strconv.Atoi(c.Query("pagina", "1"))
	limite, _ := strconv.Atoi(c.Query("limite", "20"))

	productos, err := h.service.BuscarProductos(c.Context(), query, pagina, limite)
	if err != nil {
		return response.Internal(c, "Error al buscar productos")
	}

	return response.Ok(c, productos, "Búsqueda completada")
}

// ListarCategorias maneja GET /api/v1/categorias
func (h *Handler) ListarCategorias(c *fiber.Ctx) error {
	categorias, err := h.service.ListarCategorias(c.Context())
	if err != nil {
		return response.Internal(c, "Error al listar categorías")
	}

	return response.Ok(c, categorias, "Categorías obtenidas")
}

// ObtenerCategoria maneja GET /api/v1/categorias/:slug
func (h *Handler) ObtenerCategoria(c *fiber.Ctx) error {
	slug := c.Params("slug")

	categoria, err := h.service.ObtenerCategoriaPorSlug(c.Context(), slug)
	if err != nil {
		return response.NotFound(c, "Categoría no encontrada")
	}

	return response.Ok(c, categoria, "Categoría obtenida")
}

// Crear maneja POST /api/v1/admin/productos (solo admin)
func (h *Handler) Crear(c *fiber.Ctx) error {
	var input CrearProductoInput
	if err := c.BodyParser(&input); err != nil {
		return response.BadRequest(c, "Cuerpo inválido")
	}
	producto, err := h.service.CrearProducto(c.Context(), input)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}
	return response.Created(c, producto, "Producto creado")
}

// Actualizar maneja PUT /api/v1/admin/productos/:id (solo admin)
func (h *Handler) Actualizar(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "ID inválido")
	}
	var input CrearProductoInput
	if err := c.BodyParser(&input); err != nil {
		return response.BadRequest(c, "Cuerpo inválido")
	}
	producto, err := h.service.ActualizarProducto(c.Context(), id, input)
	if err != nil {
		return response.BadRequest(c, err.Error())
	}
	return response.Ok(c, producto, "Producto actualizado")
}

// Eliminar maneja DELETE /api/v1/admin/productos/:id (solo admin)
func (h *Handler) Eliminar(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "ID inválido")
	}
	if err := h.service.EliminarProducto(c.Context(), id); err != nil {
		return response.BadRequest(c, err.Error())
	}
	return response.Ok(c, nil, "Producto eliminado")
}
