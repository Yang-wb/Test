package client

import (
	"leango/crawler/engine"
	"leango/crawler_distributed/config"
	"leango/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
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

			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error:%v item:%v", err, item)
			}
		}
	}()
	return out, nil
}
