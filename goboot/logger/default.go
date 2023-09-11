package logger

import (
	"bytes"
	"fmt"
	"goboot/domain/model"
	"time"
)

func PrintLogger(level model.LogLevel, format string, args ...any) {
	buff := new(bytes.Buffer)
	defer func() {
		buff.Reset()
		buff = nil
	}()
	Blue(buff, "Admin", true)
	// 写入日志日期
	times := time.Now().Format("2006-01-02 15:04:05")
	White(buff, times, true)
	switch level {
	case model.Info:
		Blue(buff, "Info", true)
	case model.Debug:
		Green(buff, "Debug", true)
	case model.Warn:
		Yellow(buff, "Warn", true)
	case model.Fail:
		Magenta(buff, "Fail", true)
	case model.Error:
		Red(buff, "Error", true)
	default:
		White(buff, "Unknown", true)
	}
	buff.WriteString(fmt.Sprintf(format, args...))
	fmt.Println(buff.String())
}
