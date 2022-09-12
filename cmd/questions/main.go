package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var userAnswer Answer = parseJson(request.Body)
	return events.APIGatewayProxyResponse{
		Body:       getNextQuestionBody(userAnswer),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

func parseJson(requestBody string) Answer {
	var args Answer
	json.Unmarshal([]byte(requestBody), &args)
	return args
}

//TODO figure out how to put this into a model package... or if that is a good practice here
// these are the classes/types described in the doc diagram

// next implement question routing and figure out the response object shape and how to marshall it
// also figure out how to import https://github.com/samply/golang-fhir-models to the project

// This function handles a QuestionRequest and determines what Question the user should be asked
// next based on the Script.
func getNextQuestionBody(answer Answer) string {
	nextQuestion, qerr := NextQuestion(answer)

	if qerr != nil {
		return qerr.Error()
	}

	var response = QuestionWrapper{
		Question: *nextQuestion,
		ScriptId: answer.ScriptId,
	}

	var bytes, jerr = json.MarshalIndent(response, "", "\t")

	if jerr != nil {
		return qerr.Error()
	}

	return fmt.Sprint(string(bytes))
}
