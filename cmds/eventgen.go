package cmds

import (
	"github.com/spf13/cobra"

	"github.com/iopipe/eva/templates"
)

var cmdFlagMakeEventHost string
var cmdFlagMakeEventUri string
var cmdFlagMakeEventAuthorization string

var cmdMakeEvent = &cobra.Command{
	Use:   "generate",
	Short: "Generate an event.",
	Run: func(cmd *cobra.Command, args []string) {
		templates.CreateCloudfrontEvent(cmdFlagMakeEventHost, cmdFlagMakeEventUri, cmdFlagMakeEventAuthorization)
	},
}

func initMakeEvent() {
	cmdMakeEvent.Flags().StringVarP(&cmdFlagMakeEventHost, "host", "H", "", "HTTP(s) host for event data.")
	cmdMakeEvent.Flags().StringVarP(&cmdFlagMakeEventUri, "path", "p", "", "HTTP(s) path or uri.")
	cmdMakeEvent.Flags().StringVarP(&cmdFlagMakeEventAuthorization, "auth", "u", "", "Authorization header")

	rootCmd.AddCommand(cmdMakeEvent)
}
