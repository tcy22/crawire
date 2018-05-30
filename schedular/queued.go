package schedular

import "crawier/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
	//每个worker对外的接口都是一个channel request,每个work都是一个channel
	//把每个任务都送给channel request
	// 将开的N个goroutine都放在一个channel里面
}

//构造chan
func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request){
	panic("implement me")
}

//往scheduler的requestchan里面发任务
func (s *QueuedScheduler) Submit(r engine.Request){
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request){
	s.workerChan <- w
}

func (s *QueuedScheduler) Run(){
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func(){
		var requestQ []engine.Request    //请求队列
		var workerQ  []chan engine.Request //work队列，此队列中每个元素是可以放request的channel
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0 && len(workerQ)>0{
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ,r)  //收到request,排队
			case w := <-s.workerChan:
				workerQ = append(workerQ,w)    //收到worker，排队
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()
}