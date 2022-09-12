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
		QuestionId: "q0",
		ScriptId:   TestScript.Id,
		Payload:    "",
	}

	afterStarter, _ := NextQuestion(starter)
	assert.Equal(t, Question0.Id, afterStarter.Id)

	answer0 := Answer{
		Id:         "test-answer",
		UserId:     "test-user",
		QuestionId: "q0",
		ScriptId:   TestScript.Id,
		Payload:    "I am the response to q0",
	}

	// Expect that Question1 comes after Question0
	after0, _ := NextQuestion(answer0)
	assert.Equal(t, Question1.Id, after0.Id)

	answer1 := Answer{
		Id:         "test-answer",
		UserId:     "test-user",
		QuestionId: "q1",
		ScriptId:   TestScript.Id,
		Payload:    "I am a the response to q1",
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
		Payload:    "I am a the response to q2",
	}

	// Expect that Question2 comes after Question1
	after2, _ := NextQuestion(answer2)
	assert.Equal(t, Summary.Id, after2.Id)
}
