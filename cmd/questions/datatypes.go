package main

//TODO figure out how local go modules work
// I want to put this in its own module so that different modules can share these but couldn't get the local import to work...

type Question struct {
	Id           string
	Body         string
	AnswerFormat AnswerFormat
}

type Script struct {
	Id          string
	QuestionIds []string // This should be an ordered set but it seems like go just wants me to use Slices everywhere so lets roll with it?
}

type ScriptEntry struct {
	Question Question
}

type Conversation struct {
	Id string
}

type AnswerFormat struct {
	Id string
}

type Answer struct {
	Id         string
	UserId     string
	QuestionId string
	ScriptId   string
}

type User struct {
	Id string
}
