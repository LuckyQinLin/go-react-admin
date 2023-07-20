package response

type BusinessError struct {
	Code HttpCode
	Msg  string
}

func (b BusinessError) Error() string {
	return b.Msg
}

func NewBusinessError(code HttpCode) *BusinessError {
	return &BusinessError{
		Code: code,
		Msg:  Menus[code],
	}
}

func CustomBusinessError(code HttpCode, msg string) *BusinessError {
	return &BusinessError{
		Code: code,
		Msg:  msg,
	}
}

func LoginBusinessError(msg string) *BusinessError {
	return &BusinessError{
		Code: LoginFailed,
		Msg:  msg,
	}
}
