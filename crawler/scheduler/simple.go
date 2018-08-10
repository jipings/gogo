package scheduler

import "imooc.com/tutorial/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()

}
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan

}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {

}

func (s *SimpleScheduler) Run() {

	s.workerChan = make(chan engine.Request)

}
