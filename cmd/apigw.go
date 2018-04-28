package cmd

import (
	"fmt"
	"github.com/iopipe/eva/pkg/templates"
	"github.com/spf13/cobra"
)

var apigwCmd = &cobra.Command{
	Use:   "apigw",
	Short: "Generate an API Gw event.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		request, _ := CliParseHTTP(cmd, args)
		result := templates.HandleApiGwEvent(request)
		fmt.Println(result)
	},
}

func init() {
	SetHttpCobraFlags(apigwCmd)
	generateCmd.AddCommand(apigwCmd)
}
