package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Schedule    SimpleSchedule
	WorkerCount int // 并发数
	// ItemChan chan interface{}
}

type SimpleSchedule interface {
	Submit(Request)
	Config(in chan Request)
}

type SimpleScheduler struct {
	workChan chan Request
}

func (s *SimpleScheduler) Submit(request Request) {
	// panic("implement me")
	go func() {
		s.workChan <- request
	}()
}

func (s *SimpleScheduler) Config(in chan Request) {
	// panic("implement me")
	s.workChan = in
	// fmt.Println("len: ", len(s.workChan))
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)

	e.Schedule.Config(in)

	for i := 0; i < e.WorkerCount; i++ {
		createSimpleWorker(in, out)
	}

	for _, r := range seeds {
		e.Schedule.Submit(r)
	}

	for{
		result := <- out

		items := result.Items
		go func(item2 []interface{}) {
			for _, item := range item2 {
				log.Printf("Got Item %v\n", item)
			}
		}(items)

		for _, r := range result.Requests{
			e.Schedule.Submit(r)
		}

	}
}

func createSimpleWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			r := <-in

			log.Printf("Fetching ... %s\n", r.Url)
			result, err := worker(r)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
