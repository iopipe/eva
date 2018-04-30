package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	db "github.com/iopipe/eva/data"
	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect an event history record.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prefix := ""
		if len(args) > 1 {
			fmt.Println("[")
			prefix = " "
		}
		for arg := range args {
			eventIdStr := args[arg]
			eventId, err := strconv.Atoi(eventIdStr)
			if err != nil {
				log.Fatal(err)
			}
			event := db.GetEvent(db.EventId(eventId))

			encoded, err := json.MarshalIndent(event, prefix, " ")
			if err != nil {
				log.Fatal(err)
			}
			if len(args) > 1 {
				fmt.Print(prefix)
			}
			fmt.Print(string(encoded))
			if len(args)-1 != arg {
				fmt.Print(",")
			}
			fmt.Print("\n")
		}
		if len(args) > 1 {
			fmt.Println("]")
		}
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
