package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c UserController) Register(ctx *gin.Context) {
	//进行身份验证
	//进行逻辑处理
	ctx.String(http.StatusOK, "haha")
	//数据入库
}
