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
		var results []db.InvocationLog
		for arg := range args {
			invocationId, err := strconv.Atoi(args[arg])
			if err != nil {
				log.Fatal(err)
			}
			event := db.GetInvocation(db.InvocationId(invocationId))
			results = append(results, *event)
		}
		encoded, err := json.MarshalIndent(results, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(encoded))
	},
}

var invocationsStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "stats for invocation",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Stats:")
			queryResult := db.GetStats()

			// Query result are document IDs
			for id := range queryResult {
				fmt.Println(id)
			}
			return
		}

		var results []map[string]interface{}
		for arg := range args {
			invocationId, err := strconv.Atoi(args[arg])
			if err != nil {
				log.Fatal(err)
			}
			event := *db.GetInvocation(db.InvocationId(invocationId))

			stats := db.GetStat(event.StatId)
			results = append(results, stats)
		}
		encoded, err := json.MarshalIndent(results, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(encoded))
	},
}

func init() {
	rootCmd.AddCommand(invocationsCmd)
	invocationsCmd.AddCommand(invocationsInspectCmd)
	invocationsCmd.AddCommand(invocationsStatsCmd)
}
