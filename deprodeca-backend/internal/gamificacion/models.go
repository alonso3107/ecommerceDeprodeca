package gamificacion

import "time"

// Gamificacion representa el estado de gamificación de un usuario.
type Gamificacion struct {
	ID            int64     `json:"id"`
	UsuarioID     int64     `json:"usuario_id"`
	PuntosTotales int64     `json:"puntos_totales"`
	Rango         Rango     `json:"rango"`
	VolumenTotal  float64   `json:"volumen_total"`
	UpdatedAt     time.Time `json:"updated_at"`

	// Campos JOIN
	UsuarioNombre  string `json:"usuario_nombre,omitempty"`
	UsuarioEmpresa string `json:"usuario_empresa,omitempty"`
}

// LeaderboardItem es una entrada del ranking.
type LeaderboardItem struct {
	UsuarioID      int64   `json:"usuario_id"`
	UsuarioNombre  string  `json:"usuario_nombre"`
	UsuarioEmpresa string  `json:"usuario_empresa"`
	PuntosTotales  int64   `json:"puntos_totales"`
	Rango          Rango   `json:"rango"`
	VolumenTotal   float64 `json:"volumen_total"`
	Posicion       int     `json:"posicion"`
}
