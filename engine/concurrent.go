package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)        //往scheduler的chan里面发任务
	ConfigureMasterWorkerChan(chan Request) //构造chan_in
}

func (e *ConcurrentEngine) Run(seeds ... Request){
/*	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}*/

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i:=0; i < e.WorkerCount; i++ {
		createWorker(in,out)
	}
    //把对engine的请求全部交给scheduler
	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	//计数器
	itemCount := 0
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

func createWorker(in chan Request,out chan ParseResult){
	go func() {
		for{
			Request := <- in
			result,err := worker(Request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}