package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Logger registra método, ruta, status y duración de cada request.
func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		inicio := time.Now()

		err := c.Next()

		duracion := time.Since(inicio)
		status := c.Response().StatusCode()
		metodo := c.Method()
		ruta := c.Path()

		log.Printf("[HTTP] %s %s → %d (%s)", metodo, ruta, status, duracion)

		return err
	}
}
