package controllers

import (
	"HQ/logger"
	"HQ/logic"
	"HQ/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	CategoryAddSuccess = "添加成功"
)

type AdminController struct {
}

type CategoriesController struct {
}

// GetAllCategories 获取所有分类名称
// @Summary 获取所有分类名称
// @Description 获取所有分类名称
// @Tags categories
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /admin/categories [get]
func (c CategoriesController) GetAllCategories(ctx *gin.Context) {
	categories := []models.CategoriesParam{}
	//1.查数据库
	err := logic.QueryAllCategories(&categories)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, CodeMsgDetail(CodeQueryFailed, err.Error()))
		logger.CreateLogger().Error("GetAllCategories failed",
			zap.Error(err))
		return
	}
	//2.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeQuerySuccess, categories))
}

// QueryOneCategory 查询一个分类
// @Summary 查询一个分类
// @Description 根据ID查询特定分类
// @Tags categories
// @Param id path int true "分类ID"
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /admin/categories/{id} [get]
func (c CategoriesController) QueryOneCategory(ctx *gin.Context) {
	//1.绑定参数
	idStr := ctx.Param("id")
	categoryId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, CodeMsgDetail(CodeQueryFailed, err.Error()))
		return
	}

	//2.查询数据库
	var categoryName = ""
	if err := logic.QueryOneCategory(categoryId, &categoryName); err != nil {
		ctx.JSON(http.StatusNotFound, CodeMsgDetail(CodeQueryFailed, err.Error()))
		return
	}
	//2.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeQuerySuccess, models.CategoriesParam{
		ID:   categoryId,
		Name: categoryName,
	}))
}

// AddCategory 添加一个分类
// @Summary 添加一个分类
// @Description 创建新的分类
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.CategoryName true "分类信息" example({"name": "新分类名称"})
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse "添加成功"
// @Failure 400 {object} models.APIResponse "请求参数错误"
// @Failure 500 {object} models.APIResponse "服务器内部错误"
// @Router /admin/categories [post]
func (c CategoriesController) AddCategory(ctx *gin.Context) {
	//1.绑定数据
	categoryName := models.CategoryName{}
	if err := ctx.ShouldBindJSON(&categoryName); err != nil {
		ctx.JSON(http.StatusBadRequest, CodeMsgDetail(CodeAddFailed, err.Error()))
		return
	}
	//2.添加到数据库
	if err := logic.AddCategory(categoryName.Name); err != nil {
		ctx.JSON(http.StatusInternalServerError, CodeMsgDetail(CodeAddFailed, err.Error()))
		return
	}
	//3.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeAddSuccess, CategoryAddSuccess))
}

// UpdateCategory 更新一个分类
// @Summary 更新一个分类
// @Description 根据ID更新分类信息
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Param name body models.CategoryName true "分类名称" example({"name": "新分类名称"})
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse "更新成功"
// @Failure 400 {object} models.APIResponse "请求参数错误"
// @Failure 500 {object} models.APIResponse "服务器内部错误"
// @Router /admin/categories/{id} [put]
func (c CategoriesController) UpdateCategory(ctx *gin.Context) {
	//1.绑定参数
	idStr := ctx.Param("id")
	categoryID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, CodeMsgDetail(CodeUpdateFailed, err.Error()))
		return
	}
	categoryName := models.CategoryName{}
	if err := ctx.ShouldBindJSON(&categoryName); err != nil {
		ctx.JSON(http.StatusBadRequest, CodeMsgDetail(CodeUpdateFailed, err.Error()))
		return
	}
	//2.更新数据库
	if err := logic.UpdateCategory(categoryID, categoryName.Name); err != nil {
		ctx.JSON(http.StatusInternalServerError, CodeMsgDetail(CodeServerBusy, nil))
		return
	}
	//3.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeUpdateSuccess, "更新成功"))
}

// DeleteCategory 删除一个分类
// @Summary 删除一个分类
// @Description 根据ID删除分类
// @Tags categories
// @Param id path int true "分类ID"
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse "请求参数错误"
// @Failure 500 {object} models.APIResponse "服务器内部错误"
// @Router /admin/categories/{id} [delete]
func (c CategoriesController) DeleteCategory(ctx *gin.Context) {
	//1.绑定参数
	idStr := ctx.Param("id")
	categoryId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, CodeMsgDetail(CodeDeleteFailed, err.Error()))
		return
	}

	//2.删除数据库
	if err := logic.DeleteCategory(categoryId); err != nil {
		ctx.JSON(http.StatusInternalServerError, CodeMsgDetail(CodeDeleteFailed, err.Error()))
		return
	}
	//3.返回给客户端
	ctx.JSON(http.StatusOK, CodeMsgDetail(CodeDeleteSuccess, "删除成功"))
}
