# 使用beego搭建的博客平台
 
参考1：[GoWeb——Beego框架的使用](https://blog.csdn.net/cold___play/article/details/131125246)  
参考2：[beego官方](https://github.com/beego/beedoc/tree/master/zh-CN)  
参考3：[其他](https://github.com/rubyhan1314/Golang-100-Days/tree/master/Day38-41(beego框架开发博客系统)) 

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
1. go git xxxx
2. 创建一个mysql数据库，并在conf里的app.conf中修改数据库相关信息
3. xxxx执行
4. 访问：8080即可
5. 配置nginx，配置域名
6. 申请ssh证书

## 开发步骤

1. 建立项目目录
2. 导入静态文件资源包括 static和view里的html文件
3. 创建数据库BlogBeego
4. 链接数据库，并建立数据表，在utils里的utils.go中连接数据库，建立数据表。
5. 编写controller 里的register.go和login.go，以及对数据库的相关操作，在models里的users_models.go编写。
6. 编写controller 里 base.go和exit.go和home.go用来记录session和检测是否登录。
7. 编写controller 里article_add.go 实现博客内容的编写，以及article_models里的InsertContent实现内容的插入
8. 编写 models里的 article_models.go 实现SelectPage（分页查询）、SelectPageAll(查询博客总数)、以及SelectTag（查询tag信息）
9. 编写models里的home_models.go 实现博客的显示和翻页及tage的显示
10. 编写controller 里的update.go和delete.go实现对blok的修改和删除，编写models里的article_models 里的UpdateContent（更新），DeleteContent(删除)操作
11. 编写controller 里的tag.go 实现tag标签的访问，models里的article_models.go 里的SelectTagCout 实现名字和数量的查询
12. 编写controller 里的aboutme.go和album.go 分别实现关于自己，和上传相册，在models里的album_models.go 里编写InsertAlbum(插入数据库照片）SelectAlbum(查询照片)，注意在utils文件夹下注册一个album数据表来存储照片信息

## 重点

1. 如果"xxx/xxxx"目录下配置的文件有init函数，在main文件中使用 _ "xxx/xxxx"即可调用 "xxx/xxxx"目录下的init初始化函数(golang默认会加载所以导入的包,不管是否忽视,都会导入)

2. routers目录下route.go文件中

	```golang
	beego.SetStaticPath("/static", "static")
	左边是url路径，右边是项目下的文件目录
	```
	这一句是注册了 static 目录为静态处理的目录
	
3. beego的orm只能链接数据库，创建数据表，不能自动创建数据库，需要手动创建数据库
4. session 的使用

	```golang
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

5. 目录conf下的app.conf 是配置文件，修改mysql的连接和session都在里面

	```conf
	appname = blog_beego
	port = 8080
	runmode = dev
	
	# mysql配置
	Name = mysql
	user = root
	pwd = 12345678
	host = 127.0.0.1
	port = 3306
	dbname = BlogBeego
	
	# session配置
	sessionon = true
	sessionprovider = "file"
	sessionname = "qianfengblog"
	sessiongcmaxlifetime = 1800
	sessionproviderconfig = "./tmp"
	sessioncookielifetime = 1800
	```


6. 连接数据库，动态创建数据表

	```golang
	func init() {
		// 数据库链接，AppConfig是从app.conf配置文件中提取信息
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
	```
	
7. md5加密，对密码进行加密

	```golang
	// MD5加密
	func MD5Str(src string) string {
		return fmt.Sprintf("%x", md5.Sum([]byte(src)))
	}
	```
	
8. 翻页 数据库指定范围查询

	```mysql
	// 分页查询,查询指定页数据,page为第几页,num为查询多少条
	func SelectPage(page int, num int) ([]utils.Article, error) {
		if page < 1 {
			return []utils.Article{}, errors.New("查询页数必须大于1")
		}
		start := (page - 1) * num
		_, err := Om.Raw("select * from Article limit ?,?", start, num).QueryRows(&PageData)
		if err != nil {
			fmt.Println("分页查询失败", err)
			return []utils.Article{}, err
		} else if len(PageData) == 0 {
			fmt.Println("查询数据不存在")
			return PageData, errors.New("查询页数超范围,或不存在,请检查后再次查询")
		}
		return PageData, nil
	
	}
	
	// 查询博客数据总条数,用于确定分页的范围
	func SelectPageAll() int {
		cout, err := Om.QueryTable(Article).Count()
		// tm := Om.Raw("select cout(*) from Article")
		if err != nil {
			fmt.Println("查询总数据出错", err)
			return 0
		}
		return int(cout)
	}
	```
	
	内容显示及翻页的代码实现

	```golang
	// 显示查询内容
	func MakeHomeBlocks(page []utils.Article, isLogin bool) template.HTML {
		// 返回的界面内容
		htmlHome := ""
		homepage := HomeBlockParam{}
		for _, val := range page {
			// 将数据库model转换为首页模板所需要的model
			homepage.Id = val.Id
			homepage.Title = val.Title
			homepage.Tags = createTagsLinks(val.Tage)
			homepage.Short = val.Short
			homepage.Content = val.Content
			homepage.Author = val.Author
			homepage.CreateTime = val.Createtime.Format("2006-01-02 15:04:05")
			homepage.Link = "/article/" + strconv.Itoa(val.Id)
			homepage.UpdateLink = "/article/update?id=" + strconv.Itoa(val.Id)
			homepage.DeleteLink = "/article/delete?id=" + strconv.Itoa(val.Id)
			homepage.IsLogin = isLogin
	
			//处理变量
			//ParseFile解析该文件，用于插入变量
			t, _ := template.ParseFiles("views/block/home_block.html")
			buffer := bytes.Buffer{}
			//就是将html文件里面 空白框 替换为待显示的数据
			t.Execute(&buffer, homepage)
			htmlHome += buffer.String()
		}
		// 返回整个界面
		return template.HTML(htmlHome)
	}
	
	// 将tags字符串转化成首页模板所需要的数据结构
	func createTagsLinks(tags string) []TagLink {
		var tagLink []TagLink
		tagsPamar := strings.Split(tags, "&")
		for _, tag := range tagsPamar {
			tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
		}
		return tagLink
	}
	
	// -----------翻页-----------
	// page是当前的页数
	func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
		pageCode := HomeFooterPageCode{}
		//查询出总的条数
		num := SelectPageAll()
		//从配置文件中读取每页显示的条数
		pageRow, _ := beego.AppConfig.Int("pagenum")
		//计算出总页数
		allPageNum := (num-1)/pageRow + 1
	
		// 当前页和可以显示的页数
		pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)
	
		//当前页数小于等于1，那么上一页的按钮不能点击
		if page <= 1 {
			pageCode.HasPre = false
		} else {
			pageCode.HasPre = true
		}
	
		//当前页数大于等于总页数，那么下一页的按钮不能点击
		if page >= allPageNum {
			pageCode.HasNext = false
		} else {
			pageCode.HasNext = true
		}
		pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
		pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
		return pageCode
	}
	
	```
	
9. golang实现markdown语法解析
	
	前端实现,在show_article.html页面上导入样式包：
	
	```html
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    ...
	    <link href="../static/css/lib/highlight.css" rel="stylesheet">
	</head>
	```
	后端实现
	
	```golang
	go get github.com/russross/blackfriday
	go get github.com/PuerkitoBio/goquery
	go get github.com/sourcegraph/syntaxhighlight
	
	func SwitchMarkdownToHtml(content string) template.HTML {
		markdown := blackfriday.MarkdownCommon([]byte(content))
	
		//获取到html文档
		doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
		/**
		对document进程查询，选择器和css的语法一样
		第一个参数：i是查询到的第几个元素
		第二个参数：selection就是查询到的元素
		 */
		doc.Find("code").Each(func(i int, selection *goquery.Selection) {
			light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
			selection.SetHtml(string(light))
			fmt.Println(selection.Html())
			fmt.Println("light:", string(light))
			fmt.Println("\n\n\n")
		})
		htmlString, _ := doc.Html()
		return template.HTML(htmlString)
	}
	
	```
	
10. 注意，更新和删除数据库时一定要指定对应id 否则会是灾难性后果！！！

	```mysql
	// 更新博客内容
	func UpdateContent(id int, title, author, tags, short, content string) error {
		_, err := Om.Raw("update Article set title=?,author=?,tage=?,short=?,content=?,createtime=? where id = ?", title, author, tags, short, content, time.Now(), id).Exec()
		if err != nil {
			// fmt.Println(err)
			return err
		}
		return nil
	}
	
	// 删除指定id博客信息
	func DeleteContent(id int) (bool, error) {
		_, err := Om.Raw("delete from Article where id = ?", id).Exec()
		if err != nil {
			// fmt.Println(err)
			return false, err
		}
		return true, nil
	}
	
	
	```
	
11. 实现文件的上传和访问(保存到数据库，拷贝到本地)
	
	```
	// 显示相册
	func (this *Album) Get() {
		albums, err := models.SelectAlbum()
		if err != nil {
			fmt.Println("显示照片错误", err)
		}
		this.Data["Album"] = albums
		this.TplName = "album.html"
		this.Render()
	}
	
	// 上传照片
	func (this *Album) Post() {
	
		fmt.Println("fileupload...")
		fileData, fileHeader, err := this.GetFile("upload")
		if err != nil {
			this.responseErr(err)
			fmt.Println("获取待上传照片失败", err)
			return
		}
		fmt.Println("name:", fileHeader.Filename, fileHeader.Size)
		fmt.Println(fileData)
		now := time.Now()
		fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
		fileType := "other"
		//判断后缀为图片的文件，如果是图片我们才存入到数据库中
		fileExt := filepath.Ext(fileHeader.Filename)
		if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
			fileType = "img"
		}
		//文件夹路径
		fileDir := fmt.Sprintf("static/img/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
		//ModePerm是0777，这样拥有该文件夹路径的执行权限
		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			this.responseErr(err)
			fmt.Println("打开照片文件路径失败", err)
			return
		}
		//文件路径
		timeStamp := time.Now().UnixNano()
		fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
		filePathStr := filepath.Join(fileDir, fileName)
		desFile, err := os.Create(filePathStr)
		if err != nil {
			this.responseErr(err)
			fmt.Println("创建照片失败", err)
			return
		}
		//将浏览器客户端上传的文件拷贝到本地路径的文件里面
		_, err = io.Copy(desFile, fileData)
		if err != nil {
			this.responseErr(err)
			fmt.Println("拷贝照片到本地失败", err)
			return
		}
		if fileType == "img" {
			// album := models.Album{0, filePathStr, fileName, 0, timeStamp}
			// 插入照片
			models.InsertAlbum(filePathStr, fileName, 0)
		}
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功"}
		this.ServeJSON()
	}
	
	func (this *Album) responseErr(err error) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
		this.ServeJSON()
	}
	
	```


	
	
	
	
