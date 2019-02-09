package main

import (
	"go-spider/engine"
	"go-spider/main/parser"
	"go-spider/persist"
	"go-spider/scheduler"
)

func main() {
	itemsaver, err := persist.ItemSaver("laravel_collection")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.GoroutineScheduler{},
		Workers:   10,
		ItemSaver: itemsaver,
		WorkerProcessor: engine.Worker,
	}
	// e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Path:   "https://laravelcollections.com",
		Parser: engine.NewFuncParser(parser.ParseTopic, "ParseTopic"),
	})
}
