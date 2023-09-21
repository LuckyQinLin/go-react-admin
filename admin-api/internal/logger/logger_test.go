package logger

import (
	"testing"
	"time"
)

func TestLogger_changeLevel(t *testing.T) {
	//logger := DefaultLogger("~/CodeHome/GolangProject/good-tools/logs")
	//
	logger := DefaultLogger("C:\\Users\\xcmbz\\Desktop\\test_log")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			logger.Debug("一个简单的日志框架 一个简单的日志框架 一个简单的日志框架 一个简单的日志框架 => 测试日志%d", i)
		} else if i%3 == 0 {
			logger.Fail("一个简单的日志框架 一个简单的日志框架 一个简单的日志框架 一个简单的日志框架 => 失败日志%d", i)
		} else if i%5 == 0 {
			logger.Error("一个简单的日志框架 一个简单的日志框架 一个简单的日志框架 一个简单的日志框架 => 错误日志%d", i)
		} else {
			logger.Info("一个简单的日志框架 一个简单的日志框架 一个简单的日志框架 一个简单的日志框架 一个简单的日志框架 => 详情日志%d", i)
		}
	}

	time.Sleep(5 * time.Second)
}
