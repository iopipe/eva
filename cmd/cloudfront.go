package cmd

import (
	"fmt"
	"github.com/iopipe/eva/templates"
	"github.com/spf13/cobra"
)

var cloudfrontCmd = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate a cloudfront event.",
	Run: func(cmd *cobra.Command, args []string) {
		request := CliParseHTTP(cmd, args)
		result := templates.CreateCloudfrontEvent(request)
		fmt.Println(result)
	},
}

func init() {
	SetHttpCobraFlags(cloudfrontCmd)
	generateCmd.AddCommand(cloudfrontCmd)
}
