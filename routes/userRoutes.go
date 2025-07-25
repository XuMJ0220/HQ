package routes

import (
	"HQ/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes 用户路由
func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/user")

	userRouter.POST("/signup", controllers.UserController{}.Signup)
	userRouter.POST("/login",controllers.UserController{}.Login)
}
