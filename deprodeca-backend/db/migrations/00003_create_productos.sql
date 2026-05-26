-- +goose Up
-- +goose StatementBegin
CREATE TABLE productos (
    id BIGSERIAL PRIMARY KEY,
    categoria_id BIGINT NOT NULL REFERENCES categorias(id),
    nombre VARCHAR(200) NOT NULL,
    slug VARCHAR(220) NOT NULL UNIQUE,
    descripcion TEXT NOT NULL DEFAULT '',
    precio DECIMAL(10, 2) NOT NULL CHECK (precio > 0),
    stock INT NOT NULL DEFAULT 0 CHECK (stock >= 0),
    unidad VARCHAR(30) NOT NULL DEFAULT 'unidad',
    imagen_url VARCHAR(500) NOT NULL DEFAULT '',
    activo BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_productos_categoria ON productos(categoria_id);
CREATE INDEX idx_productos_slug ON productos(slug);
CREATE INDEX idx_productos_activo ON productos(activo);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS productos;
-- +goose StatementEnd
