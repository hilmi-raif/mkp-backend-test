package httpapi

import (
	"mkp-backend-test/internal/dto"
	"mkp-backend-test/internal/helper"
	"mkp-backend-test/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	auth     *service.Auth
	validate *validator.Validate
}

func NewAuth(auth *service.Auth, validate *validator.Validate) *Auth {
	return &Auth{auth: auth, validate: validate}
}

func (h *Auth) Login(c *fiber.Ctx) error {
	var input dto.Login
	if c.BodyParser(&input) != nil || h.validate.Struct(input) != nil {
		return helper.ErrorJSON(c, 400, "email atau password tidak valid")
	}
	token, err := h.auth.Login(c.UserContext(), input)
	if err != nil {
		return helper.HandleError(c, err)
	}
	return c.JSON(fiber.Map{"access_token": token})
}
