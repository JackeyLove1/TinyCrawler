package schedule

import (
	"awesomeCrawl/engine"
)

type SimpleSchdule struct {
	workChan chan engine.Request
}

func (s SimpleSchdule) Submit(request engine.Request) {
	go func() {
		s.workChan <- request
	}()
}

func (s SimpleSchdule) ConfigMasterWorkerChan(in chan engine.Request) {
	s.workChan = in
	// fmt.Println(reflect.DeepEqual(s.workChan, in))
}
