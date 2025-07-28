package logic

import (
	"HQ/dao/mysql"
	"HQ/models"
	"context"

	"gorm.io/gorm"
)

// QueryAllCategories 查询所有分类
func QueryAllCategories(categories *[]models.CategoriesParam) error {
	ctx := context.Background()
	cates, err := gorm.G[models.Categories](mysql.Db).Select("id,name").Find(ctx)
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
func QueryOneCategory(category *models.CategoriesParam) error {
	ctx := context.Background()
	cat, err := gorm.G[models.Categories](mysql.Db).Select("id,name").Where("id = ?", (*category).ID).First(ctx)
	//如果查询失败
	if err != nil {
		return err
	}
	//把从数据库得到的信息转为models.CategoriesParam后装入*category
	(*category).ID = cat.ID
	(*category).Name = cat.Name
	return nil
}

func AddCategory(category models.CategoriesParam) error {
	ctx := context.Background()
	if err := gorm.G[models.Categories](mysql.Db).Create(ctx, &models.Categories{
		Name: category.Name,
	}); err != nil {
		return err
	}
	return nil
}

func UpdateCategory(category models.CategoriesParam) error {
	ctx := context.Background()
	result := mysql.Db.WithContext(ctx).Model(&models.Categories{}).Where("id = ?", category.ID).Update("name", category.Name)
	if result.Error != nil {
		return result.Error
	}
	// 检查是否真的更新了行
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteCategory(category models.CategoriesParam) error {
	ctx := context.Background()
	_, err := gorm.G[models.Categories](mysql.Db).Where("id = ?", category.ID).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
