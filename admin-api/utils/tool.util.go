package utils

import "time"

// IsEmptyString 字符串是否为空
func IsEmptyString(target string) bool {
	return target == ""
}

// TimeDifferenceToString 时间差值转文字
func TimeDifferenceToString(target time.Time) string {
	return ""
}

// Include 包含
func Include[T ~string | ~int | ~int8 | ~int16 | ~int32 | ~int64](list []T, data T) bool {
	for _, item := range list {
		if item == data {
			return true
		}
	}
	return false
}
