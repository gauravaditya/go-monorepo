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

	cmd := clicmd.NewRootCmd(
		"core-service",
		clicmd.WithServerCmd(consumersvc),
		clicmd.WithVersionCmd(clicmd.Version),
	)

	if err := cmd.Execute(); err != nil {
		slog.Error("Failed to start consumer service", "error", err)
		os.Exit(1)
	}
}
