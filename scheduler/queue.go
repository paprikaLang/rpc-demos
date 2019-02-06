package scheduler

import "go-spider/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueueScheduler) Submit(r engine.Request) {
	go func() {
		q.requestChan <- r // 传递过来的 request 接收到就先放在队列里
	}()
}

func (q *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w  // createWorker 产生的 用于 request -> worker 的 channel
}

func (q *QueueScheduler) Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeR engine.Request
			var activeW chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeR = requestQ[0]
				activeW = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			case activeW <- activeR:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
