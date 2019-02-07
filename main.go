package main

import (
	"go-spider/engine"
	"go-spider/main/parser"
	"go-spider/persist"
	"go-spider/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.GoroutineScheduler{},
		Workers:   10,
		ItemSaver: persist.ItemSaver(),
	}
	// e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Path:       "https://laravelcollections.com",
		ParserFunc: parser.ParseTopic,
	})
}
