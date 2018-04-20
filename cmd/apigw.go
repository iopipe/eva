package cmd

import (
	"fmt"
	"github.com/iopipe/eva/templates"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
)

var cmdFlagMakeAPIGwEventHost string
var cmdFlagMakeAPIGwEventUri string
var cmdFlagMakeAPIGwEventAuthorization string

var apigwCmd = &cobra.Command{
	Use:   "apigw",
	Short: "Generate an API Gw event.",
	Run: func(cmd *cobra.Command, args []string) {
		request := &http.Request{
			Method: "GET",
			URL: &url.URL{
				Path: cmdFlagMakeAPIGwEventUri,
			},
		}
		result := templates.CreateApiGwEvent(request)
		fmt.Println(result)
	},
}

func init() {
	apigwCmd.Flags().StringVarP(&cmdFlagMakeAPIGwEventHost, "host", "H", "", "HTTP(s) host for event data.")
	apigwCmd.Flags().StringVarP(&cmdFlagMakeAPIGwEventUri, "path", "p", "", "HTTP(s) path or uri.")
	apigwCmd.Flags().StringVarP(&cmdFlagMakeAPIGwEventAuthorization, "auth", "A", "", "Authorization header")

	generateCmd.AddCommand(apigwCmd)
}
