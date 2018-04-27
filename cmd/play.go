package cmd

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	db "github.com/iopipe/eva/data"
	"github.com/iopipe/eva/listener"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play event specified by id",
	Run: func(cmd *cobra.Command, args []string) {
		eventIdStr := args[0]
		eventId, err := strconv.Atoi(eventIdStr)
		if err != nil {
			log.Fatal(err)
		}

		event := db.GetEvent(eventId)
		encoded, err := json.MarshalIndent(event, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		listener.HandleEvent(string(encoded), cmdFlagHTTPListenerPipeExec, cmdFlagHTTPListenerPipeFile, cmdFlagHTTPListenerResponseFile)

		/* convert into HTTP...
		   responseEvent := listener.HandleEvent(...)
		   w := httptest.NewRecorder()
		   responseHandler(responseEvent, w)
		   result := w.Result()
		   result.Write(os.Stdout)*/
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVarP(&cmdFlagHTTPListenerPipeExec, "exec", "e", "", "Pipe events into specified shell command.")
	playCmd.Flags().StringVarP(&cmdFlagHTTPListenerPipeFile, "request", "q", "", "Save request JSON into file.")
	playCmd.Flags().StringVarP(&cmdFlagHTTPListenerPipeFile, "response", "s", "", "Save response JSON into file.")
}
