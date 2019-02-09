package main

import (
	itemsaver "go-spider/distributed/persists/client"
	worker "go-spider/distributed/worker/client"
	"go-spider/engine"
	"go-spider/main/parser"
	"go-spider/scheduler"
)

func main() {
	itemsaver, err := itemsaver.ItemSaver(":2333")
	if err != nil {
		panic(err)
	}
	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:       &scheduler.GoroutineScheduler{},
		Workers:         10,
		ItemSaver:       itemsaver,
		WorkerProcessor: processor,
	}
	// e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Path:   "https://laravelcollections.com",
		Parser: engine.NewFuncParser(parser.ParseTopic, "ParseTopic"),
	})
}
