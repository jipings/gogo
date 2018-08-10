package engine

import "log"

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

// 重构：统一不同scheduler，减少interface的成员
// queued scheduler 和simple scheduler区别在于worker是否拥有自己的channel
// 而且关键点在于只有scheduler知道worker
// type Scheduler interface {
// 	Submit(Request)
// 	ConfigureMasterWorkerChan(chan Request)
// 	WorkerReady(chan Request)
// 	Run()
// }

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
	ReadyNotifier
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		log.Printf("Create worker #%d...", i)
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}
	log.Print("Init ready...")

	for {
		result := <-out

		for _, item := range result.Items {

			go func(item Item) { e.ItemChan <- item }(item)

		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}

	}

}

// worker的创建奥义就是去问scheduler拿request channel，scheduler怎么给就不用管了
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in

			// log.Printf("Fetching URL %s", request.Url)
			// res, err := Worker(request)
			res, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}
