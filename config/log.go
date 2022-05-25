package config

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func GetLog() *zap.Logger {
	if Logger == nil {
		core := zapcore.NewCore(getEncoder(), getWriteSyncer(), zapcore.DebugLevel)
		return zap.New(core, zap.AddCaller())
	}
	return Logger
}

func getEncoder() zapcore.Encoder {
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConf)
}

func getWriteSyncer() zapcore.WriteSyncer {
	file, err := os.OpenFile(GetConf().GetString("serve.log"), os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic("Fail to open file, err: " + err.Error())
	}
	return zapcore.AddSync(file)
}
