package main

import (
	"fmt"
	etcd_client "github.com/coreos/etcd/clientv3"
	"time"
)

func main() {

	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"192.168.56.110:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()
}
