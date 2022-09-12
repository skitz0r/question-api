package main

import "errors"

// My thought here is that this might talk to a DB or just naievely progress the User to the next question
// and roll back later if needed so we don't have frequent DB reads

var TendoScript Script = Script{
	Id:          "0",
	QuestionIds: []string{"0", "1", "2"},
	// TODO can I bake a lambda in here to figure out the next question to ask?
}

type ScriptIterator interface {
	nextIndex(answer Answer) (string, error)
}

func nextIndex(script Script, answer Answer) (int, error) {
	//TODO validate answer matches expected format and repeat if bad

	for index, id := range script.QuestionIds {
		if answer.QuestionId == id {
			return index + 1, nil
		}
	}

	// i don't know what to do here to avoid a magic number
	return -1, errors.New("no more questions to ask")
}

// Grabs the next question
func NextQuestionForUser(answer Answer) (*Question, error) {
	script := getScript(ScriptServiceSingle, answer.ScriptId)
	nextId, err := nextIndex(script, answer) // this isn't picking up the member script#nextIndex method, why?
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
		TendoScript.Id: TendoScript,
	}
	service := &ScriptService{
		ScriptMap: idToScript,
	}
	return service
}

// TODO is there a way to mark this as the API of this object? seems not very go-like
// TODO how to handle map misses? looks like you get a "zero value" of the Script type?
func getScript(service ScriptService, scriptId string) Script {
	return service.ScriptMap[scriptId]
}
