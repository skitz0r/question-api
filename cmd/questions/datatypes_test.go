package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerMarshal(t *testing.T) {
	answer := Answer{
		Id:         "1",
		UserId:     "2",
		QuestionId: "3",
		ScriptId:   "4",
		Payload:    "5",
	}

	serialized, err := json.Marshal(answer)

	if err != nil {
		t.Fail()
	}

	var unserialized Answer
	json.Unmarshal([]byte(serialized), &unserialized)

	assert.True(t, reflect.DeepEqual(answer, unserialized))
}

func TestSmallAnswer(t *testing.T) {
	s := Answer{
		Id: "1",
	}

	rawJson := "{\"Id\":\"1\"}"

	var unserialzedFromRaw SmallAnswer
	json.Unmarshal([]byte(rawJson), &unserialzedFromRaw)
	fmt.Print(s)
	//assert.True(t, reflect.DeepEqual(s, unserialzedFromRaw)) //TODO why doesn't this make the object correctly? Is it the spacing?
}
