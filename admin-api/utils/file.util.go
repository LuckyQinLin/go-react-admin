package utils

import (
	"os"
)

// CreateDir 创建目录
func CreateDir(path string) error {
	var err error
	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
		}
	}
	return err
}

// IsFileExist 判断文件是否存在
func IsFileExist(file string) error {
	var (
		err error
	)
	if _, err = os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(file)
		}
	}
	return err
}
