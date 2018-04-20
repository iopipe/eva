package cmd

import (
	"fmt"
	"github.com/iopipe/eva/templates"
	"github.com/spf13/cobra"
)

var apigwCmd = &cobra.Command{
	Use:   "apigw",
	Short: "Generate an API Gw event.",
	Run: func(cmd *cobra.Command, args []string) {
		request := CliParseHTTP(cmd, args)
		result := templates.CreateApiGwEvent(request)
		fmt.Println(result)
	},
}

func init() {
	SetHttpCobraFlags(apigwCmd)
	generateCmd.AddCommand(apigwCmd)
}
