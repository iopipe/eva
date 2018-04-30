package data

type InvocationRequest struct {
	EventId      EventId
	PipeExec     string
	PipeFile     string
	ResponseFile string
	AwsLambdaArn string
	PlayQuiet    bool
}

type InvocationLog struct {
	StatId            StatId
	ResponseId        ResponseId
	InvocationRequest InvocationRequest
}
