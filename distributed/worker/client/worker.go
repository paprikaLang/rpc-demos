package client

import (
	"go-spider/distributed/rpcconfig"
	"go-spider/distributed/worker"
	"go-spider/engine"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcconfig.RPCClient(":2334")
	if err != nil {
		return nil, err
	}
	return func(req engine.Request) ([]engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sRes []worker.ParseResult
		err := client.Call("WorkerService.Process", sReq, &sRes)
		if err != nil {
			return []engine.ParseResult{}, err
		}
		// fmt.Println(sRes)
		return worker.DeserializeResult(sRes), nil
	}, nil
}
