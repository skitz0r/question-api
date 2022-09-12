package main

import "fmt"

var userAnswers map[string](Answer) = make(map[string](Answer))

func PersistAnswer(answer Answer) {
	userAnswers[generateKey(answer.UserId, answer.ScriptId, answer.QuestionId)] = answer
}

func GetAnswer(userId string, scriptId string, questionId string) Answer {
	return userAnswers[generateKey(userId, scriptId, questionId)]
}

func generateKey(userId string, scriptId string, questionId string) string {
	// just use hacky user+script+question concat key
	return fmt.Sprintf("%s_%s_%s", userId, scriptId, questionId)
}

func GetAnswers(userId string, scriptId string) []Answer {
	var answers []Answer
	script, _ := getScript(scriptId)
	for _, questionId := range script.QuestionIds {
		answers = append(answers, GetAnswer(userId, scriptId, questionId))
	}
	return answers
}
