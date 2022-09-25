package controller

import (
	"lengo2/SecKill/SecProxy/service"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type SkillController struct {
	beego.Controller
}

func (p *SkillController) SecKill() {
	result := make(map[string]interface{})
	result["code"] = 0
	result["message"] = "success"

	defer func() {
		p.Data["josn"] = result
		p.ServeJSONP()
	}()

	productId, err := p.GetInt("product_id")
	if err != nil {
		result["code"] = 1001
		result["message"] = "invalid prodcut_id"
		return
	}

	source := p.GetString("stc")
	authcode := p.GetString("authcode")
	secTime := p.GetString("time")
	nance := p.GetString("nance")

	secRequest := &service.SecRequest{
		ProductId:    productId,
		Source:       source,
		AuthCode:     authcode,
		SecTime:      secTime,
		Nance:        nance,
		UserAuthSign: p.Ctx.GetCookie("userAuthSing"),
	}
	secRequest.UserId, err = strconv.Atoi(p.Ctx.GetCookie("userId"))
	if err != nil {
		result["code"] = service.ErrInvalidRequest
		result["message"] = err.Error()

		return
	}
	secRequest.AccessTime = time.Now()
	if len(p.Ctx.Request.RemoteAddr) > 0 {
		secRequest.ClientAddr = strings.Split(p.Ctx.Request.RemoteAddr, ":")[0]
	}
	secRequest.ClientRefence = p.Ctx.Request.Referer()

	data, code, err := service.SecKill(secRequest)
	if err != nil {
		result["code"] = code
		result["message"] = err.Error()

		return
	}
	result["code"] = code
	result["data"] = data
	return
}

func (p *SkillController) SecInfo() {
	productId, err := p.GetInt("product_id")
	result := make(map[string]interface{})
	result["code"] = 0
	result["message"] = "success"

	defer func() {
		p.Data["josn"] = result
		p.ServeJSONP()
	}()

	if err != nil {
		//result["code"] = 1001
		//result["message"] = "invalid product_id"
		//
		//logs.Error("invalid request, get product_id failed, err:%v", err)
		data, code, err := service.SecInfoList()
		if err != nil {
			result["code"] = code
			result["message"] = err.Error()

			logs.Error("invalid request, get productList failed, err:%v", err)
			return
		}
		result["code"] = code
		result["data"] = data
	} else {
		data, code, err := service.SecInfo(productId)
		if err != nil {
			result["code"] = code
			result["message"] = err.Error()

			logs.Error("invalid request, get product_id failed, err:%v", err)
			return
		}
		result["code"] = code
		result["data"] = data
	}
}
