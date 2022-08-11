package main

import (
	"fmt"
	"leango/crawler/engine"
	"leango/crawler/scheduler"
	"leango/crawler/zhenai/parser"
	"leango/crawler_distributed/config"
	"leango/crawler_distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenhun", ParserFunc: parser.ParseCityList})
}
