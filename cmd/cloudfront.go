package cmd

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"os"

	db "github.com/iopipe/eva/data"
	"github.com/iopipe/eva/pkg/templates"
	"github.com/iopipe/eva/play"
	"github.com/spf13/cobra"
)

func mkCfEvent(handler templates.RequestHandler) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		request, _ := CliParseHTTP(cmd, args)
		event := handler(request)
		eventId := db.PutEvent(event)
		invocation := playArgsToInvocation(cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda, cmdFlagPlayQuiet)
		invocation.EventId = eventId
		play.PlayEvent(invocation)

		jsonEvent, err := json.Marshal(event)
		if err != nil {
			log.Fatal(err)
		}

		/* Emulate being an HTTP server sending a response... */
		w := httptest.NewRecorder()
		templates.HandleCloudfrontResponse(jsonEvent, w)
		result := w.Result()
		result.Write(os.Stdout)
	}
}

var cloudfrontRequestCmd = &cobra.Command{
	Use:   "cloudfront-request",
	Short: "Generate a cloudfront event.",
	Args:  cobra.ExactArgs(1),
	Run:   mkCfEvent(templates.HandleCloudfrontRequestEvent),
}

var cloudfrontResponseCmd = &cobra.Command{
	Use:   "cloudfront-response",
	Short: "Generate a cloudfront request event.",
	Args:  cobra.ExactArgs(1),
	Run:   mkCfEvent(templates.HandleCloudfrontResponseEvent),
}

func init() {
	SetHttpCobraFlags(cloudfrontRequestCmd)
	SetPlayFlags(cloudfrontRequestCmd)
	generateCmd.AddCommand(cloudfrontRequestCmd)

	SetHttpCobraFlags(cloudfrontResponseCmd)
	SetPlayFlags(cloudfrontResponseCmd)
	generateCmd.AddCommand(cloudfrontResponseCmd)
}
