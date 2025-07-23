package settings

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AllConfig struct {
	App   AppConfig   `mapstructure:"app"`
	Log   LogConfig   `mapstructure:"log"`
	MySQL MySQLConfig `mapstructure:"mysql"`
	Redis RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Port    uint16 `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       uint16 `mapstructure:"db"`
}

var AllCfg AllConfig

func Init() {
	//对命令行开始进行解析
	flag.Parse()
	//path得到配置文件的路径
	path := (flag.Args())[0] + "./config.yaml"
	//加载配置文件
	viper.SetConfigFile(path)
	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() error : %v\n", err)
		// zap.L().Fatal("", zap.Error(err))
	}
	//把配置反序列化到结构体中
	if err := viper.Unmarshal(&AllCfg); err != nil {
		//zap.L().Error("viper.Unmarshal error", zap.Error(err))
		fmt.Printf("viper.Unmarshal error : %v\n", err)
	}
	//监听
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件发生改变")
		//把配置反序列化到结构体中
		if err := viper.Unmarshal(&AllCfg); err != nil {
			//zap.L().Error("viper.Unmarshal error", zap.Error(err))
			fmt.Printf("viper.Unmarshal error : %v\n", err)
			fmt.Printf("AllCfg : %v\n", AllCfg)
		}
	})

}
