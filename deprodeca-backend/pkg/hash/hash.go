package hash

import (
	"golang.org/x/crypto/bcrypt"
)

const costo = 12 // Costo bcrypt: balance seguridad/rendimiento

// Hashear genera un hash bcrypt de la contraseña en texto plano.
func Hashear(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Verificar compara una contraseña en texto plano con su hash bcrypt.
func Verificar(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
