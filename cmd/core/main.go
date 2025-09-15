package main

import (
	"flag"
	"fmt"
	"log/slog"

	_ "github.com/gauravaditya/go-monorepo/docs" // swag docs
	"github.com/gauravaditya/go-monorepo/internal/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title Core Service API
// @version 1.0
// @description API documentation for the core service.
// @host localhost:8080
// @BasePath /
func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.Parse()
	slog.Info("Starting core service", "port", port)
	core.LoadConfig()
	if err := core.InitDB(); err != nil {
		slog.Error("Failed to initialize database", "error", err)
		return
	}
	app := fiber.New()
	app.Use(logger.New())
	// Health endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	// Swagger UI endpoint
	app.Get("/swagger/*", swagger.HandlerDefault)
	core.RegisterRoutes(app)
	app.Listen(fmt.Sprintf(":%s", port))
}
