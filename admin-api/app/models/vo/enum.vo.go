package vo

const (
	ClaimsInfo    = "Claims"
	AuthHeader    = "Authorization"
	NewToken      = "NewToken"
	HeaderUserKey = "userId"
	WSHeaderKey   = "WsValid"
)

type BusinessType int

const (
	Other BusinessType = iota
	Add
	Update
	Delete
)

const (
	SUPER_USER_ID int64 = 1 // 超级用户ID
	SUPER_ROLE_ID int64 = 1 // 超级角色ID
)
