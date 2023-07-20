package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
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

// BcryptEncode Bcrypt加密密码
func BcryptEncode(cleartext string) (ciphertext string, err error) {
	target, err := bcrypt.GenerateFromPassword([]byte(cleartext), bcrypt.MinCost)
	if err != nil {
		return
	}
	ciphertext = string(target)
	return
}

// BcryptVerify 验证密码
func BcryptVerify(ciphertext string, cleartext string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ciphertext), []byte(cleartext))
	if err != nil {
		return false
	}
	return true
}
