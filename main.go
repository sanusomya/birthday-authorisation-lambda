package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// handleRequest is the main Lambda handler function.
func handleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {

	token := event.Headers["authorization"]
	secretValue := os.Getenv("validToken")
	fmt.Println(secretValue)

	if token == secretValue {
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: true}, nil
	}

	return events.APIGatewayV2CustomAuthorizerSimpleResponse{IsAuthorized: false}, errors.New("Unauthorized")

}

func main() {
	lambda.Start(handleRequest)

}
