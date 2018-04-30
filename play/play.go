package play

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	db "github.com/iopipe/eva/data"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func PlayEvent(invocation *db.InvocationRequest) db.InvocationId {
	result := &db.InvocationLog{
		InvocationRequest: *invocation,
	}
	var err error
	var responseEvent []byte = []byte("")

	lambdaEventBytes, err := db.GetEventJson(invocation.EventId)
	if err != nil {
		log.Fatal(err)
	}
	lambdaEvent := string(lambdaEventBytes)

	if invocation.PipeFile == "-" || (!invocation.PlayQuiet && invocation.PipeFile == "") {
		fmt.Println(lambdaEvent)
	}
	if invocation.PipeExec != "" {
		cmd := exec.Command("bash", "-c", invocation.PipeExec)
		cmd.Stdin = strings.NewReader(lambdaEvent)
		/* Capture invocations from IOpipe libraries */
		cmd.Env = append(os.Environ(),
			"IOPIPE_DEBUG=true")
		responseEvent, err = cmd.Output()
		if err != nil {
			log.Fatal("Error executing command.\nError: ", err)
		}
	}
	if invocation.AwsLambdaArn != "" {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		client := lambda.New(sess, &aws.Config{})
		lambdaResult, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String(invocation.AwsLambdaArn), Payload: []byte(lambdaEvent)})

		if err != nil {
			log.Fatal("Error calling lambda: ", err)
		}

		responseEvent = lambdaResult.Payload
	}
	if bytes.HasPrefix(responseEvent, []byte("IOPIPE-DEBUG:")) {
		endIOpipe := bytes.Index(responseEvent, []byte("\n"))
		var statData map[string]interface{}
		json.Unmarshal(responseEvent[13:endIOpipe], &statData)
		result.StatId = db.PutStat(statData)
	}
	if invocation.ResponseFile == "-" || (!invocation.PlayQuiet && invocation.ResponseFile == "") {
		fmt.Println(string(responseEvent))
	}

	var responseData map[string]interface{}
	if bytes.HasPrefix(responseEvent, []byte("IOPIPE-DEBUG:")) {
		endIOpipe := bytes.Index(responseEvent, []byte("\n"))
		if err = json.Unmarshal(responseEvent[13:endIOpipe], &responseData); err == nil {
			result.ResponseId = db.PutResponse(responseData)
		}
	} else {
		if err = json.Unmarshal(responseEvent, &responseData); err == nil {
			result.ResponseId = db.PutResponse(responseData)
		}
	}
	return db.PutInvocation(*result)
}
