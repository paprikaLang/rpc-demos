package main

import (
	"go-spider/distributed/persists/client"
	"go-spider/engine"
	"go-spider/main/parser"
	"go-spider/scheduler"
)

func main() {
	itemsaver, err := client.ItemSaver(":2333")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.GoroutineScheduler{},
		Workers:   10,
		ItemSaver: itemsaver,
	}
	// e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Path:       "https://laravelcollections.com",
		ParserFunc: parser.ParseTopic,
	})
}
