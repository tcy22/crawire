package engine

import (
	"fmt"
	"log"
	"crawier/fetcher"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)        //往scheduler的chan里面发任务
	ConfigureMasterWorkerChan(chan Request) //构造chan_in
}

func (e *ConcurrentEngine) Run(seeds ... Request){

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)//构造输入chan

	for i:=0; i < e.WorkerCount; i++ {
		createWorker(in,out)         //开10个可以爬取信息和解析信息的goroutine
	}
    //把对engine的请求全部交给scheduler
	for _,r := range seeds {//往scheduler的chan里面发任务,代替了队列，抢占chan
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
			Request := <- in       //输入的request
			result,err := workerc(Request)
			if err != nil {
				continue
			}
			out <- result         //爬取URL，且解析出有用内容的返回信息
		}
	}()
}

func workerc(r Request)(ParseResult,error){
	log.Printf("Fetching %s",r.Url)
	body,err := fetcher.Fetcher(r.Url) //从网络上获取数据，然后由不同的解析器解析数据
	if err != nil {
		log.Printf("Fetcher:error fetching url %s,%v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil
}