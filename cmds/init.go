package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "usage",
	Short: "Usage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	initMakeEvent()
	initHTTPListener()
	rootCmd.AddCommand(cmdGetUrl)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
