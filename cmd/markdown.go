package cmd

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"

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
			if filename != "docs/eva.md" {
				return filename
			}
			now := time.Now().Format(time.RFC3339)
			name := filepath.Base(filename)
			base := strings.TrimSuffix(name, path.Ext(name))
			url := "/commands/" + strings.ToLower(base) + "/"
			return fmt.Sprintf(fmTemplate, now, strings.Replace(base, "_", " ", -1), base, url)
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
