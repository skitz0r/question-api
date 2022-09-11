package main

import "errors"

// realistically these should be objects in a DB, but we're gonna spoof it with a [questionId -> Question] map
// and just hard code a bunch of nebulous question objects

// visible for testing
var Question0 = Question{
	Id:           "0",
	Body:         "Hi [Patient First Name], on a scale of 1-10, would you recommend Dr [Doctor Last Name] to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend",
	AnswerFormat: IntegerFormat,
}

// TODO this question should branch if they say no and ask for clarification
var Question1 = Question{
	Id:           "1",
	Body:         "Thank you. You were diagnosed with [Diagnosis]. Did Dr [Doctor Last Name] explain how to manage this diagnosis in a way you could understand?",
	AnswerFormat: BooleanFormat,
}

var Question2 = Question{
	Id:           "2",
	Body:         "We appreciate the feedback, one last question: how do you feel about being diagnosed with [Diagnosis]?",
	AnswerFormat: TextFormat,
}

func GetQuestion(questionId string) (*Question, error) {
	switch questionId {
	case "0":
		return &Question0, nil
	case "1":
		return &Question1, nil
	case "2":
		return &Question2, nil
	}
	return nil, errors.New("Unknown question ID")
}
