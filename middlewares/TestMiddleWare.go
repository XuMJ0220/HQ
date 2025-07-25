package middlewares

import (
	"HQ/pkg/JWT"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TestMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//从请求头获取token
		v := ctx.Request.Header.Get("Authorization")
		if v == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请传输token",
			})
			ctx.Abort()
			return
		}
		//
		token := strings.SplitN(v, " ", 2)
		if token[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token传输格式不对",
			})
			ctx.Abort()
			return
		}
		//进行解析
		claims, err := JWT.ParseToken(token[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token解析有误",
			})
			ctx.Abort()
			return
		}
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}
