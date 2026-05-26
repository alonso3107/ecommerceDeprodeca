-- name: ObtenerGamificacionPorUsuario :one
SELECT id, usuario_id, puntos_totales, rango, volumen_total, updated_at
FROM gamificacion WHERE usuario_id = $1;

-- name: CrearGamificacion :one
INSERT INTO gamificacion (usuario_id, puntos_totales, rango, volumen_total)
VALUES ($1, 0, 'bronce', 0)
RETURNING id, usuario_id, puntos_totales, rango, volumen_total, updated_at;

-- name: ActualizarPuntosYRango :exec
UPDATE gamificacion
SET puntos_totales = puntos_totales + $2, volumen_total = volumen_total + $3, rango = $4, updated_at = NOW()
WHERE usuario_id = $1;

-- name: ObtenerLeaderboard :many
SELECT g.usuario_id, g.puntos_totales, g.rango, g.volumen_total,
       u.nombre AS usuario_nombre, u.empresa AS usuario_empresa
FROM gamificacion g
JOIN usuarios u ON u.id = g.usuario_id
ORDER BY g.puntos_totales DESC
LIMIT $1;
