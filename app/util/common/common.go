package common

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"
)

func GetPassword(password string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, password)
	return hex.EncodeToString(MD5.Sum(nil))
}

func GetToken(account string, password string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	str := account + password + t
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, str)
	return hex.EncodeToString(MD5.Sum(nil))
}
