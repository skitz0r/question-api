package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	mappedArguments := parseJson(request.Body)
	return events.APIGatewayProxyResponse{
		Body:       mappedArguments.UserId + mappedArguments.Payload,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

func parseJson(body string) QuestionRequest {
	var args QuestionRequest
	json.Unmarshal([]byte(body), &args)
	return args
}

type QuestionRequest struct { // these members must be capitalized? why?
	UserId     string
	QuestionId string
	ScriptId   string
	Payload    string
}

//TODO figure out how to put this into a model package... or if that is a good practice here
// these are the classes/types described in the doc diagram

// next implement question routing and figure out the response object shape and how to marshall it
// also figure out how to import https://github.com/samply/golang-fhir-models to the project

// This function handles a QuestionRequest and determines what Question the user should be asked
// next based on the Script.
func routeQuestionRequests(request QuestionRequest) {

}
