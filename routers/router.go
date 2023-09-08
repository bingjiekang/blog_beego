package routers

import (
	"blog_beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 静态路由配置
	beego.SetStaticPath("/static", "static")
	// 初始路由
	beego.Router("/", &controllers.Home{})
	// 注册路由
	beego.Router("/register", &controllers.Register{})
	// 登录界面
	beego.Router("/login", &controllers.Login{})
	// 退出
	beego.Router("/exit", &controllers.Exit{})
	// 写博客
	beego.Router("/article/add", &controllers.ArticleAdd{})
	// 显示文章详情
	// beego.Router("/article/:id", &controllers.{})
}
