package main

import (
	"leango/crawler/engine"
	"leango/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{Url: "http://www.zhenai.com/zhenhun", ParserFunc: parser.ParseCityList})
}
