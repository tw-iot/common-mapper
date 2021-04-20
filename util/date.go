package util

import "time"

//毫秒 时间戳
func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

//秒 时间戳
func GetSecondTimestamp() int64 {
	return time.Now().Unix()
}
