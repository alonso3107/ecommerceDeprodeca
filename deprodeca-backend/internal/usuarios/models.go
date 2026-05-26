package usuarios

import "time"

// Usuario representa el perfil de un usuario (sin datos sensibles).
type Usuario struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Nombre    string    `json:"nombre"`
	Empresa   string    `json:"empresa"`
	RUC       string    `json:"ruc"`
	Telefono  string    `json:"telefono"`
	Rol       string    `json:"rol"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ActualizarPerfilInput contiene los campos editables del perfil.
type ActualizarPerfilInput struct {
	Nombre   string `json:"nombre"`
	Empresa  string `json:"empresa"`
	Telefono string `json:"telefono"`
}
