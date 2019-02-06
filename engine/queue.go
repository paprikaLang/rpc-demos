package engine

// import "log"

// type QueueEngine struct {
// 	Scheduler Scheduler
// 	Workers   int
// }

// type Scheduler interface {
// 	Submit(Request)
// 	QueueSchedulerWorkerChan(chan Request)
// 	WorkerReady(chan Request)
// 	Run()
// }

// func (q *QueueEngine) Run(seeds ...Request) {

// 	out := make(chan ParseResult)
// 	q.Scheduler.Run()
// 	for i := 0; i < q.Workers; i++ {
// 		createWorker(out, q.Scheduler) //每个worker都会创建一个request channel
// 	}

// 	for _, r := range seeds {
// 		q.Scheduler.Submit(r)
// 	}
// 	itemsCount := 0
// 	for {
// 		result := <-out
// 		for _, item := range result.Items {
// 			log.Printf("Got item #%d: %v ", itemsCount, item)
// 			itemsCount++
// 		}
// 		for _, request := range result.Requests {
// 			q.Scheduler.Submit(request)
// 		}
// 	}
// }

// func createWorker(out chan ParseResult, s Scheduler) {
// 	in := make(chan Request)
// 	go func() {
// 		for {
// 			s.WorkerReady(in)
// 			request := <-in
// 			result, err := worker(request)
// 			if err != nil {
// 				continue
// 			}
// 			out <- result
// 		}
// 	}()
// }
