package scheduler

import (
	"go-spider/engine"
)

type GoroutineScheduler struct {
	workerChan chan engine.Request
}

func (s *GoroutineScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *GoroutineScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *GoroutineScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *GoroutineScheduler) WorkerReady(w chan engine.Request) {

}
