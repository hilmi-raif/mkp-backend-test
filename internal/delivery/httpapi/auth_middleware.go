package httpapi

import (
	"strings"

	"mkp-backend-test/internal/helper"

	"github.com/gofiber/fiber/v2"
)

func RequireToken(secret []byte) fiber.Handler {
	return func(c *fiber.Ctx) error {
		parts := strings.Fields(c.Get("Authorization"))
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return helper.Unauthorized(c)
		}
		if !helper.ValidToken(secret, parts[1]) {
			return helper.Unauthorized(c)
		}
		return c.Next()
	}
}
