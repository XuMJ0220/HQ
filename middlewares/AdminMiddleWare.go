package middlewares

import (
	"HQ/controllers"
	"HQ/pkg/JWT"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	AuthorizationNil         = "请求头中Authorization字段为空"
	AuthorizationFormatError = "请求头中Authorization字段格式错误"
	TokenParseError          = "token解析失败"
	NotAdmin                 = "非管理员"
)

func AdminMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//1.从请求体的头部获取token
		v := ctx.Request.Header.Get("Authorization")
		if v == "" {
			ctx.JSON(http.StatusOK, controllers.CodeMsgDetail(controllers.CodeHeaderAuthFailed, AuthorizationNil))
			ctx.Abort()
			return
		}
		token := strings.SplitN(v, " ", 2)
		if token[0] != "Bearer" {
			ctx.JSON(http.StatusOK, controllers.CodeMsgDetail(controllers.CodeHeaderAuthFailed, AuthorizationFormatError))
			ctx.Abort()
			return
		}
		//2.验证token是否有效
		claims, err := JWT.ParseToken(token[1])
		if err != nil {
			ctx.JSON(http.StatusOK, controllers.CodeMsgDetail(controllers.CodeHeaderAuthFailed, TokenParseError))
			ctx.Abort()
			return
		}
		//3.验证是否为管理员
		if claims.Role != 1 {
			ctx.JSON(http.StatusOK, controllers.CodeMsgDetail(controllers.CodeHeaderAuthFailed, NotAdmin))
			ctx.Abort()
			return
		}
		//4.将claims中的信息放入上下文
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
