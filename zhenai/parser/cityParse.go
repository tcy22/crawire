package parser

import (
"regexp"
"crawier/engine"
)

var cityRe  = regexp.MustCompile(
	`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRe  = regexp.MustCompile(
	`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

//城市列表解析器
func ParseCity(contents []byte) engine.ParseResult{
	matches := cityRe.FindAllSubmatch(contents,-1)     //[][][]byte

	result := engine.ParseResult{}
	for _,m := range matches {
		url := string(m[1])
		name := string(m[2])
		//result.Items = append(result.Items,"User "+name)
		result.Requests = append(result.Requests,engine.Request{
			Url:        url,
			//定义了一个闭包，在两个参数不同的函数中起了一个比较好的桥梁
			ParserFunc: func(contents []byte) engine.ParseResult{
				return parserProfile(contents,url,name)
			},
		})
	}
	//爬下一页
	matches = cityUrlRe.FindAllSubmatch(contents,-1)
	for _,m := range matches {
		result.Requests = append(result.Requests,engine.Request{
			Url:        string(m[1]),
			//定义了一个闭包，在两个参数不同的函数中起了一个比较好的桥梁
			ParserFunc: ParseCity,
		})
	}
	return result
}

