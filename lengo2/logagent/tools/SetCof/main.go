package main

import (
	"context"
	"encoding/json"
	"fmt"
	"lengo2/logagent/tailf"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const (
	EtcdKey = "/oldboy/backend/logagent/config/192.168.1.109"
)

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.56.110:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	var logConfArr []tailf.CollectConf
	logConfArr = append(
		logConfArr,
		tailf.CollectConf{
			LogPath: "F:\\搬迁软件\\学习的\\test\\controller.log",
			Topic:   "nginx_log",
		},
	)

	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed, ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//cli.Delete(ctx, EtcdKey)
	//return
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	SetLogConfToEtcd()
}
