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
		var results []map[string]interface{}
		for arg := range args {
			eventId, err := strconv.Atoi(args[arg])
			if err != nil {
				log.Fatal(err)
			}
			event := db.GetEvent(db.EventId(eventId))
			results = append(results, event)
		}
		encoded, err := json.MarshalIndent(results, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(encoded))
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
