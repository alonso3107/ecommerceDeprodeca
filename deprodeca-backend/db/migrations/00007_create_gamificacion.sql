-- +goose Up
-- +goose StatementBegin
CREATE TABLE gamificacion (
    id BIGSERIAL PRIMARY KEY,
    usuario_id BIGINT NOT NULL UNIQUE REFERENCES usuarios(id),
    puntos_totales BIGINT NOT NULL DEFAULT 0 CHECK (puntos_totales >= 0),
    rango VARCHAR(20) NOT NULL DEFAULT 'bronce'
        CHECK (rango IN ('bronce', 'plata', 'oro', 'platino')),
    volumen_total DECIMAL(14, 2) NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_gamificacion_usuario ON gamificacion(usuario_id);
CREATE INDEX idx_gamificacion_rango ON gamificacion(rango);
CREATE INDEX idx_gamificacion_puntos ON gamificacion(puntos_totales DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS gamificacion;
-- +goose StatementEnd
