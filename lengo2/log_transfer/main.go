package main

import "github.com/astaxie/beego/logs"

func main() {
	err := initConfig("init", "./log_transfer/conf/log_transfer.conf")
	if err != nil {
		panic(err)
		return
	}

	err = initLogger(logConfig.logPath, logConfig.logLevel)
	if err != nil {
		panic(err)
		return
	}

	err = initKafka(logConfig.kafkaAddr, logConfig.topic)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
	}
	logs.Debug("init kafka succ")

	err = initEs(logConfig.esAddr)
	if err != nil {
		logs.Error("init es failed, err:%v", err)
	}

	logs.Debug("init es succ")

	err = run()
	if err != nil {
		logs.Error("run run, err:%v", err)
	}

	logs.Warn("warning,log_transfer is exited")
}
