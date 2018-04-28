package play

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func PlayEvent(lambdaEvent string, pipeExec string, pipeFile string, responseFile string) []byte {
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
