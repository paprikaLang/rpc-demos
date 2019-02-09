package main

import (
	"flag"
	"fmt"
	"go-spider/distributed/persists"
	"go-spider/distributed/rpcconfig"
	"log"

	"github.com/olivere/elastic"
)

var port = flag.Int("port", 0, "the port to listen ")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(saveServer(fmt.Sprintf(":%d", *port), "distributed"))
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
