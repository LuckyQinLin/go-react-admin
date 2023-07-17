package logger

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"sync"
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Debug
	Warn
	Error
	Fail
)

// ILogger 日志抽象方法
type ILogger interface {
	GetPrefix() string                    // 获取前缀
	PrintLogger(LogLevel, string, ...any) // 输出日志
}

type ToolLogger struct {
	Prefix      string         // 日志前缀
	TimeFormat  string         // 时间格式
	filePath    string         // 日志文件路径
	logLevel    LogLevel       // 默认日志等级
	logChannel  chan *LogEntry // 写入日志的实体
	file        *os.File       // 对应的日志文件
	wg          sync.WaitGroup // 并发分组
	maxFileSize int64
	fileIndex   int
}

func DefaultLogger(logFile string) *ToolLogger {
	logger := &ToolLogger{
		filePath:    logFile,
		logLevel:    Info,
		logChannel:  make(chan *LogEntry, 10),
		Prefix:      "Admin",
		TimeFormat:  "2006-01-02 15:04:05",
		maxFileSize: 1 * 1024 * 1024,
		fileIndex:   1,
	}
	filePath := buildPath(logFile, logger.fileIndex)
	file, err := openFile(filePath)
	if err != nil {
		PrintLogger(Error, "创建文件失败: %s", err.Error())
	}
	logger.file = file
	logger.wg.Add(1)
	go logger.processLog()
	return logger
}

func (l *ToolLogger) GetPrefix() string {
	return l.Prefix
}

func (l *ToolLogger) PrintLogger(level LogLevel, format string, args ...any) {
	l.log(level, "", format, args...)
}

// processLog 处理日志
func (l *ToolLogger) processLog() {
	buff := new(bytes.Buffer)
	fileBuff := new(bytes.Buffer)
	defer func() {
		buff = nil
		fileBuff = nil
		_ = l.file.Close()
		l.wg.Done()
	}()
	for entry := range l.logChannel {
		Blue(buff, l.Prefix, true)
		NotColor(fileBuff, l.Prefix, true)
		// 写入日志日期
		times := entry.Timestamp.Format(l.TimeFormat)
		NotColor(fileBuff, times, true)
		White(buff, times, true)
		// 写入日志类型
		l.changeLevel(buff, entry.Level)
		l.changeLevelNotColor(fileBuff, entry.Level)
		// 写入日志内容
		fileBuff.WriteString(entry.Message)
		buff.WriteString(entry.Message)
		fileBuff.WriteByte('\n')
		// 写入文件
		l.writeFile(fileBuff, entry.Timestamp)
		// 写入控制台
		l.console(true, buff)
	}
}

// writeFile 写文件
func (l *ToolLogger) writeFile(buff *bytes.Buffer, times time.Time) {
	stat, _ := l.file.Stat()
	if times.Day() != time.Now().Day() {
		_ = l.file.Close()
		file, _ := openFile(path.Join(l.filePath, time.Now().Format("2006-01-02"), fmt.Sprintf("access_log_%03d.log", 1)))
		l.fileIndex++
		l.file = file
	}
	if stat.Size() >= l.maxFileSize {
		fmt.Println("当前文件大小：", stat.Size(), l.maxFileSize)
		// 分隔文件
		_ = l.file.Sync()
		_ = l.file.Close()

		oldName := path.Join(l.filePath, time.Now().Format("2006-01-02"), fmt.Sprintf("access_log_%03d.log", l.fileIndex))
		newName := path.Join(l.filePath, time.Now().Format("2006-01-02"), fmt.Sprintf("access_log_%03d.log", l.fileIndex+1))

		if err := os.Rename(oldName, newName); err != nil {
			PrintLogger(Error, "重命名失败: %s", err.Error())
		}
		file, err := openFile(newName)
		if err != nil {
			PrintLogger(Error, "创建文件失败: %s", err.Error())
		}
		l.fileIndex++
		l.file = file
	}
	if _, err := l.file.Write(buff.Bytes()); err != nil {
		PrintLogger(Error, "写入文件失败: %s", err.Error())
	}
	defer buff.Reset()
}

// console 写入控制台
func (l *ToolLogger) console(isConsole bool, buff *bytes.Buffer) {
	if !isConsole {
		return
	}
	fmt.Println(buff.String())
	defer buff.Reset()
}

func (l *ToolLogger) changeLevel(buffer *bytes.Buffer, level LogLevel) {
	switch level {
	case Info:
		Blue(buffer, "Info", true)
	case Debug:
		Green(buffer, "Debug", true)
	case Warn:
		Yellow(buffer, "Warn", true)
	case Fail:
		Magenta(buffer, "Fail", true)
	case Error:
		Red(buffer, "Error", true)
	default:
		White(buffer, "Unknown", true)
	}
}

func (l *ToolLogger) changeLevelNotColor(buffer *bytes.Buffer, level LogLevel) {
	switch level {
	case Info:
		buffer.WriteString("[Info] ")
	case Debug:
		buffer.WriteString("[Debug] ")
	case Warn:
		buffer.WriteString("[Warn] ")
	case Fail:
		buffer.WriteString("[Fail] ")
	case Error:
		buffer.WriteString("[Error] ")
	default:
		buffer.WriteString("[Unknown] ")
	}
}

func (l *ToolLogger) log(level LogLevel, namespace string, format string, args ...any) {
	if level < l.logLevel {
		return
	}
	l.logChannel <- NewEntry(level, namespace, fmt.Sprintf(format, args...))
}

func (l *ToolLogger) Info(format string, args ...any) {
	l.log(Info, "", format, args...)
}

func (l *ToolLogger) Debug(format string, args ...any) {
	l.log(Debug, "", format, args...)
}

func (l *ToolLogger) Warn(format string, args ...any) {
	l.log(Warn, "", format, args...)
}

func (l *ToolLogger) Fail(format string, args ...any) {
	l.log(Fail, "", format, args...)
}

func (l *ToolLogger) Error(format string, args ...any) {
	l.log(Error, "", format, args...)
}
