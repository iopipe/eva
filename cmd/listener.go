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
}

var cmdListenHTTPCloudfront = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate cloudfront requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		spawnHTTPListener(HTTPHandlerCloudfront, cmdFlagHTTPListenerAddress)
	},
}

var cmdListenHTTPApiGw = &cobra.Command{
	Use:   "apigw",
	Short: "Generate apigw requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		spawnHTTPListener(HTTPHandlerApiGw, cmdFlagHTTPListenerAddress)
	},
}

func init() {
	cmdListenHTTP.Flags().StringVarP(&cmdFlagHTTPListenerAddress, "addr", "a", ":8080", "HTTP(s) address to listen on.")

	rootCmd.AddCommand(cmdListenHTTP)
	cmdListenHTTP.AddCommand(cmdListenHTTPApiGw)
	cmdListenHTTP.AddCommand(cmdListenHTTPCloudfront)
}

func HTTPHandlerApiGw(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
	lambdaEvent := templates.CreateApiGwEvent(req)
	fmt.Println(lambdaEvent)
}

func HTTPHandlerCloudfront(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
	lambdaEvent := templates.CreateCloudfrontEvent(req)
	fmt.Println(lambdaEvent)
}

func spawnHTTPListener(handler func(w http.ResponseWriter, req *http.Request), address string) {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
