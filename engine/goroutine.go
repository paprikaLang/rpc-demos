package engine

// import "log"

// type GoroutineEngine struct {
// 	Scheduler Scheduler
// 	Workers   int
// }

// // GoorutineScheduler 实现了协议中的方法
// type Scheduler interface {
// 	Submit(Request)
// 	GoroutineSchedulerWorkerChan(chan Request)
// }

// func (c *GoroutineEngine) Run(seeds ...Request) {

// 	in := make(chan Request)
// 	out := make(chan ParseResult)
// 	c.Scheduler.GoroutineSchedulerWorkerChan(in)
// 	for i := 0; i < c.Workers; i++ {
// 		createWorker(in, out)
// 	}

// 	for _, r := range seeds {
// 		// Scheduler将request全部通过channel传递给workers, 为避免相互等待, 会为每一个request传递的过程放到goroutine中进行
// 		c.Scheduler.Submit(r)
// 	}
// 	itemsCount := 0
// 	for {
// 		result := <-out
// 		for _, item := range result.Items {
// 			log.Printf("Got item #%d: %v ", itemsCount, item)
// 			itemsCount++
// 		}
// 		for _, request := range result.Requests {
// 			c.Scheduler.Submit(request)
// 		}
// 	}
// }

// func createWorker(in chan Request, out chan ParseResult) {
// 	go func() {
// 		for {
// 			request := <-in
// 			result, err := worker(request)
// 			if err != nil {
// 				continue
// 			}
// 			out <- result
// 		}
// 	}()
// }
