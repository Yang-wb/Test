package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"lengo2/logagent/tailf"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel  string
	logPath   string
	chanSize  int
	kafkaAddr string

	collectConf []tailf.CollectConf

	etcdAddr string
	etcdKey  string
}

func loadConf(confType, filename string) error {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return err
	}

	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}

	appConfig.chanSize, err = conf.Int("collect::chan_size")
	if len(appConfig.logPath) == 0 {
		appConfig.chanSize = 100
	}

	appConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.kafkaAddr) == 0 {
		return fmt.Errorf("invalid kafka addr")
	}

	appConfig.etcdAddr = conf.String("etcd::addr")
	if len(appConfig.etcdAddr) == 0 {
		return fmt.Errorf("invalid cted addr")
	}

	appConfig.etcdKey = conf.String("etcd::configKey")
	if len(appConfig.etcdKey) == 0 {
		return fmt.Errorf("invalid cted key")
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect cof failed, err:%v\n", err)
	}
	return nil
}

func loadCollectConf(conf config.Configer) error {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		return errors.New("invalid collect::log_path")
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		return errors.New("invalid collect::topic")
	}

	appConfig.collectConf = append(appConfig.collectConf, cc)
	return nil
}
