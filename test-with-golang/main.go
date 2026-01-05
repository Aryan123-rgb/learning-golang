package main

import "strings"

var Greetings = map[string]string{
	"english":  "Hello",
	"french":   "Bonjour",
	"spanish":  "Hola",
	"german":   "Guten Tag",
	"italian":  "Ciao",
	"japanese": "こんにちは",
	"chinese":  "你好",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	_, greet := Greetings[strings.ToLower(language)]
	if !greet {
		return "Invalid Language"
	}
	return Greetings[strings.ToLower(language)] + " " + name  
}

func main() {
}
