package utils

import "time"

// GetUnixMilli 获取当前时间戳
func GetUnixMilli() int64 {
	return time.Now().UnixMilli()
}

func GetUnix() int64 {
	return time.Now().Unix()
}
