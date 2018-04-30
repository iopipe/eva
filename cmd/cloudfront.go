package cmd

import (
	db "github.com/iopipe/eva/data"
	"github.com/iopipe/eva/pkg/templates"
	"github.com/iopipe/eva/play"
	"github.com/spf13/cobra"
)

var cloudfrontCmd = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate a cloudfront event.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		request, _ := CliParseHTTP(cmd, args)
		event := templates.HandleCloudfrontEvent(request)
		eventId := db.PutEvent(event)
		invocation := playArgsToInvocation(cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda, cmdFlagPlayQuiet)
		invocation.EventId = eventId
		play.PlayEvent(invocation)
	},
}

func init() {
	SetHttpCobraFlags(cloudfrontCmd)
	SetPlayFlags(cloudfrontCmd)
	generateCmd.AddCommand(cloudfrontCmd)
}
