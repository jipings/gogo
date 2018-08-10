package persist

import (
	"context"
	"errors"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
	"imooc.com/tutorial/crawler/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)

	// 把new client从save里面分离使得client只有一个减少开销
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d, %v", itemCount, item)
			itemCount++
			err := Save(client, item, index)
			if err != nil {
				log.Printf("Item saver: error saveing items %v:%v", item, err)
			}
		}

	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) (err error) {

	log.Printf("Got Item %v", item)
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(index).    //创建是“数据库“，由配置文件给出
		Type(item.Type). //创建“表”,由Parser给出
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id) // id由Parser给出
	}

	_, err = indexService.Do(context.Background())
	// log.Printf("%+v", resp)
	if err != nil {
		return err
	}

	return nil

}
