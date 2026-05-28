package contacto

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CrearContacto(ctx context.Context, input ContactoInput) (Contacto, error)
}

type repositoryPG struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repositoryPG{db: db}
}

func (r *repositoryPG) CrearContacto(ctx context.Context, input ContactoInput) (Contacto, error) {
	if _, err := r.db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS consultas_contacto (
			id         BIGSERIAL PRIMARY KEY,
			nombre     TEXT NOT NULL,
			email      TEXT NOT NULL,
			telefono   TEXT NOT NULL DEFAULT '',
			motivo     TEXT NOT NULL DEFAULT 'No especificado',
			mensaje    TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`); err != nil {
		return Contacto{}, err
	}

	var contacto Contacto
	err := r.db.QueryRow(ctx, `
		INSERT INTO consultas_contacto (nombre, email, telefono, motivo, mensaje)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, nombre, email, telefono, motivo, mensaje, created_at
	`,
		input.Nombre,
		input.Email,
		input.Telefono,
		input.Motivo,
		input.Mensaje,
	).Scan(
		&contacto.ID,
		&contacto.Nombre,
		&contacto.Email,
		&contacto.Telefono,
		&contacto.Motivo,
		&contacto.Mensaje,
		&contacto.CreatedAt,
	)
	if err != nil {
		return Contacto{}, err
	}

	return contacto, nil
}
