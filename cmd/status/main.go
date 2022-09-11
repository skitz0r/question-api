package status

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	mappedArguments := parseJson(request.Body)
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello user %s please complete script %s", mappedArguments.Id, "100"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

func parseJson(requestBody string) StatusRequest {
	var args StatusRequest
	json.Unmarshal([]byte(requestBody), &args)
	return args
}

type StatusRequest struct {
	Id string
}
