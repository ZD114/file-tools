package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"zhangda/file-tools/controller"
	"zhangda/file-tools/object"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(Recovery)

	group := server.Group("/file-tools")
	{
		group.POST("/upload", controller.UploadFiles)
		group.POST("/search", controller.SearchPage)
	}

	return server
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("router", r)

			c.JSON(http.StatusOK, object.FailMsg("系统内部错误"))
		}
	}()
	c.Next()
}
