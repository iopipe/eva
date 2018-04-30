package cmd

import (
	"log"
	"strconv"

	db "github.com/iopipe/eva/data"
	"github.com/iopipe/eva/play"
	"github.com/spf13/cobra"
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

			invocation := playArgsToInvocation(cmdFlagPlayExecCmd, cmdFlagPlayPipeFile, cmdFlagPlayResponseFile, cmdFlagPlayExecLambda, cmdFlagPlayQuiet)
			invocation.EventId = db.EventId(eventId)
			play.PlayEvent(invocation)
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
