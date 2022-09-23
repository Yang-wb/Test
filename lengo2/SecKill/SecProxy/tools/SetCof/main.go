package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const (
	EtcdKey = "/backend/seckill/product"
)

type SecInfoConf struct {
	ProductId int
	StartTime int
	EndTime   int
	Status    int
	Total     int
	Left      int
}

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

	var secInfoConf []SecInfoConf
	secInfoConf = append(
		secInfoConf,
		SecInfoConf{
			ProductId: 1028,
			StartTime: 1663942112,
			EndTime:   1663945712,
			Status:    0,
			Total:     10,
			Left:      10,
		},
	)

	secInfoConf = append(
		secInfoConf,
		SecInfoConf{
			ProductId: 1027,
			StartTime: 1663942112,
			EndTime:   1663945712,
			Status:    0,
			Total:     10,
			Left:      10,
		},
	)

	secInfoConf = append(
		secInfoConf,
		SecInfoConf{
			ProductId: 1026,
			StartTime: 1663942112,
			EndTime:   1663945712,
			Status:    0,
			Total:     10,
			Left:      10,
		},
	)

	data, err := json.Marshal(secInfoConf)
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
