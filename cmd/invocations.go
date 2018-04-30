package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

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

var invocationsInspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect invocation",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//var results []map[string]interface{}
		var prefix = ""
		if len(args) > 1 {
			fmt.Println("[")
			prefix = " "
		}
		for arg := range args {
			invocationId, err := strconv.Atoi(args[arg])
			if err != nil {
				log.Fatal(err)
			}
			event := db.GetInvocation(db.InvocationId(invocationId))

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
	rootCmd.AddCommand(invocationsCmd)
	invocationsCmd.AddCommand(invocationsInspectCmd)
}
