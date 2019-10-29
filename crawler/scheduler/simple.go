package scheduler

import "gostudy/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
	panic("implement me")
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}


func (s SimpleScheduler) Submit(request engine.Request) {
	// send request down to
	go func() {
		s.workerChan <- request
	}()
}
