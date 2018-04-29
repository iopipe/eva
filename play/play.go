package play

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func PlayEvent(lambdaEvent string, pipeExec string, pipeFile string, responseFile string, lambdaArn string) []byte {
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
	if lambdaArn != "" {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		client := lambda.New(sess, &aws.Config{})
		result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String(lambdaArn), Payload: []byte(lambdaEvent)})

		if err != nil {
			log.Fatal("Error calling lambda: ", err)
		}

		responseEvent = result.Payload
	}
	if responseFile == "-" || responseFile == "" {
		fmt.Println(string(responseEvent))
	}
	return responseEvent
}
