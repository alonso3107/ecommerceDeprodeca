package auth

import "time"

// Usuario representa un usuario del sistema (proveedor o admin).
type Usuario struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Nombre       string    `json:"nombre"`
	Empresa      string    `json:"empresa"`
	RUC          string    `json:"ruc"`
	Telefono     string    `json:"telefono"`
	Rol          string    `json:"rol"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// RegistroInput son los datos requeridos para crear una cuenta.
type RegistroInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nombre   string `json:"nombre"`
	Empresa  string `json:"empresa"`
	RUC      string `json:"ruc"`
	Telefono string `json:"telefono"`
}

// LoginInput son las credenciales de inicio de sesión.
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse es la respuesta de login/registro exitoso.
type AuthResponse struct {
	Token     string `json:"token"`
	ExpiraEn  string `json:"expira_en"`
	UsuarioID int64  `json:"usuario_id"`
	Email     string `json:"email"`
	Nombre    string `json:"nombre"`
	Rol       string `json:"rol"`
}
