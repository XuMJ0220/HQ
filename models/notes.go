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
	Author   User     `json:"author" gorm:"foreignKey:AuthorID;references:UserID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}

// CreateNoteParam 创建笔记参数
type CreateNoteParam struct {
	Title      string `json:"title" binding:"required"`
	ContentMD  string `json:"content_md" binding:"required"`
	CategoryID int64  `json:"category_id" binding:"required"`
	Status     uint8  `json:"status" binding:"oneof=0 1"`
}

type UpdateNoteParam struct{
	Title      string `json:"title"`
	ContentMD  string `json:"content_md"`
	CategoryID int64  `json:"category_id"`
	Status     *uint8  `json:"status" ` //用指针来判断是否更新
}

type NoteResponse struct {
	ID           int64     `json:"id,string"`
	AutherName   string    `json:"author_name"`
	CategoryName string    `json:"category_name"`
	Title        string    `json:"title"`
	ContendMD    string    `json:"contend_md"`
	ContendHTML  string    `json:"contend_html"`
	Status       uint8     `json:"status"`
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
}

// TableName 指定表名
func (Note) TableName() string {
	return "notes"
}
