package client

import (
	"log"

	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler_distributed/config"
	"imooc.com/tutorial/crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {

	out := make(chan engine.Item)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	result := ""
	go func() {
		count := 0
		for {
			item := <-out
			count++
			log.Printf("got item #%d:%v", count, item)
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil || result != "ok" {
				log.Printf("Item saver: error saveing items %v:%v", item, err)
			}
		}
	}()

	return out, nil

}
