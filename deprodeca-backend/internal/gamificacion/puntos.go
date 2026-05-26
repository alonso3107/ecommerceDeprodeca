package gamificacion

// PuntosPorSol define cuántos puntos se otorgan por cada sol gastado.
const PuntosPorSol = 1

// CalcularPuntos calcula los puntos ganados por un monto de compra.
func CalcularPuntos(monto float64) int64 {
	return int64(monto) * PuntosPorSol
}
