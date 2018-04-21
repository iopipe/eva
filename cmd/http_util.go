package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var cmdFlagHttpMethod string
var cmdFlagHttpHost string
var cmdFlagHttpUri string
var cmdFlagHttpAuthorization string
var cmdFlagHttpDatafile string

func CliParseHTTP(cmd *cobra.Command, args []string) (*http.Request, error) {
	var body io.Reader
	var err error
	if cmdFlagHttpDatafile == "-" {
		fmt.Println("Reading stdin")
		body = os.Stdin
	} else if cmdFlagHttpDatafile != "" {
		body, err = os.Open(cmdFlagHttpDatafile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		body = nil //strings.NewReader("")
	}

	request := httptest.NewRequest(cmdFlagHttpMethod, args[0], body)
	request.Header.Add("Host", request.Host)
	return request, nil
}

func SetHttpCobraFlags(command *cobra.Command) {
	command.Flags().StringVarP(&cmdFlagHttpMethod, "method", "X", "GET", "HTTP Method")
	command.Flags().StringVarP(&cmdFlagHttpHost, "host", "H", "", "HTTP(s) host for event data.")
	command.Flags().StringVarP(&cmdFlagHttpUri, "path", "p", "", "HTTP(s) path or uri.")
	command.Flags().StringVarP(&cmdFlagHttpAuthorization, "auth", "A", "", "Authorization header")
	command.Flags().StringVarP(&cmdFlagHttpDatafile, "data", "d", "", "Data for body, or '-' for stdin.")
}
