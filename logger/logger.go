package logger

import (
	"flag"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	//flag解析
	flag.Parse()
	//创建有给全局的Logger
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))

}

func CreateLogger() *zap.Logger {
	t := time.Now().Format("2006-01-02")

	path := flag.Args()[0] + "./logs/" + t+".log"
	//输出到指定文件下，用lumberjack
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    200,
		MaxAge:     10,
		MaxBackups: 10,
		Compress:   false,
	})
	//日志等级
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	//指定文件配置
	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.EncodeCaller = zapcore.ShortCallerEncoder
	productionCfg.TimeKey = "时间"
	productionCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	//指定文件（这里用NewJSONEncoder是JSON格式的Encoder）
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, file, level),
	)
	return zap.New(core)
}
