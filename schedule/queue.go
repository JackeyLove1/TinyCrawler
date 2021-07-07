package schedule

import (
	"awesomeCrawl/engine"
)

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueueScheduler) WorkerReady(requests chan engine.Request) {
	q.workerChan <- requests
}

func (q *QueueScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueueScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)

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
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
