# crawire
golang----爬虫

Seed:  URL+Parser,从URL中爬出来的数据由相应的Parser解析，解析出新的URL继续加入到任务队列中，继续爬。

Engine: 将请求都加入到任务队列中，一个一个执行

Fetcher: 从网络上获取数据的模块

Parser:  解析器

城市列表解析器：传入html文档，出来结构化的数据，城市名和URL

城市解析器：传入URL，解析出该城市所有人的名字和URL

用户解析器：点一个人的名字，将该人的一些重要的信息解析出来

输入：utf-8编码的文本

输出：Request{URL,对应Parser}列表，Item列表（获取到的有价值的内容）




