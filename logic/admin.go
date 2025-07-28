package logic

import (
	"HQ/dao/mysql"
	"HQ/logger"
	"HQ/models"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// QueryAllCategories 查询所有分类
func QueryAllCategories(categories *[]models.CategoriesParam) error {
	ctx := context.Background()
	cates, err := gorm.G[models.Category](mysql.Db).Select("id,name").Find(ctx)
	//如果查询失败
	if err != nil {
		return err
	}
	//把从数据库得到的信息转为[]models.CategoriesParam后装入*categories
	for _, value := range cates {
		*categories = append(*categories, models.CategoriesParam{
			ID:   value.ID,
			Name: value.Name,
		})
	}
	return nil
}

// QueryOneCategory 查询一个分类
func QueryOneCategory(categoryId int64, name *string) error {
	ctx := context.Background()
	cat, err := gorm.G[models.Category](mysql.Db).Select("id,name").Where("id = ?", categoryId).First(ctx)
	//如果查询失败
	if err != nil {
		logger.CreateLogger().Error("QueryOneCategory failed",
			zap.Error(err),
			zap.Int64("categoryId", categoryId))
		return err
	}
	*name = cat.Name
	return nil
}

func AddCategory(categoryName string) error {
	ctx := context.Background()
	if err := gorm.G[models.Category](mysql.Db).Create(ctx, &models.Category{
		Name: categoryName,
	}); err != nil {
		logger.CreateLogger().Error("AddCategory failed",
			zap.Error(err),
			zap.String("categoryName", categoryName))
		return err
	}
	return nil
}

// UpdateCategory 更新分类
func UpdateCategory(categoryId int64, categoryName string) error {
	ctx := context.Background()
	result := mysql.Db.WithContext(ctx).Model(&models.Category{}).Where("id = ?", categoryId).Update("name", categoryName)
	if result.Error != nil {
		logger.CreateLogger().Error("UpdateCategory failed",
			zap.Error(result.Error),
			zap.Int64("categoryId", categoryId),
			zap.String("categoryName", categoryName))
		return result.Error
	}
	// 检查是否真的更新了行
	if result.RowsAffected == 0 {
		logger.CreateLogger().Error("UpdateCategory failed, no rows affected",
			zap.Int64("categoryId", categoryId),
			zap.String("categoryName", categoryName))
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteCategory(categoryId int64) error {
	ctx := context.Background()
	_, err := gorm.G[models.Category](mysql.Db).Where("id = ?", categoryId).Delete(ctx)
	if err != nil {
		logger.CreateLogger().Error("DeleteCategory failed",
			zap.Error(err),
			zap.Int64("categoryId", categoryId))
		return err
	}
	return nil
}
