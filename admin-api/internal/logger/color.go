package logger

import "bytes"

type colorLevel int

const (
	colorGreen colorLevel = iota
	colorWhite
	colorYellow
	colorRed
	colorBlue
	colorMagenta
	colorCyan
	colorReset
	colorNot
)

var (
	green   = []byte("\033[32m")
	white   = []byte("\033[37m")
	yellow  = []byte("\033[33m")
	red     = []byte("\033[31m")
	blue    = []byte("\033[34m")
	magenta = []byte("\033[35m")
	cyan    = []byte("\033[36m")
	reset   = []byte("\033[0m")
)

func color(buffer *bytes.Buffer, level colorLevel, content string, closeStr bool) {
	if closeStr {
		buffer.WriteString("[")
	}
	switch level {
	case colorGreen:
		buffer.Write(green)
	case colorWhite:
		buffer.Write(white)
	case colorYellow:
		buffer.Write(yellow)
	case colorRed:
		buffer.Write(red)
	case colorBlue:
		buffer.Write(blue)
	case colorMagenta:
		buffer.Write(magenta)
	case colorCyan:
		buffer.Write(cyan)
	default:
		break
	}
	buffer.WriteString(content)
	if level != colorNot {
		buffer.Write(reset)
	}
	if closeStr {
		buffer.WriteString("] ")
	}
}

func NotColor(buffer *bytes.Buffer, content string, closeStr bool) {
	color(buffer, colorNot, content, closeStr)
}

func Green(buffer *bytes.Buffer, content string, closeStr bool) {
	color(buffer, colorGreen, content, closeStr)
}

func White(buffer *bytes.Buffer, content string, closeStr bool) {
	color(buffer, colorWhite, content, closeStr)
}

func Yellow(buffer *bytes.Buffer, content string, closeStr bool) {
	color(buffer, colorYellow, content, closeStr)
}

func Red(buffer *bytes.Buffer, content string, closeStr bool) {
	color(buffer, colorRed, content, closeStr)
}

func Blue(buffer *bytes.Buffer, content string, closeStr bool) {
	color(buffer, colorBlue, content, closeStr)
}

func Magenta(buffer *bytes.Buffer, content string, closeStr bool) {
	color(buffer, colorMagenta, content, closeStr)
}

func Cyan(buffer *bytes.Buffer, content string, closeStr bool) {
	color(buffer, colorCyan, content, closeStr)
}
