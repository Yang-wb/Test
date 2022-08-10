package main

import (
	"fmt"
	rpcdeom "leango/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdeom.Args{10, 3}, &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
