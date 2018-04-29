package cmd

import (
	"github.com/spf13/cobra"
)

var cmdFlagPlayDriver string
var cmdFlagPlayExecCmd string
var cmdFlagPlayExecLambda string
var cmdFlagPlayPipeFile string
var cmdFlagPlayResponseFile string

func SetPlayFlags(command *cobra.Command) {
	command.PersistentFlags().StringVarP(&cmdFlagPlayExecLambda, "lambda", "l", "", "Process event(s) with specified AWS Lambda ARN")
	command.PersistentFlags().StringVarP(&cmdFlagPlayExecCmd, "command", "c", "", "Pipe event(s) into specified shell command")
	command.PersistentFlags().StringVarP(&cmdFlagPlayPipeFile, "log-event", "e", "", "Log process event(s) into file, or - for stdout")
	command.PersistentFlags().StringVarP(&cmdFlagPlayResponseFile, "log-event-response", "E", "", "Log response event(s) into file, or - for stdout")
}
