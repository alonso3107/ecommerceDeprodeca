package pedidos

import "time"

// Pedido representa una orden de compra.
type Pedido struct {
	ID        int64     `json:"id"`
	UsuarioID int64     `json:"usuario_id"`
	Total     float64   `json:"total"`
	Estado    string    `json:"estado"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Campos JOIN (admin)
	UsuarioNombre string `json:"usuario_nombre,omitempty"`
	UsuarioEmail  string `json:"usuario_email,omitempty"`
	Detalles      []PedidoDetalle `json:"detalles,omitempty"`
}

// PedidoDetalle es una línea de pedido.
type PedidoDetalle struct {
	ID              int64   `json:"id"`
	PedidoID        int64   `json:"pedido_id"`
	ProductoID      int64   `json:"producto_id"`
	ProductoNombre  string  `json:"producto_nombre,omitempty"`
	ProductoImagen  string  `json:"producto_imagen,omitempty"`
	Cantidad        int     `json:"cantidad"`
	PrecioUnitario  float64 `json:"precio_unitario"`
	Subtotal        float64 `json:"subtotal"`
}

// CrearPedidoInput son los datos para crear un pedido.
type CrearPedidoInput struct {
	Items []ItemInput `json:"items"`
}

// ItemInput es un producto y cantidad para el pedido.
type ItemInput struct {
	ProductoID int64 `json:"producto_id"`
	Cantidad   int   `json:"cantidad"`
}
