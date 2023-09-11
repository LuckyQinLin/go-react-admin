package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
	Info(string, ...any)                  //
	Debug(string, ...any)                 //
	Warn(string, ...any)                  // 警告
	Fail(string, ...any)                  // 失败
	Error(string, ...any)                 // 错误
}

type ToolLogger struct {
	Prefix       string         // 日志前缀
	TimeFormat   string         // 时间格式
	filePath     string         // 日志文件路径
	logLevel     LogLevel       // 默认日志等级
	logChannel   chan *LogEntry // 写入日志的实体
	isWriteFile  bool           // 是否写文件
	consoleOut   io.Writer      // 控制台输出
	consoleAsync bool           // 控制台输出是否异步
	consoleBuff  *bytes.Buffer  // 控制台写buff
	fileOut      io.Writer      // 文件输出
	file         *os.File       // 文件句柄
	fileBuff     *bytes.Buffer  // 文件写buff
	maxFileSize  int64          // 日志文件阈值
	wg           sync.WaitGroup // 并发分组
}

func DefaultLogger(logFile string) *ToolLogger {
	var (
		logger *ToolLogger
		file   *os.File
		err    error
	)
	logger = &ToolLogger{
		filePath:     logFile,
		logLevel:     Info,
		logChannel:   make(chan *LogEntry, 10),
		Prefix:       "Admin",
		TimeFormat:   "2006-01-02 15:04:05",
		maxFileSize:  1 * 1024 * 1024,
		isWriteFile:  true,
		consoleAsync: false,
		consoleBuff:  new(bytes.Buffer),
		consoleOut:   os.Stdout,
	}
	if file, err = openFile(buildPath(logFile)); err != nil {
		PrintLogger(Error, "创建文件失败: %s", err.Error())
		logger.isWriteFile = false
	} else {
		logger.file = file
		logger.fileBuff = new(bytes.Buffer)
		logger.fileOut = io.MultiWriter(file)
	}
	// 检测文件
	logger.checkFile(file)
	logger.wg.Add(1)
	// 启动日志监听写入
	go logger.processLog()
	return logger
}

// checkFile 检测文件是否到达阈值
func (l *ToolLogger) checkFile(file *os.File) {
	var (
		info        os.FileInfo
		filePrefix  string
		fileNewName string
		logPath     string
		count       int
		err         error
	)
	if info, err = file.Stat(); err != nil {
		PrintLogger(Error, "无法获取文件信息:%s", err.Error())
		return
	}
	if info.Size() >= l.maxFileSize {
		PrintLogger(Warn, "日志文件大小超过阈值")
		if err := file.Close(); err != nil {
			PrintLogger(Error, "关闭当前日志文件时出错:%s", err.Error())
			return
		}
		filePrefix = fmt.Sprintf("%s_%s_", l.Prefix, time.Now().Format("20060102"))
		if count, err = CountFileNum(l.filePath, filePrefix); err != nil {
			PrintLogger(Error, "统计文件数量时出错: %s", err.Error())
			return
		}

		fileNewName = buildPathWithName(l.filePath, fmt.Sprintf("%s%d.log", filePrefix, count+1))
		logPath = buildPath(l.filePath)
		if err = os.Rename(logPath, fileNewName); err != nil {
			PrintLogger(Error, "重命名日志文件时出错: %s", err.Error())
			return
		}
		if file, err = openFile(logPath); err != nil {
			PrintLogger(Error, "创建新的日志文件时出错: %s", err.Error())
			return
		}
		l.file = file
	}
}

func (l *ToolLogger) GetPrefix() string {
	return l.Prefix
}

func (l *ToolLogger) PrintLogger(level LogLevel, format string, args ...any) {
	l.log(level, "", format, args...)
}

// processLog 处理日志
func (l *ToolLogger) processLog() {
	defer func() {
		if l.fileBuff != nil {
			l.fileBuff = nil
		}
		l.consoleBuff = nil
		_ = l.file.Close()
		l.wg.Done()
	}()
	for entry := range l.logChannel {
		// 写入文件
		l.writeFile(entry)
		// 异步写入控制台
		if l.consoleAsync {
			// 写入控制台
			l.consolePrint(entry)
		}
	}
}

// writeFile 写文件
func (l *ToolLogger) writeFile(entry *LogEntry) {
	l.checkFile(l.file)
	defer l.fileBuff.Reset()
	NotColor(l.fileBuff, l.Prefix, true)
	NotColor(l.fileBuff, entry.Timestamp.Format(l.TimeFormat), true)
	l.changeLevelNotColor(l.fileBuff, entry.Level)
	l.fileBuff.WriteString(entry.Message)
	l.fileBuff.WriteByte('\n')
	if _, err := l.file.Write(l.fileBuff.Bytes()); err != nil {
		panic(err.Error())
		//PrintLogger(Error, "写入文件失败: %s", err.Error())
	}
}

// consolePrint 控制台输出
func (l *ToolLogger) consolePrint(entry *LogEntry) {
	defer l.consoleBuff.Reset()
	Blue(l.consoleBuff, l.Prefix, true)
	White(l.consoleBuff, entry.Timestamp.Format(l.TimeFormat), true)
	l.changeLevel(l.consoleBuff, entry.Level)
	l.consoleBuff.WriteString(entry.Message)
	_, _ = fmt.Fprintln(l.consoleOut, l.consoleBuff.String())
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
	entry := NewEntry(level, namespace, fmt.Sprintf(format, args...))
	if !l.consoleAsync {
		l.consolePrint(entry)
	}
	// 写入通道异步写入文件
	l.logChannel <- entry
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
