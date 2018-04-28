package main

import (
	"github.com/iopipe/eva/cmd"
)

func main() {
	/* Enter AWS Lambda mode if Lambda detected. */
	_, isLambda := os.LookupEnv("LAMBDA_TASK_ROOT")
	if isLambda {
		LambdaMain()
		return
	}
	cmd.Execute()
}
