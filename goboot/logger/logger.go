package logger

import (
	"bytes"
	"fmt"
	"goboot/domain/model"
	"io"
	"os"
	"sync"
	"time"
)

type ToolLogger struct {
	Prefix       string         // 日志前缀
	TimeFormat   string         // 时间格式
	filePath     string         // 日志文件路径
	logLevel     model.LogLevel // 默认日志等级
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

var Default = DefaultLogger("")

func DefaultLogger(logFile string) *ToolLogger {
	var (
		logger *ToolLogger
		file   *os.File
		err    error
	)
	logger = &ToolLogger{
		filePath:     logFile,
		logLevel:     model.Info,
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
		PrintLogger(model.Error, "创建文件失败: %s", err.Error())
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
		PrintLogger(model.Error, "无法获取文件信息:%s", err.Error())
		return
	}
	if info.Size() >= l.maxFileSize {
		PrintLogger(model.Warn, "日志文件大小超过阈值")
		if err := file.Close(); err != nil {
			PrintLogger(model.Error, "关闭当前日志文件时出错:%s", err.Error())
			return
		}
		filePrefix = fmt.Sprintf("%s_%s_", l.Prefix, time.Now().Format("20060102"))
		if count, err = CountFileNum(l.filePath, filePrefix); err != nil {
			PrintLogger(model.Error, "统计文件数量时出错: %s", err.Error())
			return
		}

		fileNewName = buildPathWithName(l.filePath, fmt.Sprintf("%s%d.log", filePrefix, count+1))
		logPath = buildPath(l.filePath)
		if err = os.Rename(logPath, fileNewName); err != nil {
			PrintLogger(model.Error, "重命名日志文件时出错: %s", err.Error())
			return
		}
		if file, err = openFile(logPath); err != nil {
			PrintLogger(model.Error, "创建新的日志文件时出错: %s", err.Error())
			return
		}
		l.file = file
	}
}

func (l *ToolLogger) GetPrefix() string {
	return l.Prefix
}

func (l *ToolLogger) PrintLogger(level model.LogLevel, format string, args ...any) {
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

func (l *ToolLogger) changeLevel(buffer *bytes.Buffer, level model.LogLevel) {
	switch level {
	case model.Info:
		Blue(buffer, "Info", true)
	case model.Debug:
		Green(buffer, "Debug", true)
	case model.Warn:
		Yellow(buffer, "Warn", true)
	case model.Fail:
		Magenta(buffer, "Fail", true)
	case model.Error:
		Red(buffer, "Error", true)
	default:
		White(buffer, "Unknown", true)
	}
}

func (l *ToolLogger) changeLevelNotColor(buffer *bytes.Buffer, level model.LogLevel) {
	switch level {
	case model.Info:
		buffer.WriteString("[Info] ")
	case model.Debug:
		buffer.WriteString("[Debug] ")
	case model.Warn:
		buffer.WriteString("[Warn] ")
	case model.Fail:
		buffer.WriteString("[Fail] ")
	case model.Error:
		buffer.WriteString("[Error] ")
	default:
		buffer.WriteString("[Unknown] ")
	}
}

func (l *ToolLogger) log(level model.LogLevel, namespace string, format string, args ...any) {
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
	l.log(model.Info, "", format, args...)
}

func (l *ToolLogger) Debug(format string, args ...any) {
	l.log(model.Debug, "", format, args...)
}

func (l *ToolLogger) Warn(format string, args ...any) {
	l.log(model.Warn, "", format, args...)
}

func (l *ToolLogger) Fail(format string, args ...any) {
	l.log(model.Fail, "", format, args...)
}

func (l *ToolLogger) Error(format string, args ...any) {
	l.log(model.Error, "", format, args...)
}
