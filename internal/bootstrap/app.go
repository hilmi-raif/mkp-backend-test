package bootstrap

import (
	"mkp-backend-test/internal/config"
	"mkp-backend-test/internal/delivery/httpapi"
	"mkp-backend-test/internal/repository"
	"mkp-backend-test/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func New(c config.Config) (*fiber.App, error) {
	store, err := repository.Open(c.DSN)
	if err != nil {
		return nil, err
	}
	validate := validator.New()
	auth := httpapi.NewAuth(service.NewAuth(store, c.Secret), validate)
	h := httpapi.NewShowtime(service.NewShowtime(store), validate)
	app := fiber.New()
	app.Post("/login", auth.Login)
	shows := app.Group("/jadwal-tayang", httpapi.RequireToken(c.Secret))
	shows.Post("/", h.Create)
	shows.Get("/", h.List)
	shows.Get("/:id", h.Find)
	shows.Put("/:id", h.Update)
	shows.Delete("/:id", h.Delete)
	return app, nil
}
