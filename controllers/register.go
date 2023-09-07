package controllers

import (
	"blog_beego/models"
	"blog_beego/utils"
	"fmt"

	"github.com/astaxie/beego"
)

// 注册的控制器
type Register struct {
	beego.Controller
}

// 处理get请求
func (this *Register) Get() {
	this.TplName = "register.html"
	this.Render()
}

// 处理post请求
func (this *Register) Post() {
	// 定义数据格式
	var dmate utils.DataMat

	// 获取注册信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	// fmt.Println(username, password, repassword)

	// 处理注册信息
	// 判断用户名是否存在
	if !models.SelectUsersExist(username) { // 不存在则进行注册
		lengthpsd, lengthrepsd := len(password), len(repassword)
		if lengthpsd == 0 || lengthrepsd == 0 {
			dmate.Code = 0
			dmate.Message = "密码不能为空"
		} else if lengthpsd < 5 || lengthpsd > 10 || lengthrepsd < 5 || lengthrepsd > 10 {
			dmate.Code = 0
			dmate.Message = "密码必须在5~10位之间"
		} else {
			if password != repassword {
				dmate.Code = 0
				dmate.Message = "两次密码不一致,请重新输入"
			} else {
				// 密码使用md5加密算法并保存到数据库
				mdpassword := utils.MD5Str(password)
				if err := models.InsertUsers(username, mdpassword); err != nil {
					fmt.Println(err, "加入失败")
					dmate.Code = 0
					dmate.Message = "用户插入失败,请联系管理员"
				} else {
					dmate.Code = 1
					dmate.Message = "注册成功,请登录!"
				}
			}
		}
	} else {
		dmate.Code = 0
		dmate.Message = "用户名已经存在"
	}

	this.Data["json"] = map[string]interface{}{"code": dmate.Code, "message": dmate.Message}
	this.ServeJSON()
	// this.TplName = "register.html"
	this.Render()
	// return
	// 返回结果 成功/失败
}
