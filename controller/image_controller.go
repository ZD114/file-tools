package controller

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"net/http"
	"os"
	"zhangda/file-tools/object"
)

// FileChangeBase64 文件转base64
func FileChangeBase64(c *gin.Context) {
	image, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusInternalServerError, object.FailMsg(err.Error()))

		return
	}

	file, err := image.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, object.FailMsg(err.Error()))

		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, object.FailMsg(err.Error()))

		return
	}

	base64Content := base64.StdEncoding.EncodeToString(content)
	c.JSON(http.StatusOK, base64Content)
}

// Base64ChangeFile base64转文件
func Base64ChangeFile(c *gin.Context) {
	var req struct {
		Data string `json:"data" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, object.FailMsg(err.Error()))

		return
	}

	content, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, object.FailMsg(err.Error()))

		return
	}

	// 检查文件类型是否合法
	fileType := http.DetectContentType(content)
	if _, ok := AllowedFileTypes[fileType]; !ok {
		c.JSON(http.StatusBadRequest, object.FailMsg(err.Error()))

		return
	}

	// 检查文件大小是否合法
	if len(content) > MaxFileSize {
		c.JSON(http.StatusBadRequest, object.FailMsg(err.Error()))

		return
	}

	fileName := "base64_change_file"
	// 自定义设置权限，目前设置所有用户可读写文件
	err = os.WriteFile(fileName, content, fs.FileMode(0666))
	if err != nil {
		c.JSON(http.StatusInternalServerError, object.FailMsg(err.Error()))

		return
	}

	defer os.Remove(fileName)

	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	c.Writer.Header().Set("Content-Type", fileType)
	c.File(fileName)
}

// 定义合法的文件类型和大小限制
var (
	AllowedFileTypes = map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}
	MaxFileSize = 2 * 1024 * 1024 // 2 MB
)
