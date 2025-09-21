package clicmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	// commitSHA is a constant representing the source version that
	// generated this build. It should be set during build via -ldflags.
	commitSHA string //nolint:gochecknoglobals

	// commitBranch represents branch the source version that
	// generated this build. It should be set during build via -ldflags.
	commitBranch string //nolint:gochecknoglobals

	// versionFromGit is a constant representing the version tag that
	// generated this build. It should be set during build via -ldflags.
	latestVersion string //nolint:gochecknoglobals
)

func NewVersionCommand(provider func() string, outStream io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		Long:  "Print the version information.",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprint(outStream, provider())

			return err
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
}

func Version() string {
	bytes, _ := json.Marshal(map[string]string{
		"version":   latestVersion,
		"branch":    commitBranch,
		"commitSHA": commitSHA,
	})

	return string(bytes)
}
