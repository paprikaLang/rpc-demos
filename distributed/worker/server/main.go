package main

import (
	"go-spider/distributed/rpcconfig"
	"go-spider/distributed/worker"
)

func main() {
	rpcconfig.RPCServe(":2334", worker.WorkerService{})
}
