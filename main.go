package main

import (
	"go-spider/engine"
	"go-spider/main/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://laravelcollections.com",
		ParserFunc: parser.ParseTopic,
	})
}
