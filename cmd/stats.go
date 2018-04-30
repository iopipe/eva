package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	db "github.com/iopipe/eva/data"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "List captured invocation stats",
	Run: func(cmd *cobra.Command, args []string) {
		queryResult := db.GetStats()

		// Query result are document IDs
		for id := range queryResult {
			fmt.Println(id)
		}
	},
}

var inspectStatsCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect stats record.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var results []map[string]interface{}
		for arg := range args {
			statId, err := strconv.Atoi(args[arg])
			if err != nil {
				log.Fatal(err)
			}
			stat := db.GetStat(db.StatId(statId))
			results = append(results, stat)
		}
		encoded, err := json.MarshalIndent(results, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(encoded))
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.AddCommand(inspectStatsCmd)
}
