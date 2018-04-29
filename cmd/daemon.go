package cmd
import (
	"github.com/iopipe/eva/listener"
	"github.com/iopipe/eva/pkg/templates"
	"github.com/spf13/cobra"
)

var cmdFlagHTTPListenerAddress string

var cmdDaemon = &cobra.Command{
	Use:   "daemon",
	Short: "Run HTTP daemon for listening to events.",
}

var cmdDaemonCloudfront = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate cloudfront requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		listen(templates.HandleCloudfrontEvent, templates.HandleCloudfrontResponse)
	},
}

var cmdDaemonApiGw = &cobra.Command{
	Use:   "apigw",
	Short: "Generate apigw requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		listen(templates.HandleApiGwEvent, templates.HandleApiGwResponse)
	},
}

var cmdDaemonInvocation = &cobra.Command{
	Use:   "invocations",
	Short: "Consume invocation messages from IOpipe library",
	Run: func(cmd *cobra.Command, args []string) {
		listen(templates.HandleInvocationEvent, templates.HandleInvocationResponse)
	},
}

func listen(requestTemplate templates.RequestHandler, responseTemplate templates.ResponseHandler) {
	listener.Listen(requestTemplate, responseTemplate, cmdFlagHTTPListenerAddress, cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda)
}

func init() {
	SetPlayFlags(cmdDaemon)
	cmdDaemon.PersistentFlags().StringVarP(&cmdFlagHTTPListenerAddress, "addr", "a", ":8080", "HTTP(s) address to listen on.")

	rootCmd.AddCommand(cmdDaemon)
	cmdDaemon.AddCommand(cmdDaemonApiGw)
	cmdDaemon.AddCommand(cmdDaemonCloudfront)
	cmdDaemon.AddCommand(cmdDaemonInvocation)
}
