package controllers

import "github.com/astaxie/beego"

// 用于记录session判断用户是否登录
type Base struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

// 判断是否登录
func (this *Base) Prepare() {
	loginuser := this.GetSession("loginuser")
	if loginuser != nil {
		this.IsLogin = true
		this.Loginuser = loginuser
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}
