package main

import (
	"flag"
	"fmt"
	"go-spider/distributed/rpcconfig"
	"go-spider/distributed/worker"
	"log"
)
var port = flag.Int("port", 0, "the port to listen ")
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcconfig.RPCServe(fmt.Sprintf(":%d", *port), worker.WorkerService{}))
}
