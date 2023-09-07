package main

import (
	_ "blog_beego/routers"
	_ "blog_beego/utils"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
