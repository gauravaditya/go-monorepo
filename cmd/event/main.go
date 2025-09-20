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

	cmd := clicmd.NewRootCmd(
		"event-service",
		clicmd.WithServerCmd(eventsvc),
		clicmd.WithVersionCmd(clicmd.Version),
	)

	if err := cmd.Execute(); err != nil {
		slog.Error("Failed to start event service", "error", err)
		os.Exit(1)
	}
}
