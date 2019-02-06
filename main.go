package main

import (
	"go-spider/engine"
	"go-spider/main/parser"
	"go-spider/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		Workers:   100,
	}
	e.Run(engine.Request{
		Url:        "https://laravelcollections.com",
		ParserFunc: parser.ParseTopic,
	})
}
