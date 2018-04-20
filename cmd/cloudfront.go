package cmd

import (
	"github.com/iopipe/eva/templates"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
)

var cmdFlagMakeEventHost string
var cmdFlagMakeEventUri string
var cmdFlagMakeEventAuthorization string

var cloudfrontCmd = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate a cloudfront event.",
	Run: func(cmd *cobra.Command, args []string) {
		request := &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: cmdFlagMakeEventUri,
			},
		}
		templates.CreateCloudfrontEvent(request)
	},
}

func init() {
	cloudfrontCmd.Flags().StringVarP(&cmdFlagMakeEventHost, "host", "H", "", "HTTP(s) host for event data.")
	cloudfrontCmd.Flags().StringVarP(&cmdFlagMakeEventUri, "path", "p", "", "HTTP(s) path or uri.")
	cloudfrontCmd.Flags().StringVarP(&cmdFlagMakeEventAuthorization, "auth", "A", "", "Authorization header")

	generateCmd.AddCommand(cloudfrontCmd)
}
