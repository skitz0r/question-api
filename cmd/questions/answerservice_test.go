package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersistingAnAnswer(t *testing.T) {
	answer := Answer{
		Id:         "test-answer",
		UserId:     "test-user",
		QuestionId: Question0.Id,
		ScriptId:   TestScript.Id,
		Payload:    "10",
	}
	PersistAnswer(answer)
	fromMemory := GetAnswer(answer.UserId, answer.ScriptId, answer.QuestionId)
	assert.True(t, reflect.DeepEqual(answer, fromMemory))
}
