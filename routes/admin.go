package routes

import (
	"HQ/controllers"
	"HQ/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {

	adminRouter := router.Group("/admin", middlewares.AdminMiddleWare())

	categoriesRouter := adminRouter.Group("/categories")

	categoriesRouter.GET("/", controllers.CategoriesController{}.GetAllCategories)     //查所有分类
	categoriesRouter.GET("/:id", controllers.CategoriesController{}.QueryOneCategory)  //查特定:id的分类
	categoriesRouter.POST("/", controllers.CategoriesController{}.AddCategory)         //增加一个分类,name在请求体里
	categoriesRouter.PUT("/:id", controllers.CategoriesController{}.UpdateCategory)    //改一个指定:id的分类,name在请求体里
	categoriesRouter.DELETE("/:id", controllers.CategoriesController{}.DeleteCategory) //删除一个特定:id的分类
}
