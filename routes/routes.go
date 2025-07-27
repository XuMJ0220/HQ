package routes

import (
	"github.com/gin-gonic/gin"

	_ "HQ/docs"

	swaggerfiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func RoutesInit() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerfiles.Handler))
	UserRoutes(router)
	TestRoutes(router)
	return router
}
