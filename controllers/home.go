package controllers

type Home struct {
	Base
}

func (this *Home) Get() {
	this.TplName = "home.html"
	this.Render()
}
