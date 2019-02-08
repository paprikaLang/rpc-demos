package main

import (
	"go-spider/distributed/rpcconfig"
	"go-spider/engine"
	"go-spider/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":2333"
	go saveServer(host, "distributed")
	time.Sleep(time.Second)
	client, err := rpcconfig.RPCClient(host)
	if err != nil {
		panic(err)
	}
	item := engine.Item{
		URL:  "https://laravelcollections.com/go/612",
		Type: "collections",
		Id:   "612",
		Payload: model.Profile{
			Topic:  "News",
			Domain: "github.com",
			Title:  "Laravel v5.7.25 Released &nbsp;",
		},
	}
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
