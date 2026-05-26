-- name: CrearPago :one
INSERT INTO pagos (pedido_id, monto, metodo, comprobante_url)
VALUES ($1, $2, $3, $4)
RETURNING id, pedido_id, monto, metodo, estado, comprobante_url, created_at;

-- name: ObtenerPagoPorPedido :one
SELECT id, pedido_id, monto, metodo, estado, comprobante_url, created_at
FROM pagos WHERE pedido_id = $1;

-- name: ActualizarEstadoPago :exec
UPDATE pagos SET estado = $2 WHERE id = $1;

-- name: ListarPagosPorUsuario :many
SELECT pa.id, pa.pedido_id, pa.monto, pa.metodo, pa.estado, pa.comprobante_url, pa.created_at
FROM pagos pa
JOIN pedidos pe ON pe.id = pa.pedido_id
WHERE pe.usuario_id = $1
ORDER BY pa.created_at DESC
LIMIT $2 OFFSET $3;
