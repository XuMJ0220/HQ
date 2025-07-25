package models

import (
	"time"

	"gorm.io/gorm"
)

// RegisterParam 用户注册参数
type RegisterParam struct {
	Username   string `json:"username" form:"username" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Gender     int8   `json:"gender" form:"gender" binding:"min=0,max=2"`
	Password   string `json:"password" form:"password" binding:"required"`
	RePassword string `json:"re_password" form:"re_password" binding:"required,eqfield=Password"`
}

// LoginParam 用户登录参数
type LoginParam struct{
	
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// User 用户表对应的GORM模型
type User struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement;column:id;comment:主键ID"`
	UserID    int64          `json:"user_id" gorm:"not null;uniqueIndex:idx_user_id;column:user_id;comment:用户ID"`
	Username  string         `json:"username" gorm:"type:varchar(64);not null;uniqueIndex:idx_username;column:username;comment:用户名"`
	Password  string         `json:"-" gorm:"type:varchar(64);not null;column:password;comment:密码"`
	Email     string         `json:"email" gorm:"type:varchar(64);not null;uniqueIndex:idx_email;column:email;comment:邮箱"`
	Gender    int8           `json:"gender" gorm:"type:tinyint(4);not null;default:0;column:gender;comment:性别：0-未知，1-男，2-女"`
	CreatedAt time.Time      `json:"create_time" gorm:"autoCreateTime;column:create_time;comment:创建时间"`
	UpdatedAt time.Time      `json:"update_time" gorm:"autoUpdateTime;column:update_time;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"delete_time,omitempty" gorm:"index;column:delete_time;comment:删除时间"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
