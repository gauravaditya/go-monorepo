package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gauravaditya/go-monorepo/internal/consumer"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8082", "port to listen on")
	flag.Parse()
	slog.Info("Starting consumer service", "port", port)
	consumer.LoadConfig()
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Start Kafka consumer in a goroutine
	ctx, cancel := context.WithCancel(context.Background())
	go consumer.StartConsumer(ctx, consumer.AppConfig.CoreWebhookURL, consumer.AppConfig.KafkaHost, consumer.AppConfig.KafkaPort)

	// Graceful shutdown on SIGINT/SIGTERM
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		slog.Info("Shutting down consumer service")
		cancel()
		os.Exit(0)
	}()

	app.Listen(fmt.Sprintf(":%s", port))
}
