package main

import (
	"HQ/dao/mysql"
	"HQ/dao/redis"
	"HQ/logger"
	"HQ/pkg/snowflake"
	"HQ/pkg/validator"
	"HQ/routes"
	"HQ/settings"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// @title 个人网站
// @version 1.0.0
// @description 个人网站

// @contact.name Xu Mingjie
// @contact.url
// @contact.email 1044011439@qq.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

func main() {
	//初始化配置
	settings.Init()
	//初始化日志
	logger.Init()
	//初始化验证器
	if err := validator.InitTrans("zh"); err != nil {
		zap.L().Error("Init validator failed", zap.Error(err))
	}
	//初始化MySQL
	mysql.Init()
	//初始化Redis
	redis.Init()
	//注册路由
	router := routes.RoutesInit()
	//初始化雪花算法,用来实现分布式ID
	if err := snowflake.Init("2025-01-01", 1); err != nil {
		zap.L().Error("Init snowflake failed", zap.Error(err))
	}
	//初始化翻译器
	if err := validator.InitTrans("zh"); err != nil {
		zap.L().Error("Init trans failed", zap.Error(err))
	}
	//启动服务（优雅启动或停止）
	//router := gin.Default()

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router.Handler(),
	}

	go func() {
		//服务连接
		zap.L().Info("服务器启动", zap.String("地址", "http://0.0.0.0:8080"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("ListenAndServe", zap.Error(err))
		}
	}()

	//等待终止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Shutdown Server...")

	//终止信号到来后，有5秒钟的时候给服务器去处理
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
