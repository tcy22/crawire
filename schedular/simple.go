package schedular

import "crawier/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

//构造chan
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request){
	s.workerChan = c
}

//往scheduler的chan里面发任务,代替了队列，抢占chan
func (s *SimpleScheduler) Submit(r engine.Request){
	go func(){
		s.workerChan <- r
	}()
}
