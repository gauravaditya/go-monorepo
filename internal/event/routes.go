package event

import (
	"encoding/json"
	"log/slog"

	"github.com/gauravaditya/go-monorepo/api"
	_ "github.com/gauravaditya/go-monorepo/docs/event" // swag docs path
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func (app *App) RegisterRoutes() {
	// Swagger UI endpoint
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Post("/produce", app.produceHandler)

}

// produceHandler handles HTTP requests to produce an event.
// It parses the request body into an api.Event, marshals it to JSON,
// and sends it to the event producer. Returns appropriate HTTP status codes
// and error messages for invalid payloads, marshalling failures, or production errors.
// produceHandler handles the production of an event.
//
// @Summary      Produce an event
// @Description  Parses the incoming request body as an event, marshals it to JSON, and produces the event message.
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        event  body      api.Event  true  "Event payload"
// @Success      200    {object}  map[string]interface{}  "Event produced"
// @Failure      400    {object}  map[string]interface{}  "Invalid payload"
// @Failure      500    {object}  map[string]interface{}  "Failed to marshal event or produce event"
// @Router       /events/produce [post]
func (app *App) produceHandler(c *fiber.Ctx) error {
	var event api.Event
	if err := c.BodyParser(&event); err != nil {
		slog.Error("produceHandler: invalid payload", "error", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid payload"})
	}
	slog.Info("produceHandler: producing event", "event", event, "remote_addr", c.IP())
	payload, err := json.Marshal(event)
	if err != nil {
		slog.Error("produceHandler: failed to marshal event", "error", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to marshal event"})
	}
	if err := app.ProduceEventMessage(payload, event.Name); err != nil {
		slog.Error("produceHandler: failed to produce event", "error", err)
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	slog.Info("produceHandler: event produced", "event", event)
	return c.JSON(fiber.Map{"status": "event produced"})
}
