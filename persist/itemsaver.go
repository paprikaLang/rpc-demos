package persist

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Saved item #%d: %v ", itemCount, item)
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Printf("save %v got error: %v", item, err)
			}
		}
	}()
	return out
}
func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	res, err := client.Index().Index("laravel_collection").Type("posts").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	fmt.Printf("res: %+v", res) //显示结构体字段名
	return res.Id, nil
}
