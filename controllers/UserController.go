package controllers

import (
	"HQ/logic"
	"HQ/models"
	"HQ/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	validatorPkg "github.com/go-playground/validator/v10"
)

type UserController struct {
}

// Signup 注册
// @Summary 注册新用户
// @Description  注册新用户
// @Tags user
// @Accept json
// @Produce json
// @Param username body models.RegisterParam true "用户注册"
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 200 {object} models.APIResponse
// @Router /user/signup [post]
func (c UserController) Signup(ctx *gin.Context) {
	registerParam := models.RegisterParam{}
	//进行注册信息验证
	if err := ctx.ShouldBind(&registerParam); err != nil {
		errs, ok := err.(validatorPkg.ValidationErrors)
		if !ok {
			//如果错误翻译失败了
			ctx.JSON(http.StatusOK,
				CodeMsgDetail(CodeInvalidParam, err.Error()))

		} else {
			ctx.JSON(http.StatusOK,
				CodeMsgDetail(CodeInvalidParam,
					validator.RemoveTopStruct(errs.Translate(validator.Trans))))
		}
	} else {
		//进行逻辑处理
		if userid, err := logic.Signup(registerParam); err != nil {
			// if err == errors.New("用户名已经存在") {
			// 	ctx.JSON(http.StatusOK,
			// 	CodeMsgDetail(CodeSignupFailed,err.Error()))
			// } else {
			// 	ctx.JSON(http.StatusOK,
			// 	CodeMsgDetail(CodeLoginFailed,err.Error()))
			// }
			ctx.JSON(http.StatusOK,
				CodeMsgDetail(CodeSignupFailed, err.Error()))
			return
		} else {
			//注册成功,返回给客户端信息
			ctx.JSON(http.StatusOK,
				CodeMsgDetail(CodeSignupSuccess, gin.H{
					"user_id": userid,
				}))
		}
	}
}

// Login 登录
// @Summary 用户登录
// @Description 用户登录
// @Tags user
// @Accept json
// @Produce json
// @Param username body models.LoginParam true "用户登录"
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 200 {object} models.APIResponse
// @Router /user/login [post]
func (c UserController) Login(ctx *gin.Context) {
	//进行信息验证
	loginParam := models.LoginParam{}
	if err := ctx.ShouldBind(&loginParam); err != nil {
		errs, ok := err.(validatorPkg.ValidationErrors)
		if !ok {
			//如果错误翻译失败了
			ctx.JSON(http.StatusOK,
				CodeMsgDetail(CodeInvalidParam, err.Error()))

		} else {
			ctx.JSON(http.StatusOK,
				CodeMsgDetail(CodeInvalidParam,
					validator.RemoveTopStruct(errs.Translate(validator.Trans))))
		}
	} else {
		var role int8
		//登录逻辑处理
		token, err := logic.Login(loginParam, &role)
		if err != nil {
			// if err == errors.New("用户名或密码错误") {
			// 	ctx.JSON(http.StatusOK,
			// 		CodeMsgDetail(CodeLoginFailed, "用户名或密码错误"))
			// 	return
			// } else {
			// 	ctx.JSON(http.StatusOK,
			// 		CodeMsgDetail(CodeLoginFailed, err.Error()))
			// 	return
			// }
			ctx.JSON(http.StatusOK,
				CodeMsgDetail(CodeLoginFailed, err.Error()))
			return
		} else {
			// 是普通用户
			if role == 0 {
				ctx.JSON(http.StatusOK,
					CodeMsgDetail(CodeLoginSuccess, gin.H{
						"role":  "普通用户",
						"token": token,
					}))
			} else {
				ctx.JSON(http.StatusOK,
					CodeMsgDetail(CodeAdminLogin, gin.H{
						"role":  "管理员",
						"token": token,
					}))
			}
		}
	}
}
