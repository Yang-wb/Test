package worker

import "leango/crawler/engine"

type CrawlService struct {
}

func (c CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineRequest, err := engine.Worker(engineReq)
	if err != nil {
		return nil
	}
	*result = SerializeResult(engineRequest)
	return nil
}
