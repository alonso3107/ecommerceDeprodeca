package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

// Claims personalizados para el token JWT de DEPRODECA.
type Claims struct {
	UsuarioID int64  `json:"usuario_id"`
	Email     string `json:"email"`
	Rol       string `json:"rol"` // proveedor | admin
	jwtlib.RegisteredClaims
}

// GenerarToken crea un JWT firmado con los claims del usuario.
func GenerarToken(secret string, expirationHours int, usuarioID int64, email, rol string) (string, time.Time, error) {
	expira := time.Now().Add(time.Duration(expirationHours) * time.Hour)

	claims := Claims{
		UsuarioID: usuarioID,
		Email:     email,
		Rol:       rol,
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(expira),
			IssuedAt:  jwtlib.NewNumericDate(time.Now()),
			Issuer:    "deprodeca-api",
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenStr, expira, nil
}

// ValidarToken decodifica y valida un token JWT, devolviendo los claims.
func ValidarToken(tokenStr, secret string) (*Claims, error) {
	token, err := jwtlib.ParseWithClaims(tokenStr, &Claims{}, func(t *jwtlib.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwtlib.ErrSignatureInvalid
	}

	return claims, nil
}
