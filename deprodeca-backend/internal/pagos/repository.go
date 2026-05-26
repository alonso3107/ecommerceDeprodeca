package pagos

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository define operaciones de persistencia para pagos.
type Repository interface {
	Crear(ctx context.Context, input RegistrarPagoInput) (Pago, error)
	ObtenerPorID(ctx context.Context, id int64) (Pago, error)
	ObtenerPorPedido(ctx context.Context, pedidoID int64) (Pago, error)
	ActualizarEstado(ctx context.Context, id int64, estado string) error
	ListarPorUsuario(ctx context.Context, usuarioID int64, pagina, limite int) ([]Pago, error)
	ListarTodos(ctx context.Context, pagina, limite int) ([]Pago, error)
}

type repositoryPG struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repositoryPG{db: db}
}

func (r *repositoryPG) Crear(ctx context.Context, input RegistrarPagoInput) (Pago, error) {
	var p Pago
	err := r.db.QueryRow(ctx,
		`INSERT INTO pagos (pedido_id, monto, metodo, comprobante_url)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, pedido_id, monto, metodo, estado, comprobante_url, created_at`,
		input.PedidoID, input.Monto, input.Metodo, input.ComprobanteURL,
	).Scan(&p.ID, &p.PedidoID, &p.Monto, &p.Metodo, &p.Estado, &p.ComprobanteURL, &p.CreatedAt)

	return p, err
}

func (r *repositoryPG) ObtenerPorID(ctx context.Context, id int64) (Pago, error) {
	var p Pago
	err := r.db.QueryRow(ctx,
		`SELECT id, pedido_id, monto, metodo, estado, comprobante_url, created_at
		 FROM pagos WHERE id = $1`, id,
	).Scan(&p.ID, &p.PedidoID, &p.Monto, &p.Metodo, &p.Estado, &p.ComprobanteURL, &p.CreatedAt)
	return p, err
}

func (r *repositoryPG) ActualizarEstado(ctx context.Context, id int64, estado string) error {
	_, err := r.db.Exec(ctx,
		`UPDATE pagos SET estado = $2 WHERE id = $1`, id, estado)
	return err
}

func (r *repositoryPG) ObtenerPorPedido(ctx context.Context, pedidoID int64) (Pago, error) {
	var p Pago
	err := r.db.QueryRow(ctx,
		`SELECT id, pedido_id, monto, metodo, estado, comprobante_url, created_at
		 FROM pagos WHERE pedido_id = $1`, pedidoID,
	).Scan(&p.ID, &p.PedidoID, &p.Monto, &p.Metodo, &p.Estado, &p.ComprobanteURL, &p.CreatedAt)

	return p, err
}

func (r *repositoryPG) ListarPorUsuario(ctx context.Context, usuarioID int64, pagina, limite int) ([]Pago, error) {
	offset := (pagina - 1) * limite

	rows, err := r.db.Query(ctx,
		`SELECT pa.id, pa.pedido_id, pa.monto, pa.metodo, pa.estado, pa.comprobante_url, pa.created_at
		 FROM pagos pa
		 JOIN pedidos pe ON pe.id = pa.pedido_id
		 WHERE pe.usuario_id = $1
		 ORDER BY pa.created_at DESC
		 LIMIT $2 OFFSET $3`,
		usuarioID, limite, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pagos []Pago
	for rows.Next() {
		var p Pago
		if err := rows.Scan(&p.ID, &p.PedidoID, &p.Monto, &p.Metodo, &p.Estado, &p.ComprobanteURL, &p.CreatedAt); err != nil {
			return nil, err
		}
		pagos = append(pagos, p)
	}

	return pagos, nil
}

func (r *repositoryPG) ListarTodos(ctx context.Context, pagina, limite int) ([]Pago, error) {
	offset := (pagina - 1) * limite

	rows, err := r.db.Query(ctx,
		`SELECT id, pedido_id, monto, metodo, estado, comprobante_url, created_at
		 FROM pagos ORDER BY created_at DESC LIMIT $1 OFFSET $2`,
		limite, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pagos []Pago
	for rows.Next() {
		var p Pago
		if err := rows.Scan(&p.ID, &p.PedidoID, &p.Monto, &p.Metodo, &p.Estado, &p.ComprobanteURL, &p.CreatedAt); err != nil {
			return nil, err
		}
		pagos = append(pagos, p)
	}

	return pagos, nil
}
