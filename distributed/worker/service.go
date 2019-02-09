package worker

import (
	"go-spider/engine"
)

type WorkerService struct{}

func (WorkerService) Process(req Request, result *[]ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return nil
	}

	*result = SerializeResult(engineResult)

	return nil

}
