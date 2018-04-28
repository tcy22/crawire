package parser

import (
"regexp"
"crawier/engine"
	"fmt"
)

const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//城市列表解析器
func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents,-1)     //[][][]byte

	result := engine.ParseResult{}
	for _,m := range matches {
		name := string(m[2])
		result.Items = append(result.Items,"User "+name)
		fmt.Println("cityParse,每个人的信息Url:",string(m[1]))
		result.Requests = append(result.Requests,engine.Request{
			Url:        string(m[1]),
			//定义了一个闭包，在两个参数不同的函数中起了一个比较好的桥梁
			ParserFunc: func(contents []byte) engine.ParseResult{
				return parserProfile(contents,name)
			},
		})
	}
	return result
}

