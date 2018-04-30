package cmd

import (
	db "github.com/iopipe/eva/data"
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
		event := templates.HandleApiGwEvent(request)
		eventId := db.PutEvent(event)
		invocation := playArgsToInvocation(cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda, cmdFlagPlayQuiet)
		invocation.EventId = eventId
		play.PlayEvent(invocation)
	},
}

func init() {
	SetHttpCobraFlags(apigwCmd)
	SetPlayFlags(apigwCmd)
	generateCmd.AddCommand(apigwCmd)
}
