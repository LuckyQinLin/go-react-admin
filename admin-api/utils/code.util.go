package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*=+"

// TransformMd5 将字符转换为md5
func TransformMd5(target string) string {
	d := []byte(target)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func RandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
