package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTheScriptFromTheService(t *testing.T) {
	script := ScriptServiceSingle.ScriptMap[TestScript.Id]
	assert.Equal(t, TestScript.Id, script.Id)
}

func TestScriptProgression(t *testing.T) {
	starter := Answer{
		Id:         "test-answer",
		UserId:     "test-user",
		QuestionId: Question0.Id,
		ScriptId:   TestScript.Id,
		Payload:    "",
	}

	afterStarter, _ := NextQuestion(starter)
	assert.Equal(t, Question0.Id, afterStarter.Id)

	answer0 := Answer{
		Id:         "test-answer",
		UserId:     "test-user",
		QuestionId: Question0.Id,
		ScriptId:   TestScript.Id,
		Payload:    "10",
	}

	// Expect that Question1 comes after Question0
	after0, _ := NextQuestion(answer0)
	assert.Equal(t, Question1.Id, after0.Id)

	answer1 := Answer{
		Id:         "test-answer",
		UserId:     "test-user",
		QuestionId: "q1",
		ScriptId:   TestScript.Id,
		Payload:    "False",
	}

	// Expect that Question2 comes after Question1
	after1, _ := NextQuestion(answer1)
	assert.Equal(t, Question2.Id, after1.Id)

	// the summary should come last
	answer2 := Answer{
		Id:         "test-answer",
		UserId:     "test-user",
		QuestionId: Question2.Id,
		ScriptId:   TestScript.Id,
		Payload:    "I do not feel great about this diagnosis.",
	}

	// Expect that Question2 comes after Question1
	after2, _ := NextQuestion(answer2)
	assert.Equal(t, Summary.Id, after2.Id)
}

func TestBadInputLoops(t *testing.T) {
	answer1 := Answer{
		Id:         "test-answer",
		UserId:     "test-user",
		QuestionId: Question1.Id, // expect a bool
		ScriptId:   TestScript.Id,
		Payload:    "This is definitely not a boolean.",
	}

	// Expect that we ask question 1 again
	after1, _ := NextQuestion(answer1)
	assert.Equal(t, Question1.Id, after1.Id)

}
