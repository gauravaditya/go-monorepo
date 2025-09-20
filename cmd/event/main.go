package main

import (
	"log/slog"
	"os"

	"github.com/gauravaditya/go-monorepo/internal/event"
	"github.com/gauravaditya/go-monorepo/pkg/clicmd"
	"github.com/gauravaditya/go-monorepo/pkg/server"
)

// @title Core Service API
// @version 1.0
// @description API documentation for the core service.
// @host localhost:8080
// @BasePath /
func main() {
	app := server.New("event-service")
	eventsvc := event.New(app)

	root := clicmd.NewRoot("event-service")
	root.AddCommand(clicmd.NewServer("event-service", eventsvc))

	if err := root.Execute(); err != nil {
		slog.Error("Failed to start event service", "error", err)
		os.Exit(1)
	}
}

// func main() {
// 	var port string
// 	flag.StringVar(&port, "port", "8081", "port to listen on")
// 	flag.Parse()
// 	slog.Info("Starting event service", "port", port)
// 	event.LoadConfig()
// 	app := fiber.New()
// 	app.Use(logger.New())
// 	app.Get("/health", func(c *fiber.Ctx) error {
// 		return c.SendString("OK")
// 	})
// 	event.RegisterRoutes(app)
// 	app.Listen(fmt.Sprintf(":%s", port))
// }
