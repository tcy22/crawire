package main

import (
	"crawier/engine"
	"crawier/zhenai/parser"
	"crawier/schedular"
	"crawier/save"
)

func main(){
/*	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc:  parser.ParseCityList,
	})*/
	itemChan,err := save.ItemSave()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: 		&schedular.QueuedScheduler{},
		WorkerCount:	10,
		ItemChan:  		itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc:  parser.ParseCityList,
	})
}