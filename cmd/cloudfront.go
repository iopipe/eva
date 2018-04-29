package cmd

import (
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
		result := templates.HandleCloudfrontEvent(request)
		play.PlayEvent(result, cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda, cmdFlagPlayQuiet)
	},
}

func init() {
	SetHttpCobraFlags(cloudfrontCmd)
	SetPlayFlags(cloudfrontCmd)
	generateCmd.AddCommand(cloudfrontCmd)
}
