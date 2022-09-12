package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var userAnswer Answer = parseJson(request.Body)

	// Good practice is probably to have this be a very simple implementation, should do persistence somewhere else.
	askedQuestion, _ := GetQuestion(userAnswer.QuestionId)
	expectedFormat := askedQuestion.AnswerFormat
	if IsValid(expectedFormat, userAnswer.Payload) {
		PersistAnswer(userAnswer)
	}

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

func getNextQuestionBody(answer Answer) string {

	// Get the next question, the summary, or ask the same thing again if we don't like their answer for some reason.
	nextQuestion, qerr := NextQuestion(answer)

	if qerr != nil {
		return qerr.Error()
	}

	var response = QuestionWrapper{
		Question:        *nextQuestion,
		ScriptId:        answer.ScriptId,
		PreviousAnswers: GetAnswers(answer.UserId, answer.ScriptId),
	}

	var bytes, jerr = json.MarshalIndent(response, "", "\t")

	if jerr != nil {
		return qerr.Error()
	}

	return fmt.Sprint(string(bytes))
}
