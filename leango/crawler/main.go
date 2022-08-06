package main

import (
	"leango/crawler/engine"
	"leango/crawler/persist"
	"leango/crawler/scheduler"
	"leango/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenhun", ParserFunc: parser.ParseCityList})
}
