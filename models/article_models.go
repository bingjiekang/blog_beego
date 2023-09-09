package models

import (
	"blog_beego/utils"
	"errors"
	"fmt"
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

// 根据tage查询信息
func SelectTag(tag string) ([]utils.Article, error) {
	_, err := Om.Raw("select * from Article where tage = ?", tag).QueryRows(&PageData)

	if err != nil {
		fmt.Println("根据tage查询失败", err)
		return []utils.Article{}, err
	} else if len(PageData) == 0 {
		fmt.Println("根据tage查询数据不存在")
		return PageData, errors.New("根据tage查询数据不存在")
	}
	return PageData, nil
}

// 返回tage的名称和对应数量
func SelectTagCout() (map[string]int, error) {
	_, err := Om.Raw("select tage,count(*) as cout from Article Group by tage;").QueryRows(&Tags)
	if err != nil {
		fmt.Println("查询指定tage出错", err)
		return map[string]int{}, err
	}
	var sult map[string]int = make(map[string]int)
	for _, v := range Tags {
		sult[v.Tage] = v.Cout
	}
	return sult, nil
}

// 查询指定id的博客信息
func SelectIdBlog(id int) (utils.Article, error) {
	err := Om.Raw("select * from Article where id = ?", id).QueryRow(&Article)
	if err != nil {
		fmt.Println("查询指定id出错", err)
		return utils.Article{}, err
	}
	return Article, nil
}
