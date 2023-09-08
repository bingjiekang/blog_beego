package controllers

import (
	"blog_beego/models"
	"blog_beego/utils"
	"fmt"
)

type ArticleAdd struct {
	Base
}

// 请求写博客界面
func (this *ArticleAdd) Get() {
	// 访问的时候判断是否登录
	if this.Data["IsLogin"] == true {
		this.TplName = "write_article.html"
		this.Render()
	} else {
		this.Redirect("/", 302)
	}

}

// 提交博客内容界面
func (this *ArticleAdd) Post() {
	dmate := utils.DataMat{}
	// 判断是否登录
	if this.Data["IsLogin"] == true {
		username := this.GetSession("loginuser").(string)
		fmt.Println("session记录", username)
		// username = string(username)
		// 获取标题,标签,简介,内容
		title := this.GetString("title")
		tags := this.GetString("tags")
		short := this.GetString("short")
		content := this.GetString("content")
		fmt.Println(title, tags, short, content)
		// 判断内容是否完整
		if len(title) != 0 && len(tags) != 0 && len(short) != 0 && len(content) != 0 {
			if err := models.InsertContent(title, username, tags, short, content); err != nil {
				fmt.Println("插入失败", err)
				dmate.Code = 0
				dmate.Message = "博客保存失败,请联系管理员!"
			} else {
				dmate.Code = 1
				dmate.Message = "保存成功,请查看!"
			}
		} else {
			dmate.Code = 0
			dmate.Message = "内容不完整,请补全内容!"
		}
	} else {
		dmate.Code = 0
		dmate.Message = "未登录,请先登录!"
	}
	this.Data["json"] = map[string]interface{}{"code": dmate.Code, "message": dmate.Message}
	this.ServeJSON()
	// this.TplName = "register.html"
	this.Render()
}
