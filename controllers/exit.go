package controllers

// "blog_beego/controllers"

type Exit struct {
	Base
}

// 退出,删除登录标记信息
func (this *Exit) Get() {
	this.DelSession("loginuser")
	this.Redirect("/", 302)
}
