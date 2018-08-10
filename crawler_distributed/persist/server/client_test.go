package main

import (
	"testing"
	"time"

	"imooc.com/tutorial/crawler_distributed/rpcsupport"

	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/model"
)

func TestItemSaver(t *testing.T) {

	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(time.Second * 2)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108848769",
		Id:   "108848769",
		Type: "zhenai",
		Payload: model.Profile{
			Name:     "如果芸许",
			Gender:   "女",
			Age:      35,
			Height:   155,
			Income:   "3000元以下",
			Marriage: "离异"},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s, err: %s", result, err)
	}
}
