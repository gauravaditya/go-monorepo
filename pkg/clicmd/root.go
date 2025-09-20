package clicmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Option func(*cobra.Command) *cobra.Command

func NewRootCmd(name string, options ...Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("%s Service CLI", name),
		Long:  fmt.Sprintf("%s Service API is a RESTful API.", name),
	}

	for _, opt := range options {
		cmd = opt(cmd)
	}

	return cmd
}

func WithServerCmd(runner Runner) Option {
	return func(root *cobra.Command) *cobra.Command {
		root.AddCommand(NewServer(root.Name(), runner))

		return root
	}
}

func WithVersionCmd(provider func() string) Option {
	return func(root *cobra.Command) *cobra.Command {
		root.AddCommand(NewVersionCommand(Version, root.OutOrStdout()))

		return root
	}
}
