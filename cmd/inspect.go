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
		fmt.Println(string(encoded))
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
