package main

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

type SecKillConfig struct {
	etcdConfig     EtcdConfig
	redisConfig    RedisConfig
	logConfig      LogConfig
	secProductInfo []SecProductInfoConf
}

type RedisConfig struct {
	redisAddr        string
	redisMaxIdle     int
	redisMaxActive   int
	redisIdleTimeout int
}

type EtcdConfig struct {
	etcdAddr          string
	etcdTimeout       int
	etcdSecKeyPrefix  string
	etcdSecProductKey string
}

type LogConfig struct {
	logPath  string
	logLevel string
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int
	EndTime   int
	Status    int
	Total     int
	Left      int
}

var (
	secKillConfig = &SecKillConfig{}
)

func initConfig() error {
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")

	secKillConfig.redisConfig.redisAddr = redisAddr
	secKillConfig.etcdConfig.etcdAddr = etcdAddr

	if len(redisAddr) == 0 || len(etcdAddr) == 0 {
		err := fmt.Errorf("init config faild, redis[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return err
	}

	redisMaxIdle, err := beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		err := fmt.Errorf("init config faild, redis_max_idle[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return err
	}
	secKillConfig.redisConfig.redisMaxIdle = redisMaxIdle

	redisMaxActive, err := beego.AppConfig.Int("redis_max_active")
	if err != nil {
		err := fmt.Errorf("init config faild, redisMaxActive[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return err
	}
	secKillConfig.redisConfig.redisMaxActive = redisMaxActive

	redisIdleTimeout, err := beego.AppConfig.Int("redis_idle_timeout")
	if err != nil {
		err := fmt.Errorf("init config faild, redisIdleTimeout[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return err
	}
	secKillConfig.redisConfig.redisIdleTimeout = redisIdleTimeout

	etcdTimeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		err := fmt.Errorf("init config faild, etcd_timeout config is null")
		return err
	}
	secKillConfig.etcdConfig.etcdTimeout = etcdTimeout

	etcdSecKeyPrefix := beego.AppConfig.String("etcd_sec_key")
	if len(etcdSecKeyPrefix) == 0 {
		err := fmt.Errorf("init config faild, etcdSecKeyPrefix config is null")
		return err
	}
	secKillConfig.etcdConfig.etcdSecKeyPrefix = etcdSecKeyPrefix

	etcdSecProductKey := beego.AppConfig.String("etcd_product_key")
	if len(etcdSecProductKey) == 0 {
		err := fmt.Errorf("init config faild, etcdProductKey config is null")
		return err
	}
	if strings.HasSuffix(secKillConfig.etcdConfig.etcdSecKeyPrefix, "/") {
		secKillConfig.etcdConfig.etcdSecProductKey = fmt.Sprintf("%s%s", secKillConfig.etcdConfig.etcdSecKeyPrefix, etcdSecProductKey)
	} else {
		secKillConfig.etcdConfig.etcdSecProductKey = fmt.Sprintf("%s/%s", secKillConfig.etcdConfig.etcdSecKeyPrefix, etcdSecProductKey)

	}

	secKillConfig.logConfig.logPath = beego.AppConfig.String("log_path")
	secKillConfig.logConfig.logLevel = beego.AppConfig.String("log_level")

	return nil
}
