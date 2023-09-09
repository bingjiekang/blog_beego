package controllers

type AboutMe struct {
	Base
}

// 显示关于本人的信息
func (this *AboutMe) Get() {
	this.Data["Github"] = "https://github.com/bingjiekang"
	this.Data["Email"] = "https://kangbingjie@gmail.com"
	this.Data["Blok"] = "https://kangbingjie.cn"
	this.TplName = "aboutme.html"
	this.Render()
}
