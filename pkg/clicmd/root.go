package clicmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

type Runner interface {
	Run(host, port string) error
}

func NewRoot(name string) *cobra.Command {
	return &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("%s Service CLI", name),
		Long:  fmt.Sprintf("%s Service API is a RESTful API.", name),
	}
}

func NewServer(service string, runner Runner) *cobra.Command {
	var host, port string

	cmd := &cobra.Command{
		Use:   "server",
		Short: fmt.Sprintf("Start %s service", service),
		Long:  fmt.Sprintf("Start %s service which provides RESTful APIs.", service),
		RunE: func(cmd *cobra.Command, args []string) error {
			slog.Info("Starting service...", "service", service)

			return runner.Run(host, port)
		},
	}

	cmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host to listen on")
	cmd.Flags().StringVar(&port, "port", "8080", "Port to listen on")

	return cmd
}
