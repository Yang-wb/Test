package controller

import "github.com/astaxie/beego"

type SkillController struct {
	beego.Controller
}

func (p *SkillController) SecKill() {
	p.Data["json"] = "sec kill"
	p.ServeJSONP()
}

func (p *SkillController) SecInfo() {
	p.Data["json"] = "sec info"
	p.ServeJSONP()
}
