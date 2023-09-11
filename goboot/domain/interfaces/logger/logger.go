package logger

import "goboot/domain/model"

// ILogger 抽象日志接口定义
type ILogger interface {
	GetPrefix() string                          // 获取前缀
	PrintLogger(model.LogLevel, string, ...any) // 输出日志
	Info(string, ...any)                        //
	Debug(string, ...any)                       //
	Warn(string, ...any)                        // 警告
	Fail(string, ...any)                        // 失败
	Error(string, ...any)                       // 错误
}
