package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"lengo2/logagent/kafka"
	"lengo2/logagent/tailf"
)

func main() {
	filename := "./logagent/conf/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}

	//从etcd中获取配置
	collectConf, err := initEtcd(appConfig.etcdAddr, appConfig.etcdKey)
	if err != nil {
		logs.Error("init etcd failed,err:%v", err)
		return
	}
	logs.Debug("etcd succ")

	logs.Debug("initialize succ")
	logs.Debug("load conf succ, conf:%v", appConfig)

	err = tailf.InitTail(collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}

	logs.Info("program exited")
}
