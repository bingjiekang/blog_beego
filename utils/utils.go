package utils

import (
	"crypto/md5"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 数据库链接,AppConfig是从app.conf配置文件中提取信息
	user := beego.AppConfig.String("user")
	pwd := beego.AppConfig.String("pwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	// 注册默认数据库，驱动为mysql, 第三个参数就是我们的数据库连接字符串(dbconn)。
	Dbconn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?" + "charset=utf8"
	// 连接数据库
	err := orm.RegisterDataBase("default", "mysql", Dbconn)

	// 创建数据表
	if err != nil {
		fmt.Println("err:", err)
	} else {
		orm.RegisterModel(new(Users))   // 创建用户数据表
		orm.RegisterModel(new(Article)) // 创建blok内容数据表
		orm.RunSyncdb("default", false, true)
		fmt.Println("创建成功")
	}

}

// MD5加密
func MD5Str(src string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(src)))
}
