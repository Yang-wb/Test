package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"lengo2/chat/proto"
	"net"
	"os"
)

//发送文本消息
func sendTextMessage(conn net.Conn, text string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserSendMessageCmd

	var sendReq proto.UserSendMessageReq
	sendReq.Data = text
	sendReq.UserId = userId

	data, err := json.Marshal(sendReq)
	if err != nil {
		return
	}

	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	packLen := uint32(len(data))

	//fmt.Println("packlen:", packLen)
	binary.BigEndian.PutUint32(buf[0:4], packLen)

	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data  failed")
		return
	}

	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}

	return
}

//进入聊天
func enterTalk(conn net.Conn) {
	//var destUserId int
	var msg string
	fmt.Println("please input text")
	fmt.Scanf("%s", &msg)
	sendTextMessage(conn, msg)
}

func listUnReadMsg() {
	select {
	case msg := <-msgChan:
		fmt.Println(msg.UserId, ":", msg.Data)
	default:
		return
	}
}

//进入菜单
func enterMenu(conn net.Conn) {
	fmt.Println("1. list online user")
	fmt.Println("2. talk")
	fmt.Println("3. list message")
	fmt.Println("4. exit")

	var sel int
	fmt.Scanf("%d\n", &sel)
	switch sel {
	case 1:
		//在线用户
		outputUserOnline()
	case 2:
		//聊天会话
		enterTalk(conn)
	case 3:
		//显示未读的消息
		listUnReadMsg()
		return
	case 4:
		//退出
		os.Exit(0)
	}
}

func logic(conn net.Conn) {
	enterMenu(conn)
}
