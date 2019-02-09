package main

import (
	"fmt"
	"go-spider/distributed/rpcconfig"
	"go-spider/distributed/worker"
	"testing"
	"time"
)

func TestWorkerService(t *testing.T) {
	go rpcconfig.RPCServe(":2334", worker.WorkerService{})
	time.Sleep(time.Second)

	client, err := rpcconfig.RPCClient(":2334")
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		URL: "https://laravelcollections.com/audios",
		Parser: worker.SerializedParser{
			Name: "ItemParser",
			Args: "Audios",
		},
	}
	var result []worker.ParseResult
	err = client.Call("WorkerService.Process", req, &result)
	if err != nil {
		t.Errorf("error: %v", err)
	} else {
		fmt.Println(result)
	}
}
