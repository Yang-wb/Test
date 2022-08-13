package main

import (
	"fmt"
	"leango/crawler/engine"
	"leango/crawler/scheduler"
	"leango/crawler/zhenai/parser"
	"leango/crawler_distributed/config"
	itemsaver "leango/crawler_distributed/persist/client"
	worker "leango/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenhun", Parser: engine.NewFuncParser(parser.ParseCityList, config.ParserCityList)})
}
