package main

import (
	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/persist"
	"imooc.com/tutorial/crawler/scheduler"
	"imooc.com/tutorial/crawler/zhenai/parser"
)

func main() {
	// 原先的elsatic search的命名是放在save里的，把这个拿出来放在main传入
	// 有两个好处：配置全部写在main里，测试的时候可以另外开一个index，把生产和测试分开
	itemsaver, err := persist.ItemSaver("profile")
	if err != nil {
		panic(err) //如果elastic search 没有起来爬虫就无法运行了
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      300,
		ItemChan:         itemsaver,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}
