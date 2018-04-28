package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	db "github.com/iopipe/eva/data"
	"github.com/spf13/cobra"
)

var cmdRequestbin = &cobra.Command{
	Use:   "requestbin",
	Short: "Return a URL to send HTTP(S) request to.",
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := http.Get("https://request.lol/geturl"); err == nil {
			defer resp.Body.Close()
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				fmt.Println(string(body))
				var event map[string]interface{}
				json.Unmarshal(body, &event)
				db.PutEvent(event)
			}
		}
	},
}

func init() {
	generateCmd.AddCommand(cmdRequestbin)
}
