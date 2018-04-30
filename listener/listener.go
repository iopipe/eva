package listener

import (
	"log"
	"net/http"

	db "github.com/iopipe/eva/data"
	"github.com/iopipe/eva/pkg/templates"
	"github.com/iopipe/eva/play"
)

func HTTPHandlerFactory(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, invocation *db.InvocationRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		requestEvent := requestHandler(req)
		eventId := db.PutEvent(requestEvent)
		invocation.EventId = eventId
		invocationId := play.PlayEvent(invocation)
		result := db.GetInvocation(invocationId)
		if responseData, err := db.GetResponseJson(result.ResponseId); err == nil {
			responseHandler(responseData, w)
		} else {
			log.Fatal(err)
		}
	}
}

func Listen(requestHandler templates.RequestHandler, responseHandler templates.ResponseHandler, address, execCmd, pipeFile, responseFile, lambdaArn string, playQuiet bool) {
	invocation := &db.InvocationRequest{
		PipeExec:     execCmd,
		PipeFile:     pipeFile,
		ResponseFile: responseFile,
		AwsLambdaArn: lambdaArn,
		PlayQuiet:    playQuiet,
	}
	handler := HTTPHandlerFactory(requestHandler, responseHandler, invocation)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
