package cmd

import (
	"fmt"

	db "github.com/iopipe/eva/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
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
	rootCmd.AddCommand(listCmd)
}
