package controllers

import (
	"blog_beego/models"
	"blog_beego/utils"
	"fmt"

	"github.com/astaxie/beego"
)

type Home struct {
	Base
}

func (this *Home) Get() {
	tag := this.GetString("tag")
	fmt.Println("tag:", tag)
	page, _ := this.GetInt("page")
	var artList []utils.Article

	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		this.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		// 获得每页待分页个数 设置分页
		page_nums, _ := beego.AppConfig.Int("pagenum")
		artList, _ = models.SelectPage(page, page_nums)

		// 用来查询总页数和当前页数的关系
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}
	// 显示是否登录,和是谁登录
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)

	// 指定页数的数据打包成一个html返回给this.Data["Content"]
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)

	this.TplName = "home.html"
	this.Render()
}
