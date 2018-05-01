package cmd

import (
	"fmt"

	db "github.com/iopipe/eva/data"
	"github.com/spf13/cobra"
)

// eventsCmd represents the events command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "List generated events",
	Run: func(cmd *cobra.Command, args []string) {
		queryResult := db.GetEvents()

		// Query result are document IDs
		for id := range queryResult {
			fmt.Println(id)
		}
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
}
