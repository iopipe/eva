package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/iopipe/eva/templates"
)

var cmdFlagHTTPListenerAddress string

var cmdListenHTTP = &cobra.Command{
	Use:   "listen",
	Short: "Listen on HTTP for events.",
	Run: func(cmd *cobra.Command, args []string) {
		spawnHTTPListener(cmdFlagHTTPListenerAddress)
	},
}

func init() {
	cmdListenHTTP.Flags().StringVarP(&cmdFlagHTTPListenerAddress, "addr", "a", ":8080", "HTTP(s) address to listen on.")

	rootCmd.AddCommand(cmdListenHTTP)
}

func HTTPHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")

	lambdaEvent := templates.CreateCloudfrontEvent(req)
	fmt.Println(lambdaEvent)
}

func spawnHTTPListener(address string) {
	http.HandleFunc("/", HTTPHandler)

	log.Fatal(http.ListenAndServe(address, nil))
}
