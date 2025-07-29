package logic

import (
	"HQ/dao/mysql"
	"HQ/logger"
	"HQ/models"
	"bytes"
	"context"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/renderer/html"
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

// DeleteCategory 删除分类
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

// CreateNote 创建笔记
func CreateNote(createNoteParam models.CreateNoteParam, authorID int64) (models.Note, error) {
	var htmlBuffer bytes.Buffer
	//用glodmark渲染
	markdown := goldmark.New(
		// 启用代码高亮扩展
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				// 你可以选择不同的高亮主题, 比如 "github", "monokai", "dracula"
				highlighting.WithStyle("dracula"),
				highlighting.WithFormatOptions(),
			),
		),
		// 配置渲染器选项
		goldmark.WithRendererOptions(
			// 允许在Markdown中书写原生的HTML标签，比如 <div>, <br>
			// 注意：如果你允许用户提交内容，这可能有安全风险(XSS)，但对于你自己的个人网站是安全的。
			html.WithUnsafe(),
		),
	)
	if err := markdown.Convert([]byte(createNoteParam.ContentMD), &htmlBuffer); err != nil {
		logger.CreateLogger().Error("CreateNote failed",
			zap.Error(err),
			zap.String("contentMD", createNoteParam.ContentMD))
		return models.Note{}, err
	}
	htmlContent := htmlBuffer.String()
	note := models.Note{
		Title:       createNoteParam.Title,
		ContentMD:   createNoteParam.ContentMD,
		CategoryID:  createNoteParam.CategoryID,
		Status:      createNoteParam.Status,
		AuthorID:    authorID,
		ContentHTML: htmlContent,
	}
	ctx := context.Background()

	err := gorm.G[models.Note](mysql.Db).Create(ctx, &note)
	if err != nil {
		logger.CreateLogger().Error("CreateNote failed",
			zap.Error(err),
			zap.String("contentMD", createNoteParam.ContentMD))
		return models.Note{}, err
	}
	//预加载作者和分类
	mysql.Db.Preload("Author").Preload("Category").Find(&note)
	return note, nil
}

// GetNotes 获取所有笔记
func GetNotes(notes *[]models.NoteResponse) error {
	ns := []models.Note{}
	err := mysql.Db.Preload("Author").Preload("Category").Find(&ns).Error
	if err != nil {
		logger.CreateLogger().Error("GetNotes failed",
			zap.Error(err))
		return err
	}
	if len(ns) == 0 {
		return gorm.ErrRecordNotFound
	}
	for _, v := range ns {
		(*notes) = append((*notes), models.NoteResponse{
			ID:           v.ID,
			AutherName:   v.Author.Username,
			CategoryName: v.Category.Name,
			Title:        v.Title,
			ContendMD:    v.ContentMD,
			ContendHTML:  v.ContentHTML,
			CreateAt:     v.CreatedAt,
			UpdateAt:     v.UpdatedAt,
		})
	}
	return nil
}

func GetNote(id int64, noteResponse *models.NoteResponse) error {
	note := []models.Note{}
	err := mysql.Db.Preload("Author").Preload("Category").Where("id = ?", id).Find(&note).Error
	if err != nil {
		logger.CreateLogger().Error("GetNote failed",
			zap.Error(err),
			zap.Int64("id", id))
		return err
	}
	if len(note) == 0 {
		return gorm.ErrRecordNotFound
	}
	noteResponse.ID = note[0].ID
	noteResponse.Title = note[0].Title
	noteResponse.ContendMD = note[0].ContentMD
	noteResponse.ContendHTML = note[0].ContentHTML
	noteResponse.CategoryName = note[0].Category.Name
	noteResponse.AutherName = note[0].Author.Username
	noteResponse.CreateAt = note[0].CreatedAt
	noteResponse.UpdateAt = note[0].UpdatedAt
	return nil
}

func UpdateNote(id int64, updateNoteParam models.UpdateNoteParam) error {
	//1.根据id查询记录是否存在
	note := models.Note{}
	if err := mysql.Db.Where("id = ?", id).First(&note).Error; err != nil {
		logger.CreateLogger().Error("UpdateNote failed",
			zap.Error(err),
			zap.Int64("id", id))
		return err
	}
	//创建一个map来更新字段
	updataData := make(map[string]any)
	if updateNoteParam.Title != "" {
		updataData["title"] = updateNoteParam.Title
	}
	if updateNoteParam.CategoryID != 0 {
		updataData["category_id"] = updateNoteParam.CategoryID
	}
	if updateNoteParam.Status != nil {
		updataData["status"] = *(updateNoteParam.Status)
	}
	if updateNoteParam.ContentMD != "" {
		updataData["content_md"] = updateNoteParam.ContentMD

		var htmlBuffer bytes.Buffer
		//用glodmark渲染
		markdown := goldmark.New(
			// 启用代码高亮扩展
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					// 你可以选择不同的高亮主题, 比如 "github", "monokai", "dracula"
					highlighting.WithStyle("dracula"),
					highlighting.WithFormatOptions(),
				),
			),
			// 配置渲染器选项
			goldmark.WithRendererOptions(
				// 允许在Markdown中书写原生的HTML标签，比如 <div>, <br>
				// 注意：如果你允许用户提交内容，这可能有安全风险(XSS)，但对于你自己的个人网站是安全的。
				html.WithUnsafe(),
			),
		)
		if err := markdown.Convert([]byte(updateNoteParam.ContentMD), &htmlBuffer); err != nil {
			logger.CreateLogger().Error("UpdateNote failed",
				zap.Error(err),
				zap.String("contentMD", updateNoteParam.ContentMD))
			return err
		}
		htmlContent := htmlBuffer.String()

		updataData["content_html"] = htmlContent
	}

	//如果笔记没有更改，直接返回
	if len(updataData) == 0 {
		return nil
	}
	//往数据库更新
	if err := mysql.Db.Model(&models.Note{}).Where("id = ?", id).Updates(updataData).Error; err != nil {
		logger.CreateLogger().Error("UpdateNote failed",
			zap.Error(err),
			zap.Int64("id", id))
		return err
	}
	return nil
}

func DeleNote(id int64) error {
	ctx := context.Background()
	_, err := gorm.G[models.Note](mysql.Db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		logger.CreateLogger().Error("DeleNote failed",
			zap.Error(err),
			zap.Int64("id", id))
		return err
	}
	return nil
}
