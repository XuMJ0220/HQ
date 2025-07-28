package models

import (
	"time"

	"gorm.io/gorm"
)

// Note GORM笔记表模型
type Note struct {
	ID          int64          `json:"id" gorm:"primaryKey;autoIncrement;column:id;comment:笔记id"`
	AuthorID    int64          `json:"author_id" gorm:"not null;column:author_id"`
	CategoryID  int64          `json:"category_id" gorm:"not null;column:category_id"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null;column:title"`
	ContentMD   string         `json:"content_md" gorm:"type:text;not null;column:content_md"`
	ContentHTML string         `json:"content_html" gorm:"type:text;not null;column:content_html"`
	Status      uint8          `json:"status" gorm:"type:tinyint;not null;column:status"`
	CreatedAt   time.Time      `json:"create_time" gorm:"autoCreateTime;column:create_time;comment:创建时间"`
	UpdatedAt   time.Time      `json:"update_time" gorm:"autoUpdateTime;column:update_time;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `json:"delete_time,omitempty" gorm:"index;column:delete_time;comment:删除时间"`

	//建立关联关系
	Author   User     `json:"author" gorm:"foreignKey:AuthorID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}

// TableName 指定表名
func (Note) TableName() string {
	return "notes"
}
