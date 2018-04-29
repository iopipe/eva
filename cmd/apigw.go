package cmd

import (
	"github.com/iopipe/eva/pkg/templates"
	"github.com/iopipe/eva/play"
	"github.com/spf13/cobra"
)

var apigwCmd = &cobra.Command{
	Use:   "apigw",
	Short: "Generate an API Gw event.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		request, _ := CliParseHTTP(cmd, args)
		result := templates.HandleApiGwEvent(request)
		play.PlayEvent(result, cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda)
	},
}

func init() {
	SetHttpCobraFlags(apigwCmd)
	SetPlayFlags(apigwCmd)
	generateCmd.AddCommand(apigwCmd)
}
