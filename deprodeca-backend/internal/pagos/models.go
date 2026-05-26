package pagos

import "time"

// Pago representa un comprobante de pago registrado.
type Pago struct {
	ID             int64     `json:"id"`
	PedidoID       int64     `json:"pedido_id"`
	Monto          float64   `json:"monto"`
	Metodo         string    `json:"metodo"`
	Estado         string    `json:"estado"`
	ComprobanteURL string    `json:"comprobante_url"`
	CreatedAt      time.Time `json:"created_at"`
}

// RegistrarPagoInput son los datos para subir un comprobante de pago.
type RegistrarPagoInput struct {
	PedidoID       int64   `json:"pedido_id"`
	Monto          float64 `json:"monto"`
	Metodo         string  `json:"metodo"`
	ComprobanteURL string  `json:"comprobante_url"`
}
