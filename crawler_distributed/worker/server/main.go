package main

import (
	"flag"
	"fmt"
	"log"

	"imooc.com/tutorial/crawler_distributed/rpcsupport"
	"imooc.com/tutorial/crawler_distributed/worker"
)

// golang的命令行参数
var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlerService{}))

}
