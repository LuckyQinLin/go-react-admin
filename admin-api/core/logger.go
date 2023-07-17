package core

import (
	"admin-api/internal/logger"
	"github.com/mitchellh/go-homedir"
)

var Log *logger.ToolLogger

func InitLogger() {
	logPath, _ := homedir.Expand("~/.admin/logs")
	Log = logger.DefaultLogger(logPath)
}
