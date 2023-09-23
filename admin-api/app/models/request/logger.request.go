package request

import "time"

// VisitLogRequest 访问日志
type VisitLogRequest struct {
	CommonPage
	UserName  string `json:"userName"`  // 用户名称
	Status    int    `json:"status"`    //访问结果
	StartTime string `json:"startTime"` // 开始时间
	EndTime   string `json:"endTime"`   // 结束时间
	Address   string `json:"address"`   // IP地址
}

// OperateLogRequest 操作日志
type OperateLogRequest struct {
	CommonPage
	UserName     string     `json:"userName"`     // 用户名称
	Status       int        `json:"status"`       // 访问结果
	StartTime    *time.Time `json:"startTime"`    // 开始时间
	EndTime      *time.Time `json:"endTime"`      // 结束时间
	OperateType  int        `json:"operateType"`  // 操作类型
	BusinessType int        `json:"businessType"` // 业务类型
}
