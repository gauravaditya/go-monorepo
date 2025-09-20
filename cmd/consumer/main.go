package main

import (
	"log/slog"
	"os"

	"github.com/gauravaditya/go-monorepo/internal/consumer"
	"github.com/gauravaditya/go-monorepo/pkg/clicmd"
	"github.com/gauravaditya/go-monorepo/pkg/server"
)

func main() {
	app := server.New("consumer-service")
	consumersvc := consumer.New(app)

	root := clicmd.NewRoot("consumer-service")
	root.AddCommand(clicmd.NewServer("consumer-service", consumersvc))

	if err := root.Execute(); err != nil {
		slog.Error("Failed to start core service", "error", err)
		os.Exit(1)
	}
}
