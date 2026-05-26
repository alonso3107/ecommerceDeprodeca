-- name: CrearProducto :one
INSERT INTO productos (categoria_id, nombre, slug, descripcion, precio, stock, unidad, imagen_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, categoria_id, nombre, slug, descripcion, precio, stock, unidad, imagen_url, activo, created_at, updated_at;

-- name: ObtenerProductoPorID :one
SELECT p.id, p.categoria_id, p.nombre, p.slug, p.descripcion, p.precio, p.stock, p.unidad, p.imagen_url, p.activo, p.created_at, p.updated_at,
       c.nombre AS categoria_nombre, c.slug AS categoria_slug
FROM productos p
JOIN categorias c ON c.id = p.categoria_id
WHERE p.id = $1;

-- name: ObtenerProductoPorSlug :one
SELECT p.id, p.categoria_id, p.nombre, p.slug, p.descripcion, p.precio, p.stock, p.unidad, p.imagen_url, p.activo, p.created_at, p.updated_at,
       c.nombre AS categoria_nombre, c.slug AS categoria_slug
FROM productos p
JOIN categorias c ON c.id = p.categoria_id
WHERE p.slug = $1;

-- name: ListarProductos :many
SELECT p.id, p.categoria_id, p.nombre, p.slug, p.descripcion, p.precio, p.stock, p.unidad, p.imagen_url, p.activo, p.created_at, p.updated_at,
       c.nombre AS categoria_nombre, c.slug AS categoria_slug
FROM productos p
JOIN categorias c ON c.id = p.categoria_id
WHERE p.activo = true
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListarProductosPorCategoria :many
SELECT p.id, p.categoria_id, p.nombre, p.slug, p.descripcion, p.precio, p.stock, p.unidad, p.imagen_url, p.activo, p.created_at,
       c.nombre AS categoria_nombre
FROM productos p
JOIN categorias c ON c.id = p.categoria_id
WHERE p.activo = true AND p.categoria_id = $1
ORDER BY p.nombre
LIMIT $2 OFFSET $3;

-- name: BuscarProductos :many
SELECT p.id, p.categoria_id, p.nombre, p.slug, p.precio, p.stock, p.unidad, p.imagen_url, p.activo,
       c.nombre AS categoria_nombre
FROM productos p
JOIN categorias c ON c.id = p.categoria_id
WHERE p.activo = true AND (p.nombre ILIKE '%' || $1 || '%' OR p.descripcion ILIKE '%' || $1 || '%')
ORDER BY p.nombre
LIMIT $2 OFFSET $3;

-- name: ActualizarProducto :exec
UPDATE productos SET categoria_id = $2, nombre = $3, slug = $4, descripcion = $5,
                     precio = $6, stock = $7, unidad = $8, imagen_url = $9, updated_at = NOW()
WHERE id = $1;

-- name: ActualizarStockProducto :exec
UPDATE productos SET stock = stock - $2, updated_at = NOW() WHERE id = $1 AND stock >= $2;

-- name: DesactivarProducto :exec
UPDATE productos SET activo = false, updated_at = NOW() WHERE id = $1;
