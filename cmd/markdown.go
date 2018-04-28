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
		const fmTemplate = `---
permalink: /
---

`

		filePrepender := func(filename string) string {
			if filename == "docs/eva.md" {
				return fmTemplate
			}
			return filename
		}

		err := doc.GenMarkdownTreeCustom(rootCmd, "./docs", filePrepender, func(s string) string { return s })
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	exportCmd.AddCommand(markdownCmd)
}
