package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got itme #%d: %v", itemCount, item)
			itemCount++

			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver: error:%v item:%v", err, item)
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

	resp, err := client.Index().
		Index("dating_profile").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
