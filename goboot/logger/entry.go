package logger

import (
	"goboot/domain/model"
	"time"
)

type LogEntry struct {
	Level     model.LogLevel
	Message   string
	Namespace string
	Timestamp time.Time
}

func NewEntry(level model.LogLevel, namespace string, msg string) *LogEntry {
	return &LogEntry{
		Level:     level,
		Message:   msg,
		Namespace: namespace,
		Timestamp: time.Now(),
	}
}
