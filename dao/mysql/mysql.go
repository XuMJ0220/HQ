package mysql

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB = nil

func Inti() {
	dsn := "root:ytfhqqkso1@tcp(127.0.0.1:3306)/hq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db = db
	if err != nil {
		zap.L().Fatal("MySQL连接失败",zap.Error(err))
	}
}
