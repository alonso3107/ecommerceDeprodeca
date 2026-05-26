package gamificacion

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository define operaciones de persistencia para gamificación.
type Repository interface {
	ObtenerPorUsuario(ctx context.Context, usuarioID int64) (Gamificacion, error)
	Crear(ctx context.Context, usuarioID int64) (Gamificacion, error)
	ActualizarPuntosYRango(ctx context.Context, usuarioID int64, puntos int64, volumen float64, rango Rango) error
	ObtenerLeaderboard(ctx context.Context, limite int) ([]LeaderboardItem, error)
}

type repositoryPG struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repositoryPG{db: db}
}

func (r *repositoryPG) ObtenerPorUsuario(ctx context.Context, usuarioID int64) (Gamificacion, error) {
	var g Gamificacion
	err := r.db.QueryRow(ctx,
		`SELECT id, usuario_id, puntos_totales, rango, volumen_total, updated_at
		 FROM gamificacion WHERE usuario_id = $1`, usuarioID,
	).Scan(&g.ID, &g.UsuarioID, &g.PuntosTotales, &g.Rango, &g.VolumenTotal, &g.UpdatedAt)
	return g, err
}

func (r *repositoryPG) Crear(ctx context.Context, usuarioID int64) (Gamificacion, error) {
	var g Gamificacion
	err := r.db.QueryRow(ctx,
		`INSERT INTO gamificacion (usuario_id, puntos_totales, rango, volumen_total)
		 VALUES ($1, 0, 'bronce', 0)
		 RETURNING id, usuario_id, puntos_totales, rango, volumen_total, updated_at`,
		usuarioID,
	).Scan(&g.ID, &g.UsuarioID, &g.PuntosTotales, &g.Rango, &g.VolumenTotal, &g.UpdatedAt)
	return g, err
}

func (r *repositoryPG) ActualizarPuntosYRango(ctx context.Context, usuarioID int64, puntos int64, volumen float64, rango Rango) error {
	_, err := r.db.Exec(ctx,
		`UPDATE gamificacion
		 SET puntos_totales = puntos_totales + $2,
		     volumen_total = volumen_total + $3,
		     rango = $4,
		     updated_at = NOW()
		 WHERE usuario_id = $1`,
		usuarioID, puntos, volumen, string(rango),
	)
	return err
}

func (r *repositoryPG) ObtenerLeaderboard(ctx context.Context, limite int) ([]LeaderboardItem, error) {
	rows, err := r.db.Query(ctx,
		`SELECT g.usuario_id, g.puntos_totales, g.rango, g.volumen_total,
		        u.nombre, u.empresa
		 FROM gamificacion g
		 JOIN usuarios u ON u.id = g.usuario_id
		 ORDER BY g.puntos_totales DESC
		 LIMIT $1`, limite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []LeaderboardItem
	posicion := 1
	for rows.Next() {
		var item LeaderboardItem
		if err := rows.Scan(&item.UsuarioID, &item.PuntosTotales, &item.Rango,
			&item.VolumenTotal, &item.UsuarioNombre, &item.UsuarioEmpresa); err != nil {
			return nil, err
		}
		item.Posicion = posicion
		items = append(items, item)
		posicion++
	}

	return items, nil
}
