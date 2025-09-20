package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gauravaditya/go-monorepo/api"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

type App struct {
	*fiber.App
	cfg *Config
	err error
}

func New(fapp *fiber.App, cfg ...Config) *App {
	app := &App{
		App: fapp,
	}

	if len(cfg) != 0 {
		c := cfg[0]
		app.cfg = &c
	} else {
		app.LoadConfig()
	}

	return app
}

func (app *App) hasError() bool {
	return app.err != nil
}

func (app *App) Run(host, port string) error {
	app.LoadConfig()
	app.RegisterRoutes()
	// Start Kafka consumer in a goroutine
	ctx, cancel := context.WithCancel(context.Background())
	go app.startConsumer(ctx)

	// Graceful shutdown on SIGINT/SIGTERM
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		slog.Info("Shutting down consumer service")
		cancel()
		os.Exit(0)
	}()

	if app.hasError() {
		return app.err
	}

	return app.Listen(host + ":" + port)
}

func (app *App) startConsumer(ctx context.Context) {
	if app.hasError() {
		return
	}

	broker := fmt.Sprintf("%s:%d", app.cfg.KafkaHost, app.cfg.KafkaPort)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    app.cfg.KafkaTopic,
		GroupID:  "consumer-group-1",
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	defer r.Close()

	webhookSvc := NewWebhookService()
	slog.Info("ConsumerApp: started consuming", "broker", broker, "topic", app.cfg.KafkaTopic)
	attempt := 0
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			slog.Error("ConsumerApp: error reading message", "error", err)
			time.Sleep(backoff(attempt))
			attempt++
			continue
		}
		var event api.Event
		if err := json.Unmarshal(m.Value, &event); err != nil {
			slog.Error("ConsumerApp: failed to unmarshal event from Kafka", "error", err, "raw", string(m.Value))
			continue
		}
		slog.Info("ConsumerApp: event received", "event", event)
		event.Consumed = true
		payload, err := json.Marshal(event)
		if err != nil {
			slog.Error("ConsumerApp: failed to marshal event for webhook", "error", err)
			continue
		}
		if err := webhookSvc.Call(app.cfg.CoreWebhookURL, payload); err != nil {
			slog.Error("ConsumerApp: webhook call failed", "error", err, "webhook", app.cfg.CoreWebhookURL)
			continue
		}
		slog.Info("ConsumerApp: webhook call succeeded", "webhook", app.cfg.CoreWebhookURL)
	}
}

func backoff(attempt int) time.Duration {
	baseDelay := 1 * time.Second
	maxDelay := 10 * time.Second
	delay := time.Duration(attempt) * baseDelay
	if delay > maxDelay {
		delay = maxDelay
	}
	return delay
}
