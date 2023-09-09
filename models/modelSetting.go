package models

import (
	"blog_beego/utils"

	"github.com/astaxie/beego/orm"
)

// 用来控制首页显示内容
type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

// 标签链接
type TagLink struct {
	TagName string
	TagUrl  string
}

// 用来分页的结构体
type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

// 查询tage的数量和类型
type TageCount struct {
	Tage string
	Cout int
}

// 用om来操作数据库,不切换数据库,只使用一次
var Om orm.Ormer = orm.NewOrm()

// 用户登录信息表
var Users utils.Users = utils.Users{}

// blok信息表
var Article utils.Article = utils.Article{}

// 分页查询的blok内容列表
var PageData []utils.Article = []utils.Article{}

// 查询tage的数量类型
var Tags []TageCount = []TageCount{}

// 查询照片的类型
var Albums []utils.Album = []utils.Album{}
