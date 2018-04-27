package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/iopipe/eva/templates"
)

var cmdFlagHTTPListenerAddress string
var cmdFlagHTTPListenerPipeExec string
var cmdFlagHTTPListenerPipeFile string
var cmdFlagHTTPListenerResponseFile string

var cmdListenHTTP = &cobra.Command{
	Use:   "listen",
	Short: "Listen on HTTP for events.",
}

var cmdListenHTTPCloudfront = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate cloudfront requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		spawnHTTPListener(templates.HandleCloudfrontEvent, templates.HandleCloudfrontResponse, cmdFlagHTTPListenerAddress, cmdFlagHTTPListenerPipeExec, cmdFlagHTTPListenerPipeFile, cmdFlagHTTPListenerResponseFile)
	},
}

var cmdListenHTTPApiGw = &cobra.Command{
	Use:   "apigw",
	Short: "Generate apigw requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		spawnHTTPListener(templates.HandleApiGwEvent, templates.HandleApiGwResponse, cmdFlagHTTPListenerAddress, cmdFlagHTTPListenerPipeExec, cmdFlagHTTPListenerPipeFile, cmdFlagHTTPListenerResponseFile)
	},
}

func init() {
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerAddress, "addr", "a", ":8080", "HTTP(s) address to listen on.")
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeExec, "exec", "e", "", "Pipe events into specified shell command.")
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeFile, "request", "q", "", "Save request JSON into file.")
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeFile, "response", "s", "", "Save response JSON into file.")

	rootCmd.AddCommand(cmdListenHTTP)
	cmdListenHTTP.AddCommand(cmdListenHTTPApiGw)
	cmdListenHTTP.AddCommand(cmdListenHTTPCloudfront)
}

func HTTPHandlerFactory(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, pipeExec string, pipeFile string, responseFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var responseEvent []byte = []byte("")
		var err error
		lambdaEvent := requestHandler(req)
		if pipeFile == "-" {
			fmt.Println(lambdaEvent)
		}
		if pipeExec != "" {
			cmd := exec.Command("bash", "-c", pipeExec)
			cmd.Stdin = strings.NewReader(lambdaEvent)
			responseEvent, err = cmd.Output()
			if err != nil {
				log.Fatal("Error executing command.\nError: ", err)
			}
		}
		if responseFile == "-" {
			fmt.Println(responseEvent)
		}
		responseHandler(responseEvent, w)
	}
}

func spawnHTTPListener(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, address string, pipeExec string, pipeFile string, responseFile string) {
	handler := HTTPHandlerFactory(requestHandler, responseHandler, pipeExec, pipeFile, responseFile)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
