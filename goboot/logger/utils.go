package logger

import (
	"goboot/domain/model"
	"io/fs"
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

func buildPathWithName(basePath string, name string) string {
	return path.Join(
		basePath,
		time.Now().Format("2006-01-02"),
		name,
	)
}

func CountFileNum(path string, prefix string) (int, error) {
	if strings.HasPrefix(path, "~") {
		homePath, _ := os.UserHomeDir()
		path = strings.ReplaceAll(path, "~", homePath)
	}
	count := 0
	if err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasPrefix(info.Name(), prefix) {
			count++
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return count, nil
}

func openFile(path string) (*os.File, error) {
	if strings.HasPrefix(path, "~") {
		homePath, _ := os.UserHomeDir()
		path = strings.ReplaceAll(path, "~", homePath)
	}
	dir := filepath.Dir(path)
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
			// 创建目录（如果不存在）
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				PrintLogger(model.Error, "创建目录失败: %s", err.Error())
				return nil, err
			}
		}
		// 创建文件
		PrintLogger(model.Info, "创建文件: %s", path)
		return os.Create(path)
	}
	return os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
}
