package cmd

import (
	"github.com/spf13/cobra"
)

var cmdFlagHTTPListenerAddress string
var cmdFlagHTTPListenerPipeExec string
var cmdFlagHTTPListenerPipeFile string
var cmdFlagHTTPListenerResponseFile string

func SetPlayFlags(command *cobra.Command) {
	command.PersistentFlags().StringVarP(&cmdFlagHTTPListenerAddress, "addr", "a", ":8080", "HTTP(s) address to listen on.")
	command.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeExec, "exec", "e", "", "Pipe events into specified shell command.")
	command.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeFile, "request", "q", "", "Save request JSON into file.")
	command.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeFile, "response", "s", "", "Save response JSON into file.")
}
