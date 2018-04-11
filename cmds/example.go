package cmds

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "os"

  "github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
  Use:   "geturl",
  Short: "Return a URL to send HTTP(S) request to.",
  Run: func(cmd *cobra.Command, args []string) {
    if resp, err := http.Get("https://request.lol/geturl"); (err == nil) {
      defer resp.Body.Close()
      if body, err := ioutil.ReadAll(resp.Body); (err == nil) {
        fmt.Println(string(body));
        // Do a thing.
      }
    }
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
