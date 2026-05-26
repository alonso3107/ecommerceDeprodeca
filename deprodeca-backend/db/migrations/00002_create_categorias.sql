-- +goose Up
-- +goose StatementBegin
CREATE TABLE categorias (
    id BIGSERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    slug VARCHAR(120) NOT NULL UNIQUE,
    descripcion TEXT NOT NULL DEFAULT '',
    imagen_url VARCHAR(500) NOT NULL DEFAULT '',
    activo BOOLEAN NOT NULL DEFAULT true
);

CREATE INDEX idx_categorias_slug ON categorias(slug);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categorias;
-- +goose StatementEnd
