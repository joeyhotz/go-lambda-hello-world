package main

import (
	// aws context object in go
	"context"

	// lambda events in go
	"github.com/aws/aws-lambda-go/events"

	// lambda programming model in go
	"github.com/aws/aws-lambda-go/lambda"
)

// Lambda handler
func handleRequest(ctx context.Context, request *events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	return events.LambdaFunctionURLResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
