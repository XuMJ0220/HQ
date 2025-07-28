package logic

import (
	"HQ/dao/mysql"
	"HQ/models"
	"HQ/pkg/JWT"
	"HQ/pkg/snowflake"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var (
	TokenDuration time.Duration = 24 * time.Hour
)

// Signup 注册
func Signup(registerParam models.RegisterParam) (int64, error) {

	username := registerParam.Username
	var count int64
	//err := mysql.Db.Raw("SELECT COUNT(*) FROM users WHERE username = ?", username).Count(&count).Error
	err := mysql.Db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
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

// Login 登录
func Login(loginParam models.LoginParam, role *int8) (string, error) {

	//在Mysql中是否有对应的用户名
	ctx := context.Background()
	users, err := gorm.G[models.User](mysql.Db).
		Select("username,password,role").
		Where("username = ?", loginParam.Username).
		Find(ctx)
	if err != nil {
		return "", err
	}
	//MD5加密
	password := fmt.Sprintf("%x", md5.Sum([]byte(loginParam.Password)))
	if len(users) == 0 || users[0].Password != password {
		return "", errors.New("用户名或密码错误")
	}
	*role = users[0].Role
	//生成token
	token, err := JWT.GenLoginToken(loginParam, users[0].Role, TokenDuration)
	if err != nil {
		return "", err
	}
	return token, nil
}
