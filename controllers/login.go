package controllers

import (
	"blog_beego/models"
	"blog_beego/utils"

	// "fmt"

	"github.com/astaxie/beego"
)

// 登录的控制器
type Login struct {
	beego.Controller
}

// 处理get请求
func (this *Login) Get() {
	this.TplName = "login.html"
	this.Render()
}

// 处理post请求
func (this *Login) Post() {
	// 定义数据格式
	var dmate utils.DataMat

	// 获取登录信息
	username := this.GetString("username")
	password := this.GetString("password")

	// 用户是否存在
	if !models.SelectUsersExist(username) {
		dmate.Code = 0
		dmate.Message = "用户不存在,请注册"
	} else { // 密码是否正确
		mdpwd := utils.MD5Str(password)
		if !models.ContrastUserPwd(username, mdpwd) {
			dmate.Code = 0
			dmate.Message = "密码不正确,请重新输入"
		} else {
			dmate.Code = 1
			dmate.Message = "登录成功!"
			this.SetSession("loginuser", username)
		}
	}
	this.Data["json"] = map[string]interface{}{"code": dmate.Code, "message": dmate.Message}
	this.ServeJSON()
	// this.TplName = "register.html"
	this.Render()
	// return
	// 返回结果 成功/失败
}
