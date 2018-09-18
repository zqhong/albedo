package app

import (
	"github.com/zqhong/albedo/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger() {
	logLevel := zap.InfoLevel
	if util.IsDebug() {
		logLevel = zap.DebugLevel
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "runtime/log/albedo-zap.log",
		MaxSize:    50, // megabytes
		MaxBackups: 20,
		MaxAge:     30, // days
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		logLevel,
	)

	Logger = zap.New(core)
	defer Logger.Sync()
}
