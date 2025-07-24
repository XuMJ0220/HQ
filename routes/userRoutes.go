package routes

import (
	"HQ/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/user")

	userRouter.POST("/signup", controllers.UserController{}.Signup)
}
