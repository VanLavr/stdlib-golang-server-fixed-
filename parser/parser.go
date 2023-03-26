package parser

import (
	"io/ioutil"
	"log"
)

func ParseFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("cannot parse HTML (in /crud/parser/parser.go 11:0) + %v", err.Error())
		return "error"
	}

	return string(content)
}