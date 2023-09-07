package models

import (
	"blog_beego/utils"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

var users utils.Users = utils.Users{}

// 用om来操作数据库,不切换数据库,只使用一次
var om orm.Ormer = orm.NewOrm()

// 查询用户姓名是否存在
func SelectUsersExist(username string) bool {
	err := om.Raw("select * from users where username = ?", username).QueryRow(&users)
	if err != nil {
		return false
	}
	return true
}

// 把信息插入到users表中
func InsertUsers(username, password string) error {
	_, err := om.Raw("insert into users(username,password,status,createtime) values(?,?,?,?)", username, password, 0, time.Now()).Exec()
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}

// 查询用户名和对应密码是否匹配
func ContrastUserPwd(username, password string) bool {
	err := om.Raw("select password from users where username = ? ", username).QueryRow(&users)
	if users.Password == password && users.Username == username {
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
