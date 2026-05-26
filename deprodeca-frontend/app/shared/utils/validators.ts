// Validadores comunes para formularios DEPRODECA.

/** Valida que un email tenga formato correcto. */
export function emailValido(email: string): boolean {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)
}

/** Valida que un RUC peruano tenga 11 dígitos y empiece con 10 o 20. */
export function rucValido(ruc: string): boolean {
  return /^(10|20)\d{9}$/.test(ruc)
}

/** Valida que la contraseña tenga al menos N caracteres. */
export function passwordValido(password: string, min = 6): boolean {
  return password.length >= min
}

/** Valida que un campo no esté vacío. */
export function requerido(valor: string): boolean {
  return valor.trim().length > 0
}

/** Valida que un número de teléfono peruano tenga 9 dígitos (opcional). */
export function telefonoValido(telefono: string): boolean {
  if (!telefono) return true // opcional
  return /^\d{9}$/.test(telefono)
}
