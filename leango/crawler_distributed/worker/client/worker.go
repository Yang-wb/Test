package client

import (
	"fmt"
	"leango/crawler/engine"
	"leango/crawler_distributed/config"
	"leango/crawler_distributed/rpcsupport"
	"leango/crawler_distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprint(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	return func(request engine.Request) (result engine.ParseResult, e error) {

		sReq := worker.SerializeRequest(request)

		var sResult worker.ParseResult
		err = client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
