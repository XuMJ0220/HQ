package models

import (
	"time"

	"gorm.io/gorm"
)

type CategoriesParam struct {
	ID   int64  `uri:"id"`
	Name string `json:"name"`
}

type CategoryName struct {
	Name string `json:"name" binding:"required"`
}

type Category struct {
	ID        int64          `json:"id" gorm:"primarykey;autoIncrement;column:id;comment:主键ID"`
	Name      string         `json:"name" gorm:"type:varchar(100);not null;column:name;comment:分类名称"`
	CreatedAt time.Time      `json:"create_time" gorm:"autoCreateTime;column:create_time;comment:创建时间"`
	UpdatedAt time.Time      `json:"update_time" gorm:"autoUpdateTime;column:update_time;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"delete_time,omitempty" gorm:"index;column:delete_time;comment:删除时间"`
}

func (c *Category) TableName() string {
	return "categories"
}
