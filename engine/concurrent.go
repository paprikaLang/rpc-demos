package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	Workers   int
	ItemSaver chan interface{}
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
	WorkerChannelInQueue
}

type WorkerChannelInQueue interface {
	WorkerReady(chan Request)
}

func (q *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	q.Scheduler.Run()
	for i := 0; i < q.Workers; i++ {
		createWorker(q.Scheduler.WorkerChan(), out, q.Scheduler) //每个worker都会创建一个request channelq
	}

	for _, r := range seeds {
		q.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(item interface{}) {
				q.ItemSaver <- item // loop variable item captured by func literal
			}(item)
		}
		for _, request := range result.Requests {
			q.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, w WorkerChannelInQueue) {

	go func() {
		for {
			w.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			for _, item := range result {
				out <- item
			}

		}
	}()
}
