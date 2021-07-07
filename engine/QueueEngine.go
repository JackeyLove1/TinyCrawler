package engine

type NewScheduler interface {
	Submit(Request)
	WorkerChan() chan Request
}
