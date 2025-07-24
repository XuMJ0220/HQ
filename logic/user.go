package logic

import (
	"HQ/dao/mysql"
	"HQ/models"
	"HQ/pkg/snowflake"
	"crypto/md5"
	"errors"
	"fmt"
)

// Signup 注册
func Signup(registerParam models.RegisterParam) (int64, error) {

	username := registerParam.Username
	var count int64
	err := mysql.Db.Raw("SELECT COUNT(*) FROM users WHERE username = ?", username).Count(&count).Error
	//如果mysql查询失败
	if err != nil {
		return 0, err
	}
	//用户名已经存在了
	if count > 0 {
		return 0, errors.New("用户名已经存在")
	}	
	//用雪花算法创建用户ID
	userid := snowflake.GenID()
	//对密码进行MD5加密
	password := fmt.Sprintf("%x", md5.Sum([]byte(registerParam.Password)))
	//创建用户
	user := models.User{
		UserID:   userid,
		Username: username,
		Password: password,
		Email:    registerParam.Email,
		Gender:   registerParam.Gender,
	}
	//数据插入Mysql
	err = mysql.Db.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return userid, nil
}
