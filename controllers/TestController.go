package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func (c TestController) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"username": ctx.GetString("username"),
		"hello":    "world",
	})
}
