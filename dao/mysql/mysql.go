package mysql

import (
	"fmt"
	"HQ/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go.uber.org/zap"
)

var Db *gorm.DB

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.AllCfg.MySQL.User,
		settings.AllCfg.MySQL.Password,
		settings.AllCfg.MySQL.Host,
		settings.AllCfg.MySQL.Port,
		settings.AllCfg.MySQL.DBName,
	)
	
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("MySQL连接失败", zap.Error(err))
		zap.L().Info("请确保MySQL服务已启动，数据库已创建")
		return
	}
	
	zap.L().Info("MySQL连接成功")
}
