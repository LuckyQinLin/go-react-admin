package response

type HttpCode int64

const (
	Success             HttpCode = 100200
	Failed              HttpCode = 100500
	LoginFailed         HttpCode = 100499
	DataNotExist        HttpCode = 100501
	UserPasswordError   HttpCode = 100502
	AuthNotExist        HttpCode = 100503
	AuthFail            HttpCode = 100504
	RequestParamError   HttpCode = 100505
	UserNameNotExist    HttpCode = 100506
	UserNameExist       HttpCode = 100507
	PhoneExist          HttpCode = 100508
	EmailExist          HttpCode = 100509
	TokenBuildError     HttpCode = 100510
	TokenTimeOut        HttpCode = 100511
	AddDataError        HttpCode = 100512
	SqlExecuteError     HttpCode = 100513
	DeleteSuccess       HttpCode = 100514
	ExistSameData       HttpCode = 100515
	SoftwareCreateError HttpCode = 100516
	DataNotNeedUpdate   HttpCode = 100517
	DataDeleteFail      HttpCode = 100518
	LogoutSuccess       HttpCode = 100519
	LogoutFail          HttpCode = 100520
	DateUpdateError     HttpCode = 100521
	CaptchaImageError   HttpCode = 100522
	UserNotAllowDelete  HttpCode = 100523
)

var Menus = map[HttpCode]string{
	Success:             "操作成功",
	Failed:              "操作失败",
	LoginFailed:         "登录失败",
	DataNotExist:        "数据不存在",
	UserPasswordError:   "账号密码不正确",
	AuthNotExist:        "认证信息不正确",
	AuthFail:            "校验认证信息失败",
	RequestParamError:   "请求参数错误",
	UserNameNotExist:    "用户名不存在",
	UserNameExist:       "账号已经存在",
	PhoneExist:          "用户手机号已经存在",
	EmailExist:          "用户邮箱已经注册",
	TokenBuildError:     "生成Token错误",
	TokenTimeOut:        "认证信息过期",
	AddDataError:        "增加数据失败",
	SqlExecuteError:     "SQL执行错误",
	DeleteSuccess:       "删除成功",
	ExistSameData:       "存在相同的数据",
	SoftwareCreateError: "软件创建失败",
	DataNotNeedUpdate:   "数据不需要修改",
	DataDeleteFail:      "数据删除失败",
	LogoutSuccess:       "退出登录成功",
	LogoutFail:          "退出登录失败",
	DateUpdateError:     "数据更新失败",
	CaptchaImageError:   "验证码生成失败",
	UserNotAllowDelete:  "当前用户不允许删除",
}

// Message 消息
type Message struct {
	Code HttpCode `json:"code"`
	Msg  string   `json:"message"`
	Data any      `json:"result"`
}

func Ok(data any) Message {
	return Message{
		Code: Success,
		Msg:  Menus[Success],
		Data: data,
	}
}

func Fail(data any) Message {
	return Message{
		Code: Failed,
		Msg:  Menus[Failed],
		Data: data,
	}
}

func ResultCustom(err *BusinessError) Message {
	return Message{
		Code: err.Code,
		Msg:  Menus[err.Code],
		Data: err.Error(),
	}
}

func Result(code HttpCode, data any) Message {
	return Message{
		Code: code,
		Msg:  Menus[code],
		Data: data,
	}
}

type PageData struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Data  any   `json:"records"`
}

func NewPageData(page, size int, total int64, data any) *PageData {
	return &PageData{
		Total: total,
		Page:  page,
		Size:  size,
		Data:  data,
	}
}
