package listener

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/iopipe/eva/templates"
)

func HandleEvent(lambdaEvent string, pipeExec string, pipeFile string, responseFile string) []byte {
	var responseEvent []byte = []byte("")
	var err error
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
	if responseFile == "-" || responseFile == "" {
		fmt.Println(responseEvent)
	}
	return responseEvent
}

func HTTPHandlerFactory(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, pipeExec string, pipeFile string, responseFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		lambdaEvent := requestHandler(req)
		responseEvent := HandleEvent(lambdaEvent, pipeExec, pipeFile, responseFile)
		responseHandler(responseEvent, w)
	}
}

func Listen(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, address string, pipeExec string, pipeFile string, responseFile string) {
	handler := HTTPHandlerFactory(requestHandler, responseHandler, pipeExec, pipeFile, responseFile)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
