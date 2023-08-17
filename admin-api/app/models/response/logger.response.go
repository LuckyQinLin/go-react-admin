package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

// VisitLogResponse 访问日志
type VisitLogResponse struct {
	Id        int64      `json:"id"`        // 主键
	UserName  string     `json:"userName"`  // 用户名称
	Ip        string     `json:"ip"`        // 登录地址
	Address   string     `json:"address"`   // 登录地点
	Browser   string     `json:"browser"`   // 浏览器
	Os        string     `json:"os"`        // 操作系统
	Status    int        `json:"status"`    // 登录结果
	Msg       string     `json:"msg"`       // 登录信息
	LoginTime *time.Time `json:"loginTime"` // 登录时间
}

func (s VisitLogResponse) MarshalJSON() ([]byte, error) {
	type temp VisitLogResponse
	return json.Marshal(&struct {
		temp
		LoginTime utils.DateTime `json:"loginTime"`
	}{
		temp:      (temp)(s),
		LoginTime: utils.DateTime(*s.LoginTime),
	})
}

// OperateLogResponse 操作日志
type OperateLogResponse struct {
	Id           int64      `json:"id"`           // 主键
	Title        string     `json:"title"`        // 模块名称
	BusinessType int        `json:"businessType"` // 操作类型名称
	OperatorType int        `json:"operatorType"` // 操作类型名称
	Ip           string     `json:"ip"`           // IP地址
	Address      string     `json:"address"`      // 操作地址
	Status       int        `json:"status"`       // 操作结果
	OperTime     *time.Time `json:"operTime"`     // 操作时间
	CostTime     int64      `json:"costTime"`     // 耗时
}

func (s OperateLogResponse) MarshalJSON() ([]byte, error) {
	type temp OperateLogResponse
	return json.Marshal(&struct {
		temp
		OperTime utils.DateTime `json:"operTime"`
	}{
		temp:     (temp)(s),
		OperTime: utils.DateTime(*s.OperTime),
	})
}
