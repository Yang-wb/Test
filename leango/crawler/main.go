package main

import (
	"leango/crawler/engine"
	"leango/crawler/scheduler"
	"leango/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenhun", ParserFunc: parser.ParseCityList})
}
