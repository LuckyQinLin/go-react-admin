package gin

// GinLogger 日志
type GinLogger interface {
	Println(format string, values ...any)
}
