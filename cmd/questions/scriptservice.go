package main

import (
	"fmt"
)

// My thought here is that this might talk to a DB or just progress the User to the next question
// and roll back later if needed so we don't have frequent DB reads

var TestScript Script = Script{
	Id:          "test",
	QuestionIds: []string{Question0.Id, Question1.Id, Question2.Id},
	// TODO can I bake a lambda in here to figure out the next question to ask?
}

type ScriptIndexWalker interface {
	nextIndex(answer Answer) (string, error)
}

// TODO this is very bug prone, would rather have a map of QuestionId -> QuestionId for progression
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
	script, scripterr := getScript(ScriptServiceSingle, answer.ScriptId)

	if scripterr != nil {
		return nil, scripterr
	}

	//TODO add another enpoint to start a conversation
	// empty request means serve back first question of the script until the above is done
	if answer.Payload == "" {
		return GetQuestion(script.QuestionIds[0])
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
var ScriptServiceSingle ScriptService = *NewScriptService()

type ScriptService struct {
	ScriptMap map[string]Script
}

func NewScriptService() *ScriptService {
	idToScript := map[string]Script{
		TestScript.Id: TestScript,
	}
	service := &ScriptService{
		ScriptMap: idToScript,
	}
	return service
}

// TODO is there a way to mark this as the API of this object? seems not very go-like
// TODO how to handle map misses? looks like you get a "zero value" of the Script type?
func getScript(service ScriptService, scriptId string) (*Script, error) {
	fmt.Print(scriptId)
	if script, ok := service.ScriptMap[scriptId]; ok {
		return &script, nil
	}
	return nil, fmt.Errorf("couldn't find a script with id %s", scriptId)
}
