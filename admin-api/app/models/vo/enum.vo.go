package vo

const (
	ClaimsInfo    = "Claims"
	AuthHeader    = "Authorization"
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
