package schedular

import (
	"crawier/engine"
)

type SimpleScheduler struct {//负责输入channel
	workerChan chan engine.Request
}

//构造chan
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request){
	s.workerChan = c
}

//往scheduler的chan里面发任务,代替了队列，抢占chan
func (s *SimpleScheduler) Submit(r engine.Request){
	//这里如果不加go func，会卡，因为只开了10个goroutine去接channel里的值，如果超过10个去Submit，就会卡死。
	go func(){
		s.workerChan <- r
	}()
}
