package cmds

import (
  "io"
  "log"
  "net/http"


	"github.com/spf13/cobra"
)

var cmdFlagHTTPListenerAddress string

var cmdListenHTTP = &cobra.Command{
	Use:   "listen",
	Short: "Listen on HTTP for events.",
	Run: func(cmd *cobra.Command, args []string) {
	  spawnHTTPListener(cmdFlagHTTPListenerAddress)
	},
}
func initHTTPListener() {
  cmdListenHTTP.Flags().StringVarP(&cmdFlagHTTPListenerAddress, "addr", "a", ":8080", "HTTP(s) address to listen on.")

  rootCmd.AddCommand(cmdListenHTTP)
}

func HTTPHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "hello, world!\n")
}

func spawnHTTPListener (address string) {
  http.HandleFunc("/", HTTPHandler)

  log.Fatal(http.ListenAndServe(address, nil))
}
