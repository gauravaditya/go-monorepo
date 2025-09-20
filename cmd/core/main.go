package main

import (
	"log/slog"
	"os"

	// _ "github.com/gauravaditya/go-monorepo/docs" // swag docs
	"github.com/gauravaditya/go-monorepo/internal/core"
	"github.com/gauravaditya/go-monorepo/pkg/clicmd"
	"github.com/gauravaditya/go-monorepo/pkg/server"
)

// @title Core Service API
// @version 1.0
// @description API documentation for the core service.
// @host localhost:8080
// @BasePath /
func main() {
	app := server.New("core-service")
	coresvc := core.New(app)

	root := clicmd.NewRoot("core-service")
	root.AddCommand(clicmd.NewServer("core-service", coresvc))

	if err := root.Execute(); err != nil {
		slog.Error("Failed to start core service", "error", err)
		os.Exit(1)
	}
}
