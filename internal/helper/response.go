package helper

import "github.com/gofiber/fiber/v2"

func ErrorJSON(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{"error": message})
}

func Unauthorized(c *fiber.Ctx) error {
	c.Set("WWW-Authenticate", "Bearer")
	return ErrorJSON(c, 401, "token tidak valid atau tidak ada")
}
