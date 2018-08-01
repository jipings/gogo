package main

import (
	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
