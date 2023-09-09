package models

import (
	"blog_beego/utils"
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

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

//--------------按照标签查询--------------
/*
通过标签查询首页的数据
有四种情况
	1.左右两边有&符和其他符号
	2.左边有&符号和其他符号，同时右边没有任何符号
	3.右边有&符号和其他符号，同时左边没有任何符号
	4.左右两边都没有符号

通过%去匹配任意多个字符，至少是一个
*/
func QueryArticlesWithTag(tag string) ([]utils.Article, error) {
	// sql := " where tags like '%&" + tag + "&%'"
	// sql += " or tags like '%&" + tag + "'"
	// sql += " or tags like '" + tag + "&%'"
	// sql += " or tags like '" + tag + "'"

	// sql := "select * from Article where tage = " + tag
	// fmt.Println(sql)
	//sql: like

	// tags: http&web&socket&互联网&计算机
	//       http&web
	//       web&socket&互联网&计算机
	//       web

	// http://localhost:8080?tag=web

	// %&web&%   %代表任何内容都可以匹配
	// %&web
	// web&%
	// web
	return SelectTag(tag)
}
