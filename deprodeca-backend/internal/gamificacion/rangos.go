package gamificacion

// Rango representa el nivel de gamificación de un proveedor.
type Rango string

const (
	RangoBronce  Rango = "bronce"
	RangoPlata   Rango = "plata"
	RangoOro     Rango = "oro"
	RangoPlatino Rango = "platino"
)

// Umbrales de volumen de compra para cada rango.
const (
	UmbralPlata   float64 = 5000  // S/ 5,000
	UmbralOro     float64 = 20000 // S/ 20,000
	UmbralPlatino float64 = 50000 // S/ 50,000
)

// ordenRangos mapea cada rango a su nivel numérico para comparación.
var ordenRangos = map[Rango]int{
	RangoBronce:  0,
	RangoPlata:   1,
	RangoOro:     2,
	RangoPlatino: 3,
}

// CalcularRango determina el rango según el volumen total de compras.
func CalcularRango(volumenTotal float64) Rango {
	switch {
	case volumenTotal >= UmbralPlatino:
		return RangoPlatino
	case volumenTotal >= UmbralOro:
		return RangoOro
	case volumenTotal >= UmbralPlata:
		return RangoPlata
	default:
		return RangoBronce
	}
}

// RangoSuperior determina si el nuevo rango es superior al actual.
func RangoSuperior(actual, nuevo Rango) bool {
	return ordenRangos[nuevo] > ordenRangos[actual]
}

// NombreRango devuelve el nombre legible del rango con la primera letra en mayúscula.
func NombreRango(r Rango) string {
	switch r {
	case RangoBronce:
		return "Bronce"
	case RangoPlata:
		return "Plata"
	case RangoOro:
		return "Oro"
	case RangoPlatino:
		return "Platino"
	default:
		return "Bronce"
	}
}
