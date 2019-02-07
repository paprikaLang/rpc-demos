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
func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Saved item #%d: %v ", itemCount, item)
			itemCount++
			err := save(item)
			if err != nil {
				log.Printf("save %v got error: %v", item, err)
			}
		}
	}()
	return out
}
func save(item engine.Item) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	if item.Type == "" {
		return errors.New("must supply elastic Type")
	}
	indexService := client.Index().Index("laravel_collections").Type(item.Type).BodyJson(item)
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
