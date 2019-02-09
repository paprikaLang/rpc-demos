package main

import (
	"go-spider/distributed/rpcconfig"
	"go-spider/distributed/worker"
)

func main() {
	rpcconfig.RPCServe(":9000", worker.WorkerService{})
}
