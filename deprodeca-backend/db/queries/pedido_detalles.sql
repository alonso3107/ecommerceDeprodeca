-- name: CrearPedidoDetalle :one
INSERT INTO pedido_detalles (pedido_id, producto_id, cantidad, precio_unitario, subtotal)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, pedido_id, producto_id, cantidad, precio_unitario, subtotal;

-- name: ListarDetallesPorPedido :many
SELECT pd.id, pd.pedido_id, pd.producto_id, pd.cantidad, pd.precio_unitario, pd.subtotal,
       p.nombre AS producto_nombre, p.imagen_url AS producto_imagen
FROM pedido_detalles pd
JOIN productos p ON p.id = pd.producto_id
WHERE pd.pedido_id = $1
ORDER BY pd.id;
