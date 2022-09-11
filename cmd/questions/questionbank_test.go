package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	expected string
	actual   string
}

func TestGetQuestion1(t *testing.T) {
	testCase := TestCase{
		expected: Question1.Body,
	}

	returnedQuestion, err := GetQuestion("1")

	if err != nil {
		t.Fail()
	} else {
		testCase.actual = returnedQuestion.Body
	}

	assert.Equal(t, testCase.actual, testCase.expected)
}

func TestGetQuestionError(t *testing.T) {
	_, err := GetQuestion("99")
	assert.NotNil(t, err)
}
