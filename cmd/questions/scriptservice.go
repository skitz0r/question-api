package main

import (
	"fmt"
)

var TestScript Script = Script{
	Id:          "test",
	QuestionIds: []string{Question0.Id, Question1.Id, Question2.Id},
	// TODO can I bake a lambda in here to figure out the next question dynamically? maybe an interface entry on Script
}

type ScriptIndexWalker interface {
	nextIndex(answer Answer) (string, error)
}

// TODO would rather have a map of QuestionId -> QuestionId for progression, this index stuff is not readable and buggy
func nextIndex(script Script, answer Answer) (int, error) {
	for index, id := range script.QuestionIds {
		if answer.QuestionId == id {
			if lastQuestion(index, script.QuestionIds) {
				// using a magic number here for the summary, this is bad but will have to learn more about go to avoid it
				return -1, nil
			} else {
				return index + 1, nil
			}
		}
	}
	return 0, fmt.Errorf("this question wasn't recognized as part of the script")
}

func lastQuestion(questionIndex int, questions []string) bool {
	return questionIndex+1 == len(questions)
}

func NextQuestion(answer Answer) (*Question, error) {
	script, scripterr := getScript(answer.ScriptId)

	if scripterr != nil {
		return nil, scripterr
	}

	//TODO add another enpoint to start a conversation
	// empty request means serve back first question of the script until the above is done
	if answer.Payload == "" {
		return GetQuestion(script.QuestionIds[0])
	}

	askedQuestion, _ := GetQuestion(answer.QuestionId)
	expectedFormat := askedQuestion.AnswerFormat
	if !IsValid(expectedFormat, answer.Payload) {
		return askedQuestion, nil
	}

	nextId, err := nextIndex(*script, answer)

	if nextId == -1 {
		return &Summary, nil
	}

	if err != nil {
		return nil, err
	} else {
		return GetQuestion(script.QuestionIds[nextId])
	}
}

// Define a service and singleton to represent it... just experimenting here
var ScriptServiceSingle ScriptService = *newScriptService()

type ScriptService struct {
	ScriptMap map[string]Script
}

func newScriptService() *ScriptService {
	idToScript := map[string]Script{
		TestScript.Id: TestScript,
	}
	service := &ScriptService{
		ScriptMap: idToScript,
	}
	return service
}

func getScript(scriptId string) (*Script, error) {
	if script, ok := ScriptServiceSingle.ScriptMap[scriptId]; ok {
		return &script, nil
	}
	return nil, fmt.Errorf("couldn't find a script with id %s", scriptId)
}

// "Script.GetQuestions" if we wanted to include them in the summary
func _(script Script) []Question {
	var questions []Question
	for _, questionId := range script.QuestionIds {
		question, _ := GetQuestion(questionId)
		questions = append(questions, *question)
	}
	return questions
}
