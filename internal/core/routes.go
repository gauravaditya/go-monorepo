package core

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/gauravaditya/go-monorepo/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// RegisterRoutes registers all HTTP routes for the core service.
func (app *App) RegisterRoutes() {
	if app.hasError() {
		return
	}

	staticDir := filepath.Join("cmd", "core", "static")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4200",
	}))
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	// Swagger UI endpoint
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Static("/", staticDir)
	app.Post("/register", app.RegisterHandler)
	app.Post("/webhook", app.WebhookHandler)
	app.Get("/events-data", app.EventsDataHandler)
}

// RegisterHandler handles event registration
// @Summary Register events
// @Description Register a number of events (calls event service)
// @Tags events
// @Accept json
// @Produce json
// @Param count body object true "Number of events to register"
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func (app *App) RegisterHandler(c *fiber.Ctx) error {
	type req struct {
		Count int `json:"count"`
	}
	var r req
	if err := c.BodyParser(&r); err != nil {
		slog.Error("RegisterHandler: invalid payload", "error", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid payload"})
	}
	if r.Count < 1 {
		slog.Error("RegisterHandler: count must be >= 1", "count", r.Count)
		return c.Status(400).JSON(fiber.Map{"error": "Count must be >= 1"})
	}
	eventServiceURL := app.GetEventServiceURL()
	var events []api.Event
	for i := 0; i < r.Count; i++ {
		ev := api.Event{
			Name:      fmt.Sprintf("event-%d-%d", time.Now().UnixNano(), i),
			Timestamp: time.Now().Format(time.RFC3339),
			Consumed:  false,
		}
		// Save to DB
		if err := app.CreateEvent(ev.Name, ev.Timestamp); err != nil {
			slog.Error("RegisterHandler: failed to save event to DB", "error", err, "event", ev)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save event to DB"})
		}
		// Send to event service
		if err := app.CallEventServiceWithPayload(eventServiceURL, ev); err != nil {
			slog.Error("RegisterHandler: failed to call event service", "error", err, "event", ev)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to register event with event service"})
		}
		events = append(events, ev)
	}
	slog.Info("RegisterHandler: events registered via event service", "count", r.Count)
	return c.JSON(fiber.Map{"message": "Event registration submitted", "events": events})
}

// WebhookHandler handles webhook updates
// @Summary Webhook to update event consumed count
// @Description Webhook endpoint for consumer to update event consumption
// @Tags events
// @Accept json
// @Prodmake swaggeruce json
// @Param event body EventData true "Event data"
// @Success 200 {object} map[string]interface{}
// @Router /webhook [post]
func (app *App) WebhookHandler(c *fiber.Ctx) error {
	var event api.Event
	if err := c.BodyParser(&event); err != nil {
		slog.Error("WebhookHandler: invalid payload", "error", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid payload"})
	}
	if err := app.UpdateEventConsumed(event.Name, event.Timestamp, event.Consumed); err != nil {
		slog.Error("WebhookHandler: failed to update event", "error", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update event"})
	}
	slog.Info("WebhookHandler: event updated", "event", event)
	return c.JSON(fiber.Map{"status": "ok"})
}

// EventsDataHandler returns event consumption data
// @Summary Get event consumption data
// @Description Returns all event consumption data
// @Tags events
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /events-data [get]
func (app *App) EventsDataHandler(c *fiber.Ctx) error {
	events, err := app.GetAllEvents()
	if err != nil {
		slog.Error("EventsDataHandler: failed to fetch events", "error", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch events"})
	}
	slog.Info("EventsDataHandler called", "count", len(events))
	return c.JSON(fiber.Map{"events": events})
}
