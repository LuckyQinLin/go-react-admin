package logger

import (
	"bytes"
	"fmt"
	"time"
)

func PrintLogger(level LogLevel, format string, args ...any) {
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
	case Info:
		Blue(buff, "Info", true)
	case Debug:
		Green(buff, "Debug", true)
	case Warn:
		Yellow(buff, "Warn", true)
	case Fail:
		Magenta(buff, "Fail", true)
	case Error:
		Red(buff, "Error", true)
	default:
		White(buff, "Unknown", true)
	}
	buff.WriteString(fmt.Sprintf(format, args...))
	fmt.Println(buff.String())
}
