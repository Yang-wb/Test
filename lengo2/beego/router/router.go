package router

import (
	"lengo2/beego/controller/IndexController"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
