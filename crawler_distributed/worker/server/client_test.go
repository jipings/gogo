package main

import (
	"testing"
	"time"

	"imooc.com/tutorial/crawler_distributed/config"
	"imooc.com/tutorial/crawler_distributed/rpcsupport"
	"imooc.com/tutorial/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlerService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/108848769",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "如果芸许",
		},
	}

	var res worker.ParseResult
	err = client.Call(config.CrawlerServiceRpc, req, &res)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", res)
	}

}
