package settings

import (
	"HQ/logger"
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() {
	//对命令行开始进行解析
	//flag.Parse()
	//path得到配置文件的路径
	path := (flag.Args())[0] + "./config.yaml"
	//加载配置文件
	viper.SetConfigFile(path)
	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		//fmt.Printf("viper.ReadInConfig() error : %v\n", err)
		zap.L().Fatal("", zap.Error(err))
	}
	//监听
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件发生改变")
	})
	sugar:=logger.CreateLogger().Sugar()
	sugar.Info("nodudu")
}
