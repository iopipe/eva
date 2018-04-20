package cmd

import (
	"fmt"
	"github.com/iopipe/eva/templates"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
)

var cmdFlagHttpHost string
var cmdFlagHttpUri string
var cmdFlagHttpAuthorization string

func CliParseHTTP(cmd *cobra.Command, args []string) *http.Request {
  request := &http.Request{
    Method: "GET",
    URL: &url.URL{
      Path: cmdFlagHttpUri,
    },
  }
  return request
}

var apigwCmd = &cobra.Command{
	Use:   "apigw",
	Short: "Generate an API Gw event.",
	Run: func (cmd *cobra.Command, args []string) {
   request := CliParseHTTP(cmd, args)
   result := templates.CreateApiGwEvent(request)
   fmt.Println(result)
  },
}

func SetHttpCobraFlags (command *cobra.Command) {
	command.Flags().StringVarP(&cmdFlagHttpHost, "host", "H", "", "HTTP(s) host for event data.")
	command.Flags().StringVarP(&cmdFlagHttpUri, "path", "p", "", "HTTP(s) path or uri.")
	command.Flags().StringVarP(&cmdFlagHttpAuthorization, "auth", "A", "", "Authorization header")
}

func init() {
  SetHttpCobraFlags(apigwCmd)
	generateCmd.AddCommand(apigwCmd)
}
