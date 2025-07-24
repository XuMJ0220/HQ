package controllers

import (
	"HQ/logic"
	"HQ/models"
	"HQ/pkg/validator"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	validatorPkg "github.com/go-playground/validator/v10"
)

type UserController struct {
}

// Signup 注册
func (c UserController) Signup(ctx *gin.Context) {
	registerParam := models.RegisterParam{}
	//进行注册信息验证
	if err := ctx.ShouldBindJSON(&registerParam); err != nil {
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
		if userid, err := logic.Signup(registerParam); err != nil {
			if err == errors.New("用户名已经存在") {
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "用户名已经存在",
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}
			return
		} else {
			//注册成功,返回给客户端信息
			ctx.JSON(http.StatusOK, gin.H{
				"msg":    "注册成功",
				"userid": userid,
			})
		}
	}
}
