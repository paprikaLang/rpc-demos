package client

import (
	"go-spider/distributed/worker"
	"go-spider/engine"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(req engine.Request) ([]engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sRes []worker.ParseResult
		c := <-clientChan
		err := c.Call("WorkerService.Process", sReq, &sRes)
		if err != nil {
			return []engine.ParseResult{}, err
		}
		// fmt.Println(sRes)
		return worker.DeserializeResult(sRes), nil
	}
}
