package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	etcd_client "github.com/coreos/etcd/clientv3"
	"lengo2/logagent/tailf"
	"strings"
	"time"
)

type EtcdClient struct {
	client *etcd_client.Client
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
