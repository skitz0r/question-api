package main

// My thought here is that this might talk to a DB or just naievely progress the User to the next question
// and roll back later if needed so we don't have frequent DB reads

var TendoScript Script = Script{
	Id: "0",
}

// Grabs the next
func NextQuestionForUser(user User, script Script, question Question) {

}

type ScriptService struct {
	ScriptMap map[int]Script
}

func NewScriptService() ScriptService {
	return ScriptService{
		ScriptMap: make(map[int]Script),
	}
}
