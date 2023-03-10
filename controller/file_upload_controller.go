package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"sync"
	"zhangda/file-tools/object"
)

// 定义全局变量
var globalWait sync.WaitGroup // 等待多个文件上传或下载完

// UploadFiles 上传多个文件
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
	filesize := file.Size
	dst := fmt.Sprintf("./files/%s", file.Filename)

	if filesize <= object.SmallFileSize { // 小文件

		// 如果 path 路径不存在，会有 err，然后通过 IsNotExist 判定文件路径是否存在，如果 true 则不存在，注意用 os.ModePerm 这样文件是可以写入的
		if _, err := os.Stat("./files"); os.IsNotExist(err) {
			// mkdir 创建目录，mkdirAll 可创建多层级目录
			os.MkdirAll("./files", os.ModePerm)
		}

		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}

	} else { // 大文件，进行切片上传

		err = object.SliceUpload(file)
	}

	//断点续传标志
	var rangeExt = c.GetHeader("Range")

	if rangeExt != "" {
		//断点续传
		err = object.BreakPointTrans(dst)
		if err != nil {
			return
		}
	}

	if err != nil {
		log.Printf("上传%s文件失败", file.Filename)
	}
}
