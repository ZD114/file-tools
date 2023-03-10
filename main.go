package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"zhangda/file-tools/config"
	"zhangda/file-tools/router"
	"zhangda/file-tools/util"
)

func main() {
	newRouter := router.InitRouter()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.GetConfig().Server.Port),
		Handler:      newRouter,
		ReadTimeout:  0,
		WriteTimeout: 0,
	}

	if err := s.ListenAndServe(); err != nil {
		util.Logger.Error("服务器启动异常", util.Any("serverError", err))
	}
}
