package persist

import (
	"context"
	"errors"
	"leango/crawler/engine"
	"log"

	"github.com/olivere/elastic/v7"
)

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
			log.Printf("Item Saver: got itme #%d: %v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error:%v item:%v", err, item)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("must supple type")
	}

	indexService := client.Index().
		Index(index).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
