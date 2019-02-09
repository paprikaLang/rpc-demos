package main

import (
	"flag"
	itemsaver "go-spider/distributed/persists/client"
	"go-spider/distributed/rpcconfig"
	worker "go-spider/distributed/worker/client"
	"go-spider/engine"
	"go-spider/main/parser"
	"go-spider/scheduler"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemsaver, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

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

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcconfig.RPCClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected success to %s", h)
		} else {
			log.Printf("error connected to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
