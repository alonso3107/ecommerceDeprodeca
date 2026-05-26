package productos

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository define las operaciones de persistencia para productos y categorías.
type Repository interface {
	ListarProductos(ctx context.Context, p ListarParams) ([]Producto, error)
	ObtenerProductoPorSlug(ctx context.Context, slug string) (Producto, error)
	BuscarProductos(ctx context.Context, query string, pagina, limite int) ([]Producto, error)
	ListarCategorias(ctx context.Context) ([]Categoria, error)
	ObtenerCategoriaPorSlug(ctx context.Context, slug string) (Categoria, error)
}

type repositoryPG struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repositoryPG{db: db}
}

func (r *repositoryPG) ListarProductos(ctx context.Context, p ListarParams) ([]Producto, error) {
	query := `SELECT p.id, p.categoria_id, p.nombre, p.slug, p.descripcion,
	                  p.precio, p.stock, p.unidad, p.imagen_url, p.activo, p.created_at, p.updated_at,
	                  c.nombre AS categoria_nombre, c.slug AS categoria_slug
	           FROM productos p
	           JOIN categorias c ON c.id = p.categoria_id
	           WHERE p.activo = true`

	var args []any
	argIdx := 1

	if p.CategoriaID > 0 {
		query += fmt.Sprintf(" AND p.categoria_id = $%d", argIdx)
		args = append(args, p.CategoriaID)
		argIdx++
	}

	if p.Busqueda != "" {
		query += fmt.Sprintf(" AND (p.nombre ILIKE $%d OR p.descripcion ILIKE $%d)", argIdx, argIdx)
		args = append(args, "%"+p.Busqueda+"%")
		argIdx++
	}

	query += " ORDER BY p.created_at DESC"

	offset := (p.Pagina - 1) * p.Limite
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIdx, argIdx+1)
	args = append(args, p.Limite, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var pr Producto
		if err := rows.Scan(&pr.ID, &pr.CategoriaID, &pr.Nombre, &pr.Slug, &pr.Descripcion,
			&pr.Precio, &pr.Stock, &pr.Unidad, &pr.ImagenURL, &pr.Activo,
			&pr.CreatedAt, &pr.UpdatedAt, &pr.CategoriaNombre, &pr.CategoriaSlug); err != nil {
			return nil, err
		}
		productos = append(productos, pr)
	}

	return productos, nil
}

func (r *repositoryPG) ObtenerProductoPorSlug(ctx context.Context, slug string) (Producto, error) {
	var pr Producto

	err := r.db.QueryRow(ctx,
		`SELECT p.id, p.categoria_id, p.nombre, p.slug, p.descripcion,
		        p.precio, p.stock, p.unidad, p.imagen_url, p.activo, p.created_at, p.updated_at,
		        c.nombre AS categoria_nombre, c.slug AS categoria_slug
		 FROM productos p
		 JOIN categorias c ON c.id = p.categoria_id
		 WHERE p.slug = $1`, slug,
	).Scan(&pr.ID, &pr.CategoriaID, &pr.Nombre, &pr.Slug, &pr.Descripcion,
		&pr.Precio, &pr.Stock, &pr.Unidad, &pr.ImagenURL, &pr.Activo,
		&pr.CreatedAt, &pr.UpdatedAt, &pr.CategoriaNombre, &pr.CategoriaSlug)

	return pr, err
}

func (r *repositoryPG) BuscarProductos(ctx context.Context, query string, pagina, limite int) ([]Producto, error) {
	return r.ListarProductos(ctx, ListarParams{
		Pagina:   pagina,
		Limite:   limite,
		Busqueda: strings.TrimSpace(query),
	})
}

func (r *repositoryPG) ListarCategorias(ctx context.Context) ([]Categoria, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, nombre, slug, descripcion, imagen_url, activo
		 FROM categorias WHERE activo = true ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []Categoria
	for rows.Next() {
		var c Categoria
		if err := rows.Scan(&c.ID, &c.Nombre, &c.Slug, &c.Descripcion, &c.ImagenURL, &c.Activo); err != nil {
			return nil, err
		}
		categorias = append(categorias, c)
	}

	return categorias, nil
}

func (r *repositoryPG) ObtenerCategoriaPorSlug(ctx context.Context, slug string) (Categoria, error) {
	var c Categoria
	err := r.db.QueryRow(ctx,
		`SELECT id, nombre, slug, descripcion, imagen_url, activo
		 FROM categorias WHERE slug = $1`, slug,
	).Scan(&c.ID, &c.Nombre, &c.Slug, &c.Descripcion, &c.ImagenURL, &c.Activo)

	return c, err
}
