package cmd

import (
	"fmt"

	db "github.com/iopipe/eva/data"
	"github.com/spf13/cobra"
)

// invocationsCmd represents the invocations command
var invocationsCmd = &cobra.Command{
	Use:   "invocations",
	Short: "Invocations list",
	Run: func(cmd *cobra.Command, args []string) {
		queryResult := db.GetInvocations()

		// Query result are document IDs
		for id := range queryResult {
			fmt.Println(id)
		}
	},
}

func init() {
	rootCmd.AddCommand(invocationsCmd)
}
