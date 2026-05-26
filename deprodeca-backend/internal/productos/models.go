package productos

import "time"

// Producto representa un producto del catálogo.
type Producto struct {
	ID              int64     `json:"id"`
	CategoriaID     int64     `json:"categoria_id"`
	CategoriaNombre string    `json:"categoria_nombre,omitempty"`
	CategoriaSlug   string    `json:"categoria_slug,omitempty"`
	Nombre          string    `json:"nombre"`
	Slug            string    `json:"slug"`
	Descripcion     string    `json:"descripcion"`
	Precio          float64   `json:"precio"`
	Stock           int       `json:"stock"`
	Unidad          string    `json:"unidad"`
	ImagenURL       string    `json:"imagen_url"`
	Activo          bool      `json:"activo"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Categoria representa una categoría de productos.
type Categoria struct {
	ID          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Slug        string `json:"slug"`
	Descripcion string `json:"descripcion"`
	ImagenURL   string `json:"imagen_url"`
	Activo      bool   `json:"activo"`
}

// ListarParams define parámetros de paginación y filtro.
type ListarParams struct {
	Pagina      int
	Limite      int
	Busqueda    string
	CategoriaID int64
}

// CrearProductoInput son los datos para crear/actualizar un producto.
type CrearProductoInput struct {
	Nombre      string  `json:"nombre"`
	CategoriaID int64   `json:"categoria_id"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
	Unidad      string  `json:"unidad"`
	Descripcion string  `json:"descripcion"`
	ImagenURL   string  `json:"imagen_url"`
}
