package cmd

import (
	"github.com/iopipe/eva/listener"
	"github.com/iopipe/eva/templates"
	"github.com/spf13/cobra"
)

var cmdFlagHTTPListenerAddress string
var cmdFlagHTTPListenerPipeExec string
var cmdFlagHTTPListenerPipeFile string
var cmdFlagHTTPListenerResponseFile string

var cmdListenHTTP = &cobra.Command{
	Use:   "daemon",
	Short: "Run HTTP daemon for listening to events.",
}

var cmdListenHTTPCloudfront = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate cloudfront requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		listener.Listen(templates.HandleCloudfrontEvent, templates.HandleCloudfrontResponse, cmdFlagHTTPListenerAddress, cmdFlagHTTPListenerPipeExec, cmdFlagHTTPListenerPipeFile, cmdFlagHTTPListenerResponseFile)
	},
}

var cmdListenHTTPApiGw = &cobra.Command{
	Use:   "apigw",
	Short: "Generate apigw requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		listener.Listen(templates.HandleApiGwEvent, templates.HandleApiGwResponse, cmdFlagHTTPListenerAddress, cmdFlagHTTPListenerPipeExec, cmdFlagHTTPListenerPipeFile, cmdFlagHTTPListenerResponseFile)
	},
}

var cmdListenHTTPInvocation = &cobra.Command{
	Use:   "invocations",
	Short: "Consume invocation messages from IOpipe library",
	Run: func(cmd *cobra.Command, args []string) {
		listener.Listen(templates.HandleInvocationEvent, templates.HandleInvocationResponse, cmdFlagHTTPListenerAddress, cmdFlagHTTPListenerPipeExec, cmdFlagHTTPListenerPipeFile, cmdFlagHTTPListenerResponseFile)
	},
}

func init() {
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerAddress, "addr", "a", ":8080", "HTTP(s) address to listen on.")
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeExec, "exec", "e", "", "Pipe events into specified shell command.")
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeFile, "request", "q", "", "Save request JSON into file.")
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeFile, "response", "s", "", "Save response JSON into file.")

	rootCmd.AddCommand(cmdListenHTTP)
	cmdListenHTTP.AddCommand(cmdListenHTTPApiGw)
	cmdListenHTTP.AddCommand(cmdListenHTTPCloudfront)
	cmdListenHTTP.AddCommand(cmdListenHTTPInvocation)
}
