package main

import "errors"

// realistically these should be objects in a DB, but we're gonna spoof it with a [questionId -> Question] map
// and just hard code a bunch of nebulous question objects

// visible for testing
var Question0 = Question{
	Id:           "q0",
	Body:         "Hi [Patient First Name], on a scale of 1-10, would you recommend Dr [Doctor Last Name] to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend",
	AnswerFormat: IntegerFormat,
}

// TODO this question should branch if they say no and ask for clarification
var Question1 = Question{
	Id:           "q1",
	Body:         "Thank you. You were diagnosed with [Diagnosis]. Did Dr [Doctor Last Name] explain how to manage this diagnosis in a way you could understand?",
	AnswerFormat: BooleanFormat,
}

var Question2 = Question{
	Id:           "q2",
	Body:         "We appreciate the feedback, one last question: how do you feel about being diagnosed with [Diagnosis]?",
	AnswerFormat: TextFormat,
}

var Summary = Question{
	Id:   "summary",
	Body: "I am placeholder behavior for a summary because none of your answers were stored anywhere. Thanks for playing.",
}

func GetQuestion(questionId string) (*Question, error) {
	// If you squint really hard you can pretend this is like a mocked object db
	switch questionId {
	case Question0.Id:
		return &Question0, nil
	case Question1.Id:
		return &Question1, nil
	case Question2.Id:
		return &Question2, nil
	}
	return nil, errors.New("unknown question id")
}
