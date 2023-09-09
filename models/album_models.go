package models

import (
	"blog_beego/utils"
	"fmt"
	"time"
)

// 插入照片
func InsertAlbum(filepath, filename string, status int) bool {
	_, err := Om.Raw("insert into Album(filepath,filename,status,createtime) values(?,?,?,?)", filepath, filename, status, time.Now()).Exec()
	if err != nil {
		fmt.Println("保存照片到数据库失败", err)
		return false
	}
	return true

}

// 查询照片
func SelectAlbum() ([]utils.Album, error) {
	_, err := Om.Raw("select * from Album").QueryRows(&Albums)
	if err != nil {
		fmt.Println("查询照片失败", err)
		return []utils.Album{}, err
	}
	return Albums, nil
}
