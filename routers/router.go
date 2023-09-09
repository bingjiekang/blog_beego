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
	beego.Router("/article/:id", &controllers.ArticleShow{})
	// 修改文章
	beego.Router("/article/update", &controllers.Update{})
	// 删除文章
	beego.Router("/article/delete", &controllers.DeleteBlog{})
	// 标签查看
	beego.Router("/tags", &controllers.Tag{})
	// 关于相册
	beego.Router("/album", &controllers.Album{})
	beego.Router("/upload", &controllers.Album{})
	// 关于自身
	beego.Router("/aboutme", &controllers.AboutMe{})
}
