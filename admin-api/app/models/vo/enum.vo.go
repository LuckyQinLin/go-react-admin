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
