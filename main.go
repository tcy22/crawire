package main

import (
	"crawier/engine"
	"crawier/zhenai/parser"
	"crawier/schedular"
)

func main(){
/*	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc:  parser.ParseCityList,
	})*/
	e := engine.ConcurrentEngine{
		Scheduler: &schedular.SimpleScheduler{},
		WorkerCount:10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc:  parser.ParseCityList,
	})
}