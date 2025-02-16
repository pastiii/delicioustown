package utils

import (
	"DeliciousTown/global"
	"os"
)

// 判断目录是否存在,存在返回 true
func DirExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 创建目录
func CreateDir(path string) error {
	dirExist, err := DirExist(path)
	if err != nil {
		return err
	}
	if !dirExist {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			global.DefaultLogger.Sugar().Debugf("创建[%s]目录失败: %s", path, err)
		}
	}
	return err
}