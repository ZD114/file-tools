package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"sync"
	"zhangda/file-tools/object"
	"zhangda/file-tools/util"
)

// 定义全局变量
var globalWait sync.WaitGroup // 等待多个文件上传或下载完

func UploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()

	files := form.File["file"]

	for _, file := range files {
		globalWait.Add(1)

		go uploadFile(file, c)
	}

	globalWait.Wait()
}

func uploadFile(file *multipart.FileHeader, c *gin.Context) {
	defer globalWait.Done()

	var err error = nil

	// 获取文件大小，如果小于等于1M则整个文件上传，否则采用分片方式上传
	filesize := util.GetFileSize(file.Filename)

	if filesize <= object.SmallFileSize { // 小文件
		dst := fmt.Sprintf("./files/%s", file.Filename)

		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}

	} else { // 大文件，进行切片上传

	}
}
