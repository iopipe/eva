package cmd

import (
	"fmt"
	"io"
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

var cmdListenHTTP = &cobra.Command{
	Use:   "listen",
	Short: "Listen on HTTP for events.",
}

var cmdListenHTTPCloudfront = &cobra.Command{
	Use:   "cloudfront",
	Short: "Generate cloudfront requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		spawnHTTPListener(templates.CreateCloudfrontEvent, cmdFlagHTTPListenerAddress, cmdFlagHTTPListenerPipeExec, cmdFlagHTTPListenerPipeFile)
	},
}

var cmdListenHTTPApiGw = &cobra.Command{
	Use:   "apigw",
	Short: "Generate apigw requests from HTTP listener",
	Run: func(cmd *cobra.Command, args []string) {
		spawnHTTPListener(templates.CreateApiGwEvent, cmdFlagHTTPListenerAddress, cmdFlagHTTPListenerPipeExec, cmdFlagHTTPListenerPipeFile)
	},
}

func init() {
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerAddress, "addr", "a", ":8080", "HTTP(s) address to listen on.")
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeExec, "exec", "e", "", "Pipe events into specified shell command.")
	cmdListenHTTP.PersistentFlags().StringVarP(&cmdFlagHTTPListenerPipeFile, "output", "o", "", "Pipe events into file.")

	rootCmd.AddCommand(cmdListenHTTP)
	cmdListenHTTP.AddCommand(cmdListenHTTPApiGw)
	cmdListenHTTP.AddCommand(cmdListenHTTPCloudfront)
}

func HTTPHandlerFactory(templateHandler templates.TemplateHandler, pipeExec string, pipeFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		lambdaEvent := templateHandler(req)
		if pipeFile == "-" {
			fmt.Println(lambdaEvent)
		}
		if pipeExec != "" {
			cmd := exec.Command("bash", "-c", pipeExec)
			cmd.Stdin = strings.NewReader(lambdaEvent)
			responseEvent, err := cmd.Output()
			if err != nil {
				panic(err)
			}
			io.WriteString(w, string(responseEvent))
		}
	}
}

func spawnHTTPListener(templateHandler templates.TemplateHandler, address string, pipeExec string, pipeFile string) {
	handler := HTTPHandlerFactory(templateHandler, pipeExec, pipeFile)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
