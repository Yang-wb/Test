package main

import (
	"fmt"
	"net"
)

var (
	localIpArray []string
)

func init() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic(fmt.Sprintf("get local op failed, %v", err))
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIpArray = append(localIpArray, ipnet.IP.String())
			}
		}
	}
	fmt.Println(localIpArray)
}
