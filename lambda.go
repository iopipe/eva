package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(x string, e error) {
	/* No Op */
}

func LambdaMain(handler func(interface{}, error)) {
	lambda.Start(handler)
}
