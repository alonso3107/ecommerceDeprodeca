-- name: CrearUsuario :one
INSERT INTO usuarios (email, password_hash, nombre, empresa, ruc, telefono, rol)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, email, nombre, empresa, ruc, telefono, rol, created_at, updated_at;

-- name: ObtenerUsuarioPorID :one
SELECT id, email, password_hash, nombre, empresa, ruc, telefono, rol, created_at, updated_at
FROM usuarios WHERE id = $1;

-- name: ObtenerUsuarioPorEmail :one
SELECT id, email, password_hash, nombre, empresa, ruc, telefono, rol, created_at, updated_at
FROM usuarios WHERE email = $1;

-- name: ListarUsuarios :many
SELECT id, email, nombre, empresa, ruc, telefono, rol, created_at
FROM usuarios ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: ActualizarUsuario :exec
UPDATE usuarios SET nombre = $2, empresa = $3, telefono = $4, updated_at = NOW()
WHERE id = $1;

-- name: ActualizarPassword :exec
UPDATE usuarios SET password_hash = $2, updated_at = NOW() WHERE id = $1;
