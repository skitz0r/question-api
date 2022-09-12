package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("1")
	mappedArguments := parseJson(request.Body)
	fmt.Println("2")
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello user %s please complete script with id \"%s\"", mappedArguments.Id, "test"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

func parseJson(requestBody string) StatusRequest {
	var args StatusRequest
	json.Unmarshal([]byte(requestBody), &args)
	fmt.Print(args)
	return args
}

type StatusRequest struct {
	Id string
}
