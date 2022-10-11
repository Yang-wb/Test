package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user")
}
func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func main() {
	server := NewHttpServer("test-server")

	//server.Route("/", handler)
	//server.Route("/home", home)
	//server.Route("/user", user)
	server.Route("POST", "/user/signup", SingUp)
	log.Fatal(server.Start(":8080"))

}

func SingUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		BizCode: 0,
		Msg:     "",
		Data:    123,
	}

	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入相应失败：%v", err)
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
