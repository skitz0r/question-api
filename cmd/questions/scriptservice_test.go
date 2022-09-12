package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTheScriptFromTheService(t *testing.T) {
	script := ScriptServiceSingle.ScriptMap[TendoScript.Id]
	assert.Equal(t, TendoScript.Id, script.Id)
}

func TestScriptProgression(t *testing.T) {
	answer := Answer{
		Id:         "test-answer",
		UserId:     "test-user",
		QuestionId: Question0.Id,
		ScriptId:   TendoScript.Id,
	}

	// Expect that Question1 comes after Question0
	nextq, _ := NextQuestionForUser(answer)
	assert.Equal(t, "1", nextq.Id)
}
