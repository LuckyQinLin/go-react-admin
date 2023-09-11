package model

type LogLevel int

const (
	Info LogLevel = iota
	Debug
	Warn
	Error
	Fail
)
