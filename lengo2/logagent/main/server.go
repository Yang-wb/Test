package main

import (
	"github.com/astaxie/beego/logs"
	"lengo2/logagent/kafka"
	"lengo2/logagent/tailf"
	"time"
)

func serverRun() error {
	for {
		msg := tailf.GetOneLine()
		err := sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}

	return nil
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	//logs.Debug("read msg:%s,topic:%s",msg.Msg,msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	if err != nil {
		return err
	}
	return nil
}
