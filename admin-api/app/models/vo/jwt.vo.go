package vo

// UserClaims 自定义jwt数据体
type UserClaims struct {
	UserId   int64
	DeptId   int64
	DeptName string
	Username string
	Email    string
	Phone    string
}
