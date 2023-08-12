package utils

import (
	"crypto/md5"
	"encoding/base64"
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

// GenerateRandomToken 生成随机token串
func GenerateRandomToken(length int) (string, error) {
	// 计算需要生成的字节数
	byteLength := (length * 3) / 4

	// 生成随机字节数组
	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// 将随机字节数组转换为 base64 编码的字符串
	token := base64.URLEncoding.EncodeToString(randomBytes)[:length]
	return token, nil
}
