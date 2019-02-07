package persist

import (
	"context"
	"errors"
	"fmt"
	"go-spider/engine"
	"log"

	"github.com/olivere/elastic"
)

// ItemSaver save item to elasticsearch
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Saved item #%d: %v ", itemCount, item)
			itemCount++
			err := save(client, item, index)
			if err != nil {
				log.Printf("save %v got error: %v", item, err)
			}
		}
	}()
	return out, nil
}
func save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("must supply elastic Type")
	}
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	res, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("res: %+v", res) //+显示结构体字段名
	return nil
}
