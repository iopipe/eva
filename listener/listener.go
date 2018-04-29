package listener

import (
	"log"
	"net/http"

	"github.com/iopipe/eva/pkg/templates"
	"github.com/iopipe/eva/play"
)

func HTTPHandlerFactory(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, execCmd, pipeFile, responseFile, lambdaArn string, playQuiet bool) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		lambdaEvent := requestHandler(req)
		responseEvent := play.PlayEvent(lambdaEvent, execCmd, pipeFile, responseFile, lambdaArn, playQuiet)
		responseHandler(responseEvent, w)
	}
}

func Listen(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, address, execCmd, pipeFile, responseFile, lambdaArn string, playQuiet bool) {
	handler := HTTPHandlerFactory(requestHandler, responseHandler, execCmd, pipeFile, responseFile, lambdaArn, playQuiet)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
