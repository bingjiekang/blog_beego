package models

import (
	"time"
)

// 插入博客内容
func InsertContent(title, author, tags, short, content string) error {
	_, err := Om.Raw("insert into Article(title,author,tage,short,content,createtime) values(?,?,?,?,?,?)", title, author, tags, short, content, time.Now()).Exec()
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}
