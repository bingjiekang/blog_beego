package controllers

import (
	"blog_beego/models"
	"fmt"
)

type DeleteBlog struct {
	Base
}

// 删除指定blok
func (this *DeleteBlog) Get() {

	if this.Data["IsLogin"] == true {
		id, _ := this.GetInt("id")
		fmt.Println("删除 id:", id)
		_, err := models.DeleteContent(id)
		if err != nil {
			fmt.Println("删除id失败", err)
		}
	}
	this.Redirect("/", 302)
}
