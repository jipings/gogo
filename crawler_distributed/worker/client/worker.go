package client

import (
	"log"
	"net/rpc"

	"imooc.com/tutorial/crawler_distributed/worker"

	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler_distributed/config"
)

type ClientInfo struct {
	Client *(*rpc.Client)
	Host   string
}

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlerServiceRpc, sReq, &sResult)

		if err != nil {
			log.Println("Processor call error", err)
			return engine.ParseResult{}, err

		}
		return worker.DeserializeResult(sResult), nil

	}

}
