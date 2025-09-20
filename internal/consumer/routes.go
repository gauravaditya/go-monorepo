package consumer

import (
	_ "github.com/gauravaditya/go-monorepo/docs/consumer" // swag docs path
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func (app *App) RegisterRoutes() {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	// Swagger UI endpoint
	app.Get("/swagger/*", swagger.HandlerDefault)
}
