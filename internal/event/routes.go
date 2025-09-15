package event

import (
	"encoding/json"
	"log/slog"

	"github.com/gauravaditya/go-monorepo/api"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/produce", produceHandler)
}

func produceHandler(c *fiber.Ctx) error {
       var event api.Event
       if err := c.BodyParser(&event); err != nil {
	       slog.Error("produceHandler: invalid payload", "error", err)
	       return c.Status(400).JSON(fiber.Map{"error": "Invalid payload"})
       }
       slog.Info("produceHandler: producing event", "event", event, "remote_addr", c.IP())
       // Publish to Kafka as JSON
       payload, err := json.Marshal(event)
       if err != nil {
	       slog.Error("produceHandler: failed to marshal event", "error", err)
	       return c.Status(500).JSON(fiber.Map{"error": "Failed to marshal event"})
       }
       if err := ProduceEventMessage(payload, event.Name); err != nil {
	       slog.Error("produceHandler: failed to produce event", "error", err)
	       return c.Status(500).JSON(fiber.Map{"error": err.Error()})
       }
       slog.Info("produceHandler: event produced", "event", event)
       return c.JSON(fiber.Map{"status": "event produced"})
}
