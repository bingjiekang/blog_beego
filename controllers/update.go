package controllers

import (
	"blog_beego/models"
	"blog_beego/utils"
	"fmt"
)

type Update struct {
	Base
}

func (this *Update) Get() {
	// 访问的时候判断是否登录
	if this.Data["IsLogin"] == false {
		this.Redirect("/", 302)
	} else {
		id, _ := this.GetInt("id")
		fmt.Println(id)

		//获取id所对应的文章信息
		art, err := models.SelectIdBlog(id)
		if err != nil {
			fmt.Println("查询文章信息失败", err)
		} else {
			this.Data["Title"] = art.Title
			this.Data["Tags"] = art.Tage
			this.Data["Short"] = art.Short
			this.Data["Content"] = art.Content
			this.Data["Id"] = art.Id
		}
	}
	this.TplName = "write_article.html"
	this.Render()

}

// 修改详情页
func (this *Update) Post() {
	dmate := utils.DataMat{}
	// 访问的时候判断是否登录
	if this.Data["IsLogin"] == false {
		dmate.Code = 0
		dmate.Message = "未登录,请先登录!"
	} else {
		id, _ := this.GetInt("id")
		fmt.Println(id)

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
			if err := models.UpdateContent(id, title, username, tags, short, content); err != nil {
				fmt.Println("更新失败", err)
				dmate.Code = 0
				dmate.Message = "博客更新失败,请联系管理员!"
			} else {
				dmate.Code = 1
				dmate.Message = "更新成功,请查看!"
			}
		} else {
			dmate.Code = 0
			dmate.Message = "内容不完整,请补全内容!"
		}

	}
	this.Data["json"] = map[string]interface{}{"code": dmate.Code, "message": dmate.Message}
	this.ServeJSON()
	// this.TplName = "register.html"
	this.Render()
}
