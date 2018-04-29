package cmd

import (
	"github.com/iopipe/eva/listener"
	"github.com/iopipe/eva/pkg/templates"
	"github.com/spf13/cobra"
)

var cmdFlagHTTPListenerAddress string

var cmdListenHTTP = &cobra.Command{
	Use:   "daemon",
	Short: "Run HTTP daemon for listening to events.",
}

var cmdListenHTTPCloudfront = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate cloudfront requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		listener.Listen(templates.HandleCloudfrontEvent, templates.HandleCloudfrontResponse, cmdFlagHTTPListenerAddress, cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda)
	},
}

var cmdListenHTTPApiGw = &cobra.Command{
	Use:   "apigw",
	Short: "Generate apigw requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		listener.Listen(templates.HandleApiGwEvent, templates.HandleApiGwResponse, cmdFlagHTTPListenerAddress, cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda)
	},
}

var cmdListenHTTPInvocation = &cobra.Command{
	Use:   "invocations",
	Short: "Consume invocation messages from IOpipe library",
	Run: func(cmd *cobra.Command, args []string) {
		listener.Listen(templates.HandleInvocationEvent, templates.HandleInvocationResponse, cmdFlagHTTPListenerAddress, cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda)
	},
}

func init() {
	SetPlayFlags(cmdListenHTTP)
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerAddress, "addr", "a", ":8080", "HTTP(s) address to listen on.")

	rootCmd.AddCommand(cmdListenHTTP)
	cmdListenHTTP.AddCommand(cmdListenHTTPApiGw)
	cmdListenHTTP.AddCommand(cmdListenHTTPCloudfront)
	cmdListenHTTP.AddCommand(cmdListenHTTPInvocation)
}
