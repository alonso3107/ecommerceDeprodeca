package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository define las operaciones de persistencia para auth.
// Interface pequeña y enfocada según el Principio de Segregación de Interfaces (SOLID).
type Repository interface {
	CrearUsuario(ctx context.Context, input RegistroInput, passwordHash string) (Usuario, error)
	ObtenerUsuarioPorEmail(ctx context.Context, email string) (Usuario, error)
}

// repositoryPG es la implementación concreta con PostgreSQL.
type repositoryPG struct {
	db *pgxpool.Pool
}

// NewRepository crea una nueva instancia del repositorio de auth.
func NewRepository(db *pgxpool.Pool) Repository {
	return &repositoryPG{db: db}
}

func (r *repositoryPG) CrearUsuario(ctx context.Context, input RegistroInput, passwordHash string) (Usuario, error) {
	var u Usuario

	err := r.db.QueryRow(ctx,
		`INSERT INTO usuarios (email, password_hash, nombre, empresa, ruc, telefono, rol)
		 VALUES ($1, $2, $3, $4, $5, $6, 'proveedor')
		 RETURNING id, email, password_hash, nombre, empresa, ruc, telefono, rol, created_at, updated_at`,
		input.Email, passwordHash, input.Nombre, input.Empresa, input.RUC, input.Telefono,
	).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Nombre, &u.Empresa,
		&u.RUC, &u.Telefono, &u.Rol, &u.CreatedAt, &u.UpdatedAt)

	return u, err
}

func (r *repositoryPG) ObtenerUsuarioPorEmail(ctx context.Context, email string) (Usuario, error) {
	var u Usuario

	err := r.db.QueryRow(ctx,
		`SELECT id, email, password_hash, nombre, empresa, ruc, telefono, rol, created_at, updated_at
		 FROM usuarios WHERE email = $1`,
		email,
	).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Nombre, &u.Empresa,
		&u.RUC, &u.Telefono, &u.Rol, &u.CreatedAt, &u.UpdatedAt)

	return u, err
}
