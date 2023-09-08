package controllers

import (
	"blog_beego/models"
	"blog_beego/utils"
	"fmt"
	"strconv"
)

type ArticleShow struct {
	Base
}

// 用来显示详情页
func (this *ArticleShow) Get() {
	// 获取url后面的id
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	// var art utils.Article
	art, err := models.SelectIdBlog(id)
	if err != nil {
		fmt.Println("获取id失败,请重新尝试")
		this.Redirect("/", 302)
	} else {
		// this.Data["Title"] = art.Title
		// this.Data["Content"] = art.Content
		// this.TplName = "show_article.html"
		this.Data["Title"] = art.Title
		this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
		this.TplName = "show_article.html"
	}
	this.Render()
}
