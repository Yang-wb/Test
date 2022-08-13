package client

import (
	"leango/crawler/engine"
	"leango/crawler_distributed/config"
	"leango/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client, err := rpcsupport.NewClient(fmt.Sprint(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}
	return func(request engine.Request) (result engine.ParseResult, err error) {

		sReq := worker.SerializeRequest(request)

		var sResult worker.ParseResult
		c := <-clientChan
		err = c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
