package usuarios

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository define las operaciones de persistencia para usuarios.
type Repository interface {
	ObtenerPorID(ctx context.Context, id int64) (Usuario, error)
	Actualizar(ctx context.Context, id int64, input ActualizarPerfilInput) error
}

type repositoryPG struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repositoryPG{db: db}
}

func (r *repositoryPG) ObtenerPorID(ctx context.Context, id int64) (Usuario, error) {
	var u Usuario

	err := r.db.QueryRow(ctx,
		`SELECT id, email, nombre, empresa, ruc, telefono, rol, created_at, updated_at
		 FROM usuarios WHERE id = $1`, id,
	).Scan(&u.ID, &u.Email, &u.Nombre, &u.Empresa, &u.RUC,
		&u.Telefono, &u.Rol, &u.CreatedAt, &u.UpdatedAt)

	return u, err
}

func (r *repositoryPG) Actualizar(ctx context.Context, id int64, input ActualizarPerfilInput) error {
	_, err := r.db.Exec(ctx,
		`UPDATE usuarios SET nombre = $2, empresa = $3, telefono = $4, updated_at = NOW()
		 WHERE id = $1`,
		id, input.Nombre, input.Empresa, input.Telefono,
	)
	return err
}
