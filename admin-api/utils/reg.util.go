package utils

import "regexp"

// IsEnglishChineseNumericOnly 仅支持英文、中文和数字（不包含空格）
func IsEnglishChineseNumericOnly(str string) bool {
	reg := regexp.MustCompile("^[a-zA-Z\\p{Han}\\d]+$")
	return reg.MatchString(str)
}
