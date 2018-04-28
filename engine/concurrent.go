package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)        //往scheduler的chan里面发任务
	ConfigureMasterWorkerChan(chan Request) //构造chan_in
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ... Request){
	out := make(chan ParseResult)
	e.Scheduler.Run()   //等待新的任务的到来

	for i:=0; i < e.WorkerCount; i++ {
		createWorker(out,e.Scheduler)
	}
    //把对engine的请求全部交给scheduler
	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0    //计数器
	for{
		result := <- out
		for _,item := range result.Items{
			fmt.Printf("Got item #%d: %v",itemCount,item)
			itemCount++
		}
		for _,request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParseResult,s Scheduler){
	in := make(chan Request)
	go func() {
		for{
			// tell scheduler I am ready
			s.WorkerReady(in)
			Request := <- in
			result,err := worker(Request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}