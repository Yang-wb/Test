package main

import (
	"fmt"
	"leango/crawler_distributed/config"
	"leango/crawler_distributed/persist"
	"leango/crawler_distributed/rpcsupport"
	"log"

	"github.com/olivere/elastic/v7"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.SaverRpc(host, &persist.ItemSaveService{Client: client, Index: index})
}
