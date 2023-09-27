package models

import (
	"fmt"
	"time"
)

// 查询用户姓名是否存在
func SelectUsersExist(username string) bool {
	err := Om.Raw("select * from users where username = ?", username).QueryRow(&Users)
	if err != nil {
		return false
	}
	return true
}

// 把信息插入到users表中
func InsertUsers(username, password string) error {
	_, err := Om.Raw("insert into users(username,password,status,createtime) values(?,?,?,?)", username, password, 0, time.Now()).Exec()
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}

// 查询用户名和对应密码是否匹配
func ContrastUserPwd(username, password string) bool {
	err := Om.Raw("select password from users where username = ? ", username).QueryRow(&Users)
	if Users.Password == password && Users.Username == username {
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
