-- name: CrearCategoria :one
INSERT INTO categorias (nombre, slug, descripcion, imagen_url)
VALUES ($1, $2, $3, $4)
RETURNING id, nombre, slug, descripcion, imagen_url, activo;

-- name: ObtenerCategoriaPorID :one
SELECT id, nombre, slug, descripcion, imagen_url, activo
FROM categorias WHERE id = $1;

-- name: ObtenerCategoriaPorSlug :one
SELECT id, nombre, slug, descripcion, imagen_url, activo
FROM categorias WHERE slug = $1;

-- name: ListarCategorias :many
SELECT id, nombre, slug, descripcion, imagen_url, activo
FROM categorias WHERE activo = true ORDER BY nombre;
