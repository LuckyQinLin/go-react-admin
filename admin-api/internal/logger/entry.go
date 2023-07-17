package logger

import "time"

type LogEntry struct {
	Level     LogLevel
	Message   string
	Namespace string
	Timestamp time.Time
}

func NewEntry(level LogLevel, namespace string, msg string) *LogEntry {
	return &LogEntry{
		Level:     level,
		Message:   msg,
		Namespace: namespace,
		Timestamp: time.Now(),
	}
}
