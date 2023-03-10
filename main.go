package main

import (
	"fmt"
	"net/http"
	"zhangda/file-tools/config"
	"zhangda/file-tools/router"
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
		return
	}
}
