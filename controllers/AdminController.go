package controllers

import (
	"HQ/logic"
	"HQ/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	CategoryAddSuccess = "添加成功"
)

type AdminController struct {
}

type CategoriesController struct {
}

// GetAllCategories 获取所有分类名称
func (c CategoriesController) GetAllCategories(ctx *gin.Context) {
	categories := []models.CategoriesParam{}
	//1.查数据库
	err := logic.QueryAllCategories(&categories)
	if err != nil {
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeQueryFailed, err.Error()))
		return
	}
	//2.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeQuerySuccess, categories))
}

// QueryOneCategory 查询一个分类
func (c CategoriesController) QueryOneCategory(ctx *gin.Context) {
	//1.绑定参数
	category := models.CategoriesParam{}
	if err := ctx.ShouldBindUri(&category); err != nil {
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeQueryFailed, err.Error()))
		return
	}
	//2.查询数据库
	if err := logic.QueryOneCategory(&category); err != nil {
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeQueryFailed, err.Error()))
		return
	}
	//2.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeQuerySuccess, category))
}

// AddCategory 添加一个分类
func (c CategoriesController) AddCategory(ctx *gin.Context) {
	//1.绑定数据
	category := models.CategoriesParam{}
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeAddFailed, err.Error()))
		return
	}
	//2.添加到数据库
	if err := logic.AddCategory(category); err != nil {
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeAddFailed, err.Error()))
		return
	}
	//3.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeAddSuccess, CategoryAddSuccess))
}

// UpdateCategory 更新一个分类
func (c CategoriesController)UpdateCategory(ctx *gin.Context){
	//1.绑定参数
	category:=models.CategoriesParam{}
	if err:=ctx.ShouldBindJSON(&category);err!=nil{
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeUpdateFailed, err.Error()))
		return
	}
	if err:=ctx.ShouldBindUri(&category);err!=nil{
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeUpdateFailed, err.Error()))
		return
	}
	//2.更新数据库
	if err:=logic.UpdateCategory(category);err!=nil{
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeUpdateFailed, err.Error()))
		return
	}
	//3.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeUpdateSuccess, "更新成功"))
}

// DeleteCategory 删除一个分类
func (c CategoriesController) DeleteCategory(ctx *gin.Context) {
	//1.绑定参数
	category := models.CategoriesParam{}
	if err := ctx.ShouldBindUri(&category); err != nil {
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeDeleteFailed, err.Error()))
		return
	}
	//2.删除数据库
	if err := logic.DeleteCategory(category); err != nil {
		ctx.JSON(http.StatusOK, CodeMsgDetail(CodeDeleteFailed, err.Error()))
		return
	}
	//3.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeDeleteSuccess, "删除成功"))
}
