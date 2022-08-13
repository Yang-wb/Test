package main

import (
	"flag"
	"leango/crawler/engine"
	"leango/crawler/scheduler"
	"leango/crawler/zhenai/parser"
	"leango/crawler_distributed/config"
	itemsaver "leango/crawler_distributed/persist/client"
	"leango/crawler_distributed/rpcsupport"
	worker "leango/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_host", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenhun", Parser: engine.NewFuncParser(parser.ParseCityList, config.ParserCityList)})
}

func createClientPool(host []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range host {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
