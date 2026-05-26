package response

import "github.com/gofiber/fiber/v2"

// APIResponse es la envoltura estándar para todas las respuestas JSON de la API.
// Formato: { "success": bool, "data": any, "message": string }
type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

// Ok envía una respuesta exitosa con código 200 y los datos proporcionados.
func Ok(c *fiber.Ctx, data any, message string) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

// Created envía una respuesta exitosa con código 201.
func Created(c *fiber.Ctx, data any, message string) error {
	return c.Status(fiber.StatusCreated).JSON(APIResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

// Error envía una respuesta de error con el código HTTP especificado.
func Error(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(APIResponse{
		Success: false,
		Data:    nil,
		Message: message,
	})
}

// BadRequest envía un error 400.
func BadRequest(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusBadRequest, message)
}

// Unauthorized envía un error 401.
func Unauthorized(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusUnauthorized, message)
}

// Forbidden envía un error 403.
func Forbidden(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusForbidden, message)
}

// NotFound envía un error 404.
func NotFound(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusNotFound, message)
}

// Internal envía un error 500.
func Internal(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusInternalServerError, message)
}
