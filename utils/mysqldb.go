package utils

import "time"

// 用户的数据表
type Users struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0正常,1删除
	Createtime time.Time
}

// 返回数据格式
type DataMat struct {
	Code    int
	Message string
}

// 文章表
type Article struct {
	Id         int
	Title      string
	Author     string
	Tage       string
	Short      string
	Content    string `orm:"type(text)"` // longtext
	Createtime time.Time
}

// 照片存储表
type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	Createtime time.Time
}
