package contacto

import "time"

type ContactoInput struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
	Motivo   string `json:"motivo"`
	Mensaje  string `json:"mensaje"`
}

type Contacto struct {
	ID        int64     `json:"id"`
	Nombre    string    `json:"nombre"`
	Email     string    `json:"email"`
	Telefono  string    `json:"telefono"`
	Motivo    string    `json:"motivo"`
	Mensaje   string    `json:"mensaje"`
	CreatedAt time.Time `json:"created_at"`
}
