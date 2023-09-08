package models

import (
	"blog_beego/utils"

	"github.com/astaxie/beego/orm"
)

// 用om来操作数据库,不切换数据库,只使用一次
var Om orm.Ormer = orm.NewOrm()

// 用户登录信息表
var Users utils.Users = utils.Users{}

// blok信息表
var Article utils.Article = utils.Article{}
