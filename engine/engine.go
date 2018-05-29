package engine

import (
	"crawier/fetcher"
	"log"
)

func Run(seeds ...Request){
	var requests []Request
	for _,r := range seeds {
		requests = append(requests,r)
	}

	for len(requests)>0 {
		r := requests[0]
		requests =requests[1:]

		log.Printf("Fetching %s",r.Url)
		body,err := fetcher.Fetcher(r.Url) //从网络上获取数据，然后由不同的解析器解析数据
		if err != nil {
			log.Printf("Fetcher:error fetching url %s,%v",r.Url,err)
			continue
		}

		ParseResult := r.ParserFunc(body)
		requests = append(requests,ParseResult.Requests...)//把slice里面的内容展开一个一个加进里面

		for _,item := range ParseResult.Items {
			log.Printf("Got item %s",item)
		}
	}
}