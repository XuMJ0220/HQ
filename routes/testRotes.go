package routes

import (
	"HQ/controllers"
	"HQ/middlewares"

	"github.com/gin-gonic/gin"
)

func TestRoutes(c *gin.Engine) {

	userRoute := c.Group("/test")
	userRoute.GET("", middlewares.TestMiddleWare(), controllers.TestController{}.Test)
}
