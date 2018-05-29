package engine

type Request struct { //请求的URL和对应解析器（把从URL中爬出来的数据进行解析）
	Url        string
	ParserFunc  func([]byte) ParseResult
}

type ParseResult struct {//数据解析后的返回信息,Items（有用的信息）
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult{
	return ParseResult{}
}
