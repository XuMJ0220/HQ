package controllers

import (
	"HQ/models"
	"HQ/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	validatorPkg "github.com/go-playground/validator/v10"
)

type UserController struct {
}

func (c UserController) Signup(ctx *gin.Context) {
	registerParam := models.RegisterParam{}
	//进行注册信息验证
	if err := ctx.ShouldBindJSON(registerParam); err != nil {
		errs, ok := err.(validatorPkg.ValidationErrors)
		if !ok {
			//如果错误翻译失败了
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})

		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": validator.RemoveTopStruct(errs.Translate(validator.Trans)),
			})

		}
	} else {
		//进行逻辑处理

		//给客户端返回个消息
		ctx.JSON(http.StatusOK, gin.H{
			"msg": registerParam,
		})
	}

}
