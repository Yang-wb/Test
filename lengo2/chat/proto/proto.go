package proto

import "lengo2/chat/common"

//消息
type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

//登录协议
type LoginCmd struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}

//注册协议
type RegisterCmd struct {
	User common.User `json:"user"`
}

//登录返回 成功返回用户列表
type LoginCmdRes struct {
	Code  int    `json:"code"`
	User  []int  `json:"users"`
	Error string `json:"error"`
}

//用户状态通知
type UserStatusNotify struct {
	UserId int `json:"user_id"`
	Status int `json:"user_status"`
}

//发送消息
type UserSendMessageReq struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}

//接收消息
type UserRecvMessageReq struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}
