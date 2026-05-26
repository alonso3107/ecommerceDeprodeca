package middleware

import (
	"strings"

	"deprodeca-backend/pkg/jwt"
	"deprodeca-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// JWTAuth protege rutas que requieren autenticación.
// Extrae el token del header Authorization: Bearer <token>.
func JWTAuth(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return response.Unauthorized(c, "Token no proporcionado")
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return response.Unauthorized(c, "Formato de token inválido. Usar: Bearer <token>")
		}

		tokenStr := parts[1]

		claims, err := jwt.ValidarToken(tokenStr, secret)
		if err != nil {
			return response.Unauthorized(c, "Token inválido o expirado")
		}

		// Inyectar claims en el contexto para los handlers
		c.Locals("usuarioID", claims.UsuarioID)
		c.Locals("email", claims.Email)
		c.Locals("rol", claims.Rol)

		return c.Next()
	}
}

// RequireRol verifica que el usuario autenticado tenga uno de los roles permitidos.
// Debe usarse DESPUÉS de JWTAuth.
func RequireRol(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rol := c.Locals("rol")
		if rol == nil {
			return response.Forbidden(c, "Acceso denegado: sin rol asignado")
		}

		rolStr := rol.(string)
		for _, permitido := range roles {
			if rolStr == permitido {
				return c.Next()
			}
		}

		return response.Forbidden(c, "Acceso denegado: rol no autorizado")
	}
}
