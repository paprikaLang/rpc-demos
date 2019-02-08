package client

import (
	"go-spider/distributed/rpcconfig"
	"go-spider/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcconfig.RPCClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item rpc saver: got item"+"#%id: %v", itemCount, item)
			itemCount++
			result := ""
			err := client.Call("ItemSaverService.Save", item, &result)
			if err != nil {
				log.Printf("item rpc service: saved item: %v, err: %v", item, err)
			}
		}
	}()
	return out, nil
}
