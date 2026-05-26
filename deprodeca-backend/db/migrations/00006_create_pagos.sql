-- +goose Up
-- +goose StatementBegin
CREATE TABLE pagos (
    id BIGSERIAL PRIMARY KEY,
    pedido_id BIGINT NOT NULL REFERENCES pedidos(id),
    monto DECIMAL(12, 2) NOT NULL CHECK (monto > 0),
    metodo VARCHAR(30) NOT NULL DEFAULT 'transferencia'
        CHECK (metodo IN ('transferencia', 'yape', 'plin', 'efectivo')),
    estado VARCHAR(20) NOT NULL DEFAULT 'pendiente'
        CHECK (estado IN ('pendiente', 'verificado', 'rechazado')),
    comprobante_url VARCHAR(500) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_pagos_pedido ON pagos(pedido_id);
CREATE INDEX idx_pagos_estado ON pagos(estado);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pagos;
-- +goose StatementEnd
