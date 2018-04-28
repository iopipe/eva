package cmd

import (
	"github.com/spf13/cobra"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Help for eva",
	Long: `Eva is a CLI application that enables developers
to work with events to store, replay, deliver,
and proxy. It is designed to work with event-driven
serverless systems.

Use "eva docs" to generate documentation and view docs.`,
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
