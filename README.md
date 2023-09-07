# 使用beego搭建的博客平台

参考1：[千锋教育](https://github.com/rubyhan1314/Golang-100-Days/tree/master/Day38-41(beego框架开发博客系统))   
参考2：[GoWeb——Beego框架的使用](https://blog.csdn.net/cold___play/article/details/131125246)  
参考2：[beego官方](https://github.com/beego/beedoc/tree/master/zh-CN)

## 目录介绍

- conf   #配置文件目录
	- app.conf #配置文件
- controllers #控制器目录
	- default.go #默认控制器文件
- main.go #入口
- models #模型目录
- routers #路由目录
	- router.go #路由文件
- static #静态文件目录
	- css # css文件目录
	- img #图片文件目录
	- js #js文件目录
- tests #测试文件目录
	- default_test.go #默认测试文件
- views #视图目录
	- index.tpl #默认视图文件

## 使用方法


## 开发步骤

1. 建立项目目录
2. 导入静态文件资源包括 static和view里的html文件
3. 创建数据库BlogBeego
4. 链接数据库，并建立数据表
5. 编写controller 里的register.go和login.go
6. 编写controller 里 base.go和exit.go和home.go用来记录session和检测是否登录

## 重点

1. 如果"xxx/xxxx"目录下配置的文件有init函数，在main文件中使用 _ "xxx/xxxx"即可调用 "xxx/xxxx"目录下的init初始化函数(golang默认会加载所以导入的包,不管是否忽视,都会导入)

2. routers目录下route.go文件中

	```golang
	beego.SetStaticPath("/static", "static")
	左边是url路径，右边是项目下的文件目录
	```
	这一句是注册了 static 目录为静态处理的目录
	
3. beefo的orm只能链接数据库，创建数据表，不能自动创建数据库，需要手动创建数据库
4. session 的使用

	```
	// 登录的时候设置session
	this.SetSession("loginuser", username)
	// 退出的时候删除session
	this.DelSession("loginuser")
	// 通过session判断是否登录
	func (this *Base) Prepare() {
		loginuser := this.GetSession("loginuser")
		if loginuser != nil {
			this.IsLogin = true
			this.Loginuser = loginuser
		} else {
			this.IsLogin = false
		}
		this.Data["IsLogin"] = this.IsLogin
	}
	```
