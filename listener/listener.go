package listener

import (
	"log"
	"net/http"

	"github.com/iopipe/eva/play"
	"github.com/iopipe/eva/templates"
)

func HTTPHandlerFactory(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, pipeExec string, pipeFile string, responseFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		lambdaEvent := requestHandler(req)
		responseEvent := play.PlayEvent(lambdaEvent, pipeExec, pipeFile, responseFile)
		responseHandler(responseEvent, w)
	}
}

func Listen(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, address string, pipeExec string, pipeFile string, responseFile string) {
	handler := HTTPHandlerFactory(requestHandler, responseHandler, pipeExec, pipeFile, responseFile)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
