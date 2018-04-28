package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// markdownCmd represents the markdown command
var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Export to markdown",
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	exportCmd.AddCommand(markdownCmd)
}
