-- name: CrearPedido :one
INSERT INTO pedidos (usuario_id, total, estado)
VALUES ($1, $2, 'pendiente')
RETURNING id, usuario_id, total, estado, created_at, updated_at;

-- name: ObtenerPedidoPorID :one
SELECT id, usuario_id, total, estado, created_at, updated_at
FROM pedidos WHERE id = $1;

-- name: ListarPedidosPorUsuario :many
SELECT id, usuario_id, total, estado, created_at, updated_at
FROM pedidos WHERE usuario_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: ListarTodosPedidos :many
SELECT p.id, p.usuario_id, p.total, p.estado, p.created_at, p.updated_at,
       u.nombre AS usuario_nombre, u.email AS usuario_email
FROM pedidos p
JOIN usuarios u ON u.id = p.usuario_id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2;

-- name: ActualizarEstadoPedido :exec
UPDATE pedidos SET estado = $2, updated_at = NOW() WHERE id = $1;

-- name: ObtenerVolumenComprasUsuario :one
SELECT COALESCE(SUM(total), 0)::bigint AS volumen_total
FROM pedidos WHERE usuario_id = $1 AND estado IN ('confirmado', 'enviado', 'entregado');
