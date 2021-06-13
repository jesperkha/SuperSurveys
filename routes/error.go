package routes

import (
	"log"
)


type RequestError struct {
	Type string
	Msg  string
}

var ErrorMessages = map[string]string{
	"400": "Bad request",
	"404": "Not found",
	"500": "Server error",
}

func debug(msg string) {
	log.Println(msg)
}
