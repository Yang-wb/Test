package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"lengo2/logagent/tailf"
	"strings"
	"time"
)

type EtcdClient struct {
	client *etcd_client.Client
	keys   []string
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr, key string) (conf []tailf.CollectConf, err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"192.168.56.110:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect failed, err:", err)
		return nil, err
	}

	fmt.Println("connect succ")
	defer cli.Close()

	etcdClient = &EtcdClient{client: cli}

	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	var collectConf []tailf.CollectConf
	for _, ip := range localIpArray {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed, err:%v", err)
			continue
		}
		cancel()

		logs.Debug("resp from etcd:%v", resp.Kvs)
		for _, v := range resp.Kvs {
			if string(v.Value) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed, err:%v", err)
					continue
				}

				logs.Debug("log config is %v", collectConf)
			}
		}
	}

	return collectConf, nil
}

func initEtcdWatcher() {
	for _, key := range etcdClient.keys {
		watchKey(key)
	}
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
		var conf []tailf.CollectConf

		getConfSucc := true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", key)
					continue
				} else if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err := json.Unmarshal(ev.Kv.Value, &conf)
					if err != nil {
						logs.Error("key [%s], JNmarshal[%s], err:%v", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd,%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", conf)
				tailf.UpdateConfig(conf)
			}
		}
	}
}
