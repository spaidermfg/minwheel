package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//初始化zapLogger
func InitZapLogger(mode string) (err error) {
	encoder := getEncoder()
	writerSyncer := getLogWriter()
	level := new(zapcore.Level)
	err = level.UnmarshalText([]byte("debug"))
	if err != nil {
		return
	}

	var core zapcore.Core
	if mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writerSyncer, level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writerSyncer, level)
	}

	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) //替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.L().Info("[init log success]")
	return nil
}

//zapcore.Field: 一组键值对参数
//编码器，解决如何写入日志
func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}

//指定日志将写到哪里
//添加日志切割归档功能，使用第三方库Lumberjack来实现
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    200,
		MaxBackups: 5,
		MaxAge:     30,
	}
	return zapcore.AddSync(lumberJackLogger)
}
