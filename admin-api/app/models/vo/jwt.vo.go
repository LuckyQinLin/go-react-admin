package vo

import "github.com/golang-jwt/jwt/v5"

// UserClaims 自定义jwt数据体
type UserClaims struct {
	UserId   int64
	Username string
	Email    string
	jwt.RegisteredClaims
}
