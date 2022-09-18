package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type LofConfig struct {
	kafkaAddr string
	topic     string
	esAddr    string
	logPath   string
	logLevel  string
}

var (
	logConfig *LofConfig
)

func initConfig(confType string, filename string) error {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return err
	}

	logConfig = &LofConfig{}
	logConfig.logLevel = conf.String("logs::log_level")
	if len(logConfig.logLevel) == 0 {
		logConfig.logLevel = "debug"
	}

	logConfig.logPath = conf.String("logs::log_path")
	if len(logConfig.logPath) == 0 {
		logConfig.logPath = "./logs"
	}

	logConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(logConfig.kafkaAddr) == 0 {
		return fmt.Errorf("invalid kafka addr")
	}

	logConfig.topic = conf.String("kafka::topic")
	if len(logConfig.topic) == 0 {
		return fmt.Errorf("invalid kafka topic")
	}

	logConfig.esAddr = conf.String("es::addr")
	if len(logConfig.esAddr) == 0 {
		return fmt.Errorf("invalid es addr")
	}

	return nil
}
