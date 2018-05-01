package cmd

import (
	db "github.com/iopipe/eva/data"
)

func playArgsToInvocation(execCmd, pipeFile, responseFile, lambdaArn string, playQuiet bool) *db.InvocationRequest {
	return &db.InvocationRequest{
		PipeExec:     execCmd,
		PipeFile:     pipeFile,
		ResponseFile: responseFile,
		AwsLambdaArn: lambdaArn,
		PlayQuiet:    playQuiet,
	}
}
