package engine

type Engine struct{
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

type Scheduler interface {
	ReadyNotifier
	WorkChan() chan Request
	Run()
	Submit(Request)
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}