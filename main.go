package main

import (
	"awesomeCrawl/engine"
	"awesomeCrawl/parse"
)

const seed = "https://book.douban.com"

// single
//func main() {
//	e := engine.SimpleEngine{}
//	e.Run(engine.Request{
//		Url: seed,
//		ParseFunc: parse.ParseTag,
//	})
//}

// concurrent
func main() {
	e := engine.ConcurrentEngine{
		Schedule: &engine.SimpleScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{
		Url: seed,
		ParseFunc: parse.ParseTag,
	})
}
