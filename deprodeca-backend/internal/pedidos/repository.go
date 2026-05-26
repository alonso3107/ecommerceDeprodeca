package pedidos

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository define las operaciones de persistencia para pedidos.
type Repository interface {
	CrearPedido(ctx context.Context, usuarioID int64, total float64) (int64, error)
	AgregarDetalle(ctx context.Context, pedidoID, productoID int64, cantidad int, precioUnitario, subtotal float64) error
	ObtenerPedidoPorID(ctx context.Context, id int64) (Pedido, error)
	ListarDetallesPorPedido(ctx context.Context, pedidoID int64) ([]PedidoDetalle, error)
	ListarPedidosPorUsuario(ctx context.Context, usuarioID int64, pagina, limite int) ([]Pedido, error)
	ListarTodos(ctx context.Context, pagina, limite int) ([]Pedido, error)
	ActualizarEstado(ctx context.Context, pedidoID int64, estado string) error
}

type repositoryPG struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repositoryPG{db: db}
}

func (r *repositoryPG) CrearPedido(ctx context.Context, usuarioID int64, total float64) (int64, error) {
	var id int64
	err := r.db.QueryRow(ctx,
		`INSERT INTO pedidos (usuario_id, total, estado)
		 VALUES ($1, $2, 'pendiente') RETURNING id`,
		usuarioID, total,
	).Scan(&id)
	return id, err
}

func (r *repositoryPG) AgregarDetalle(ctx context.Context, pedidoID, productoID int64, cantidad int, precioUnitario, subtotal float64) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO pedido_detalles (pedido_id, producto_id, cantidad, precio_unitario, subtotal)
		 VALUES ($1, $2, $3, $4, $5)`,
		pedidoID, productoID, cantidad, precioUnitario, subtotal)
	return err
}

func (r *repositoryPG) ObtenerPedidoPorID(ctx context.Context, id int64) (Pedido, error) {
	var p Pedido
	err := r.db.QueryRow(ctx,
		`SELECT id, usuario_id, total, estado, created_at, updated_at
		 FROM pedidos WHERE id = $1`, id,
	).Scan(&p.ID, &p.UsuarioID, &p.Total, &p.Estado, &p.CreatedAt, &p.UpdatedAt)
	return p, err
}

func (r *repositoryPG) ListarDetallesPorPedido(ctx context.Context, pedidoID int64) ([]PedidoDetalle, error) {
	rows, err := r.db.Query(ctx,
		`SELECT pd.id, pd.pedido_id, pd.producto_id, pd.cantidad, pd.precio_unitario, pd.subtotal,
		        p.nombre AS producto_nombre, p.imagen_url AS producto_imagen
		 FROM pedido_detalles pd
		 JOIN productos p ON p.id = pd.producto_id
		 WHERE pd.pedido_id = $1 ORDER BY pd.id`, pedidoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var detalles []PedidoDetalle
	for rows.Next() {
		var d PedidoDetalle
		if err := rows.Scan(&d.ID, &d.PedidoID, &d.ProductoID, &d.Cantidad,
			&d.PrecioUnitario, &d.Subtotal, &d.ProductoNombre, &d.ProductoImagen); err != nil {
			return nil, err
		}
		detalles = append(detalles, d)
	}
	return detalles, nil
}

func (r *repositoryPG) ListarPedidosPorUsuario(ctx context.Context, usuarioID int64, pagina, limite int) ([]Pedido, error) {
	offset := (pagina - 1) * limite

	rows, err := r.db.Query(ctx,
		`SELECT id, usuario_id, total, estado, created_at, updated_at
		 FROM pedidos WHERE usuario_id = $1
		 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
		usuarioID, limite, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []Pedido
	for rows.Next() {
		var p Pedido
		if err := rows.Scan(&p.ID, &p.UsuarioID, &p.Total, &p.Estado, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		pedidos = append(pedidos, p)
	}

	return pedidos, nil
}

func (r *repositoryPG) ActualizarEstado(ctx context.Context, pedidoID int64, estado string) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE pedidos SET estado = $1, updated_at = NOW() WHERE id = $2`,
		estado, pedidoID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("pedido no encontrado: %d", pedidoID)
	}
	return nil
}

func (r *repositoryPG) ListarTodos(ctx context.Context, pagina, limite int) ([]Pedido, error) {
	offset := (pagina - 1) * limite

	rows, err := r.db.Query(ctx,
		`SELECT p.id, p.usuario_id, p.total, p.estado, p.created_at, p.updated_at,
		        u.nombre AS usuario_nombre, u.email AS usuario_email
		 FROM pedidos p
		 JOIN usuarios u ON u.id = p.usuario_id
		 ORDER BY p.created_at DESC LIMIT $1 OFFSET $2`,
		limite, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []Pedido
	for rows.Next() {
		var p Pedido
		if err := rows.Scan(&p.ID, &p.UsuarioID, &p.Total, &p.Estado,
			&p.CreatedAt, &p.UpdatedAt, &p.UsuarioNombre, &p.UsuarioEmail); err != nil {
			return nil, err
		}
		pedidos = append(pedidos, p)
	}

	return pedidos, nil
}

// ─── Helpers compartidos ───────────────────────────

// ObtenerProductoPrecioYStock obtiene precio y stock de un producto.
func ObtenerProductoPrecioYStock(ctx context.Context, db *pgxpool.Pool, productoID int64) (float64, int, string, error) {
	var precio float64
	var stock int
	var nombre string
	err := db.QueryRow(ctx,
		`SELECT precio, stock, nombre FROM productos WHERE id = $1 AND activo = true`, productoID,
	).Scan(&precio, &stock, &nombre)
	return precio, stock, nombre, err
}

// DescontarStock descuenta stock de un producto (con check de stock suficiente).
func DescontarStock(ctx context.Context, db *pgxpool.Pool, productoID int64, cantidad int) error {
	tag, err := db.Exec(ctx,
		`UPDATE productos SET stock = stock - $2, updated_at = NOW()
		 WHERE id = $1 AND stock >= $2`, productoID, cantidad)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("stock insuficiente para producto ID %d", productoID)
	}
	return nil
}
