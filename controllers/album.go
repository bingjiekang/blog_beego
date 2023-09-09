package controllers

import (
	"blog_beego/models"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Album struct {
	Base
}

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
