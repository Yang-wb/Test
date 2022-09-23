package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/mvcc/mvccpb"

	"github.com/astaxie/beego/logs"
	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"
)

var (
	redisPoll  *redis.Pool
	etcdClient *etcd_client.Client
)

func initRedis() (err error) {
	redisPoll = &redis.Pool{
		MaxIdle:     secKillConfig.redisConfig.redisMaxIdle,
		MaxActive:   secKillConfig.redisConfig.redisMaxActive,
		IdleTimeout: time.Duration(secKillConfig.redisConfig.redisIdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConfig.redisConfig.redisAddr)
		},
	}

	conn := redisPoll.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err:%v", err)
		return
	}
	return
}

func initEtcd() (err error) {
	etcdClient, err = etcd_client.New(etcd_client.Config{
		Endpoints:   []string{secKillConfig.etcdConfig.etcdAddr},
		DialTimeout: time.Duration(secKillConfig.etcdConfig.etcdTimeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	return
}

func initLogs() (err error) {
	config := make(map[string]interface{})
	config["filename"] = secKillConfig.logConfig.logPath
	config["level"] = convertLogLevel(secKillConfig.logConfig.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

func convertLogLevel(level string) int {

	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}

	return logs.LevelDebug
}

func loadSecConf() (err error) {
	key := secKillConfig.etcdConfig.etcdSecProductKey
	resp, err := etcdClient.Get(context.Background(), key)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err:%v", key, err)
	}
	var secProductInfo []SecProductInfoConf
	for k, v := range resp.Kvs {
		logs.Debug("key[%v],value[%v]", k, v)
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			logs.Error("Unmarshal sec product info failed, err:%v", err)
			return
		}
		logs.Debug("sec info conf is [%v]", secProductInfo)
	}

	secKillConfig.secProductInfo = secProductInfo
	return
}

func initSecProductWatcher() {
	go watchKey(secKillConfig.etcdConfig.etcdSecProductKey)
}

func watchKey(key string) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"192.168.56.110:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect failed, err:", err)
	}

	fmt.Println("connect succ")
	defer cli.Close()

	logs.Debug("etcd:" + key)

	for {
		rch := cli.Watch(context.Background(), key)
		var secProductInfo []SecProductInfoConf

		getConfSucc := true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", key)
					continue
				} else if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err := json.Unmarshal(ev.Kv.Value, &secProductInfo)
					if err != nil {
						logs.Error("key [%s], JNmarshal[%s], err:%v", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd,%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", secProductInfo)
				updateSecProductInfo(secProductInfo)
			}
		}
	}
}

func updateSecProductInfo(secProductInfo []SecProductInfoConf) {

}

func initSec() (err error) {

	err = initLogs()
	if err != nil {
		logs.Error("init logger failed, err %v", err)
		return err
	}

	err = initRedis()
	if err != nil {
		logs.Error("init redis failed, err %v", err)
		return err
	}

	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed , err %v", err)
		return err
	}

	err = loadSecConf()
	if err != nil {
		logs.Error("init loadSecConf failed , err %v", err)
	}

	initSecProductWatcher()

	return nil
}
