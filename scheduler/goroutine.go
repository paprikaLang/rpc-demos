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

func (s *GoroutineScheduler) GoroutineSchedulerWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
