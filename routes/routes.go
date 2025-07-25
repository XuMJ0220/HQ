package routes

import (
	"github.com/gin-gonic/gin"
)

func RoutesInit() *gin.Engine {
	router := gin.Default()
	UserRoutes(router)
	TestRoutes(router)
	return router
}
