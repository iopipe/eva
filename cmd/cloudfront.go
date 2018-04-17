package cmd

import (
	"github.com/iopipe/eva/templates"
	"github.com/spf13/cobra"
)

var cmdFlagMakeEventHost string
var cmdFlagMakeEventUri string
var cmdFlagMakeEventAuthorization string

var cloudfrontCmd = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate a cloudfront event.",
	Run: func(cmd *cobra.Command, args []string) {
		templates.CreateCloudfrontEvent(cmdFlagMakeEventHost, cmdFlagMakeEventUri, cmdFlagMakeEventAuthorization)
	},
}

func init() {
	cloudfrontCmd.Flags().StringVarP(&cmdFlagMakeEventHost, "host", "H", "", "HTTP(s) host for event data.")
	cloudfrontCmd.Flags().StringVarP(&cmdFlagMakeEventUri, "path", "p", "", "HTTP(s) path or uri.")
	cloudfrontCmd.Flags().StringVarP(&cmdFlagMakeEventAuthorization, "auth", "u", "", "Authorization header")

	generateCmd.AddCommand(cloudfrontCmd)
}
