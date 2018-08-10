package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/olivere/elastic.v5"
	"imooc.com/tutorial/crawler_distributed/config"
	"imooc.com/tutorial/crawler_distributed/persist"
	"imooc.com/tutorial/crawler_distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	// err := serveRpc(":1234", "dating_profile")
	// if err != nil {
	// 	panic(err)
	// }
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))

}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	// &persist.ItemSaverService才
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})

}
