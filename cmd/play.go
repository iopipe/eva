package cmd

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	db "github.com/iopipe/eva/data"
	"github.com/iopipe/eva/play"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play event specified by id",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for arg := range args {
			eventIdStr := args[arg]
			eventId, err := strconv.Atoi(eventIdStr)
			if err != nil {
				log.Fatal(err)
			}

			event := db.GetEvent(eventId)
			encoded, err := json.MarshalIndent(event, "", " ")
			if err != nil {
				log.Fatal(err)
			}
			play.PlayEvent(string(encoded), cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda, cmdFlagPlayQuiet)
		}

		/* convert into HTTP...
		   responseEvent := listener.PlayEvent(...)
		   w := httptest.NewRecorder()
		   responseHandler(responseEvent, w)
		   result := w.Result()
		   result.Write(os.Stdout)*/
	},
}

func init() {
	SetPlayFlags(playCmd)
	rootCmd.AddCommand(playCmd)
}
