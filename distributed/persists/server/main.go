package main

import (
	"go-spider/distributed/persists"
	"go-spider/distributed/rpcconfig"
	"log"

	"github.com/olivere/elastic"
)

func main() {
	log.Fatal(saveServer(":2333", "distributed"))
}

func saveServer(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcconfig.RPCServe(host, &persists.ItemSaverService{
		Client: client,
		Index:  index,
	})

}
