package controllers

import (
	"blog_beego/models"
	"fmt"
)

type Tag struct {
	Base
}

func (this *Tag) Get() {
	tages, err := models.SelectTagCout()
	if err != nil {
		fmt.Println("查询标签出错,", err)
	} else {
		this.Data["Tags"] = tages
	}
	this.TplName = "tags.html"
	this.Render()

}
