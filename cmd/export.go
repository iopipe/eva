package cmd

import (
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export documentation to file or directory",
}

func init() {
	docsCmd.AddCommand(exportCmd)
}
