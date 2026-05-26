-- +goose Up
-- +goose StatementBegin
CREATE TABLE usuarios (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    nombre VARCHAR(150) NOT NULL,
    empresa VARCHAR(200) NOT NULL,
    ruc VARCHAR(11) NOT NULL,
    telefono VARCHAR(20) NOT NULL DEFAULT '',
    rol VARCHAR(20) NOT NULL DEFAULT 'proveedor' CHECK (rol IN ('proveedor', 'admin')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_usuarios_email ON usuarios(email);
CREATE INDEX idx_usuarios_rol ON usuarios(rol);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS usuarios;
-- +goose StatementEnd
