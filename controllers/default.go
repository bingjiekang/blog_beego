package controllers

import (
	"github.com/astaxie/beego"
)

type Default struct {
	beego.Controller
}

func (this *Default) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}
