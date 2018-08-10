package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"

	"imooc.com/tutorial/crawler_distributed/rpcsupport"

	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/scheduler"
	"imooc.com/tutorial/crawler/zhenai/parser"
	itemSaver "imooc.com/tutorial/crawler_distributed/persist/client"
	worker "imooc.com/tutorial/crawler_distributed/worker/client"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemsaver, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err) //如果elastic search 没有起来爬虫就无法运行了
	}
	hosts := strings.Split(*workerHosts, ",")
	pool := createClientPool(hosts)
	processor := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      300,
		ItemChan:         itemsaver,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}

// TODO: rpc的server程序挂了重启会无法进行工作而且也没法看到报错
// 应该是阻塞在channel上了
func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		c, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, c)
			log.Printf("connceting to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				clientChan <- client
			}
		}
	}()

	return clientChan
}
