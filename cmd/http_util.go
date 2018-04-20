package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
	"net/http/httptest"
	"strings"
)

var cmdFlagHttpMethod string
var cmdFlagHttpHost string
var cmdFlagHttpUri string
var cmdFlagHttpAuthorization string

func CliParseHTTP(cmd *cobra.Command, args []string) (*http.Request, error) {
	body := strings.NewReader("Hello world")
	request := httptest.NewRequest(cmdFlagHttpMethod, args[0], body)
	request.Header.Add("Host", request.Host)
	return request, nil
}

func SetHttpCobraFlags(command *cobra.Command) {
	command.Flags().StringVarP(&cmdFlagHttpMethod, "method", "X", "GET", "HTTP Method")
	command.Flags().StringVarP(&cmdFlagHttpHost, "host", "H", "", "HTTP(s) host for event data.")
	command.Flags().StringVarP(&cmdFlagHttpUri, "path", "p", "", "HTTP(s) path or uri.")
	command.Flags().StringVarP(&cmdFlagHttpAuthorization, "auth", "A", "", "Authorization header")
}
