package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhangda/file-tools/object"
)

// SearchPage 分页查询列表
func SearchPage(ctx *gin.Context) {
	searchParam := &object.FileSearchParam{}

	if err := ctx.BindJSON(&searchParam); err != nil {
		ctx.JSON(http.StatusBadRequest, object.FailCodeMsg(400, err.Error()))

		return
	}

	if res, err := object.GetFileItems(*searchParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, object.FailCodeMsg(500, err.Error()))
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}
