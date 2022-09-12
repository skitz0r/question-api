package main

import "strconv"

var BooleanFormat = AnswerFormat{
	Id: "bool",
}

var TextFormat = AnswerFormat{
	Id: "text",
}

var IntegerFormat = AnswerFormat{
	Id: "int",
}

type FormateValidate interface {
	IsValid(data string) bool
}

func IsValid(format AnswerFormat, data string) bool {
	switch format {
	case BooleanFormat:
		_, err := strconv.ParseBool(data)
		return err == nil
	case TextFormat:
		return true
	case IntegerFormat:
		_, err := strconv.ParseInt(data, 0, 32)
		return err == nil
	}
	return false
}
