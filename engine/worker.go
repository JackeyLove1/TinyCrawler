package engine

import (
	"awesomeCrawl/fetcher"
	"log"
)

// 工作函数
func worker(r Request) (ParseResult, error){
	doc, err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("fetch url error: %s\n", err)
		return ParseResult{}, err
	}

	return r.ParseFunc(doc)
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier){
	go func() {
		for{
			ready.WorkerReady(in) // 告诉调度器任务空闲
			req := <- in
			result, err := worker(req)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}