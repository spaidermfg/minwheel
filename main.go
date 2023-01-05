package main

import (
	"minwheel/logger"

	"go.uber.org/zap"
)

//logger使用
func init() {
	if err := logger.InitZapLogger("dev"); err != nil { //初始化日志
		zap.L().Error("[init log fail]")
		return
	}
	defer zap.L().Sync() //将缓存中的日志同步到日志文件中
}

func main() {

}
