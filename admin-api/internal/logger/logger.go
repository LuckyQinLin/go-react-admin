package logger

import (
	"bytes"
	"fmt"
	"io"
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
	isWriteFile bool           // 是否写文件
	wg          sync.WaitGroup // 并发分组
	out         io.Writer
	fileSize    int64
	maxFileSize int64
}

func DefaultLogger(logFile string) *ToolLogger {
	logger := &ToolLogger{
		filePath:    logFile,
		logLevel:    Info,
		logChannel:  make(chan *LogEntry, 10),
		Prefix:      "Admin",
		TimeFormat:  "2006-01-02 15:04:05",
		maxFileSize: 10 * 1024 * 1024,
		isWriteFile: true,
	}
	file, err := openFile(buildPath(logFile))
	if err != nil {
		logger.isWriteFile = false
		PrintLogger(Error, "创建文件失败: %s", err.Error())
	}
	info, _ := file.Stat()
	logger.fileSize = info.Size()
	logger.out = io.MultiWriter(file, os.Stdout)
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
		// 写入文件
		l.writeFile(entry, func() []byte {
			defer fileBuff.Reset()
			NotColor(fileBuff, l.Prefix, true)
			NotColor(fileBuff, entry.Timestamp.Format(l.TimeFormat), true)
			l.changeLevelNotColor(fileBuff, entry.Level)
			fileBuff.WriteString(entry.Message)
			fileBuff.WriteByte('\n')
			return fileBuff.Bytes()
		})
		// 写入控制台
		l.console(true, func() []byte {
			defer buff.Reset()
			Blue(buff, l.Prefix, true)
			White(buff, entry.Timestamp.Format(l.TimeFormat), true)
			l.changeLevel(buff, entry.Level)
			buff.WriteString(entry.Message)
			return buff.Bytes()
		})
	}
}

// writeFile 写文件
func (l *ToolLogger) writeFile(entry *LogEntry, handler func() []byte) {
	stat, _ := l.file.Stat()
	if entry.Timestamp.Day() != time.Now().Day() {
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
	buffer := handler()
	if _, err := l.file.Write(buffer); err != nil {
		PrintLogger(Error, "写入文件失败: %s", err.Error())
	}
}

// console 写入控制台
func (l *ToolLogger) console(isConsole bool, handler func() []byte) {
	if !isConsole {
		return
	}
	l.out.Write(handler())
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
