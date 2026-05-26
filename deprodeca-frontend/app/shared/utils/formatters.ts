// Formateadores de datos para DEPRODECA.
// Precios en soles peruanos, fechas en español, números legibles.

/**
 * Formatea un número como precio en soles peruanos.
 * Ej: 12.5 → "S/ 12.50"
 */
export function formatearPrecio(precio: number): string {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
    minimumFractionDigits: 2,
  }).format(precio)
}

/**
 * Formatea una fecha ISO a formato legible en español.
 * Ej: "2025-05-25T10:30:00Z" → "25 may 2025, 05:30"
 */
export function formatearFecha(iso: string, conHora = false): string {
  const opciones: Intl.DateTimeFormatOptions = {
    day: "numeric",
    month: "short",
    year: "numeric",
  }
  if (conHora) {
    opciones.hour = "2-digit"
    opciones.minute = "2-digit"
  }
  return new Date(iso).toLocaleDateString("es-PE", opciones)
}

/**
 * Formatea un número grande con separadores de miles.
 * Ej: 50000 → "50,000"
 */
export function formatearNumero(n: number): string {
  return new Intl.NumberFormat("es-PE").format(n)
}

/**
 * Capitaliza la primera letra de un string.
 */
export function capitalizar(texto: string): string {
  return texto.charAt(0).toUpperCase() + texto.slice(1)
}
