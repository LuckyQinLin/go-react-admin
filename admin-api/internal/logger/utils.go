package logger

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// buildPath 构建路径
func buildPath(logFile string) string {
	return path.Join(
		logFile,
		time.Now().Format("2006-01-02"),
		"out.log",
	)
}

func openFile(path string) (*os.File, error) {
	if strings.HasPrefix(path, "~") {
		homePath, _ := os.UserHomeDir()
		path = strings.ReplaceAll(path, "~", homePath)
	}
	dir := filepath.Dir(path)
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		// 创建目录（如果不存在）
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			PrintLogger(Error, "创建目录失败: %s", err.Error())
			return nil, err
		}
	}
	// 创建文件
	PrintLogger(Info, "创建文件: %s", path)
	return os.Create(path)
}
