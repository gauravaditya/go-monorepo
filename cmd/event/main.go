package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/gauravaditya/go-monorepo/internal/event"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8081", "port to listen on")
	flag.Parse()
	slog.Info("Starting event service", "port", port)
	event.LoadConfig()
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	event.RegisterRoutes(app)
	app.Listen(fmt.Sprintf(":%s", port))
}
