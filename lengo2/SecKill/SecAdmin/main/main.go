package main

import (
	"fmt"
	_ "lengo2/SecKill/SecAdmin/router"

	"github.com/astaxie/beego"
)

func main() {
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init database failed, err:%v", err))
		return
	}
	beego.Run()
}
