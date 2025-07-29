package routes

import (
	"HQ/controllers"
	"HQ/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1/admin", middlewares.AdminMiddleWare())

	categoriesRouter := v1.Group("/categories")

	categoriesRouter.GET("/", controllers.CategoriesController{}.GetAllCategories)     //查所有分类
	categoriesRouter.GET("/:id", controllers.CategoriesController{}.QueryOneCategory)  //查特定:id的分类
	categoriesRouter.POST("/", controllers.CategoriesController{}.AddCategory)         //增加一个分类,name在请求体里
	categoriesRouter.PUT("/:id", controllers.CategoriesController{}.UpdateCategory)    //改一个指定:id的分类,name在请求体里
	categoriesRouter.DELETE("/:id", controllers.CategoriesController{}.DeleteCategory) //删除一个特定:id的分类

	notesRouter := v1.Group("notes")

	notesRouter.POST("/", controllers.NotesController{}.AddNote) //添加一则笔记
	notesRouter.GET("/", controllers.NotesController{}.GetAllNotes) //查询所有笔记
	notesRouter.GET("/:id",controllers.NotesController{}.GetOneNote) //查询一条笔记
	notesRouter.PUT("/:id",controllers.NotesController{}.UpdateNote) //更新一条指定:id的笔记
	notesRouter.DELETE("/:id", controllers.NotesController{}.DeleteNote) //删除一条指定:id的笔记
}
