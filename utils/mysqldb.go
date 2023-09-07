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
