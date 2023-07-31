package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

// DictInfoResponse 字典详情
type DictInfoResponse struct {
	DictId   int64  `json:"dictId"`   // 主键
	DictName string `json:"dictName"` // 字典名称
	DictType string `json:"dictType"` // 字典类型
	Status   int    `json:"status"`   // 状态
	Remark   string `json:"remark"`   // 备注
}

// DictTableResponse 字典表属性
type DictTableResponse struct {
	DictId     int64      `json:"dictId"`     // 主键
	DictName   string     `json:"dictName"`   // 字典名称
	DictType   string     `json:"dictType"`   // 字典类型
	Status     int        `json:"status"`     // 状态
	Remark     string     `json:"remark"`     // 备注
	CreateTime *time.Time `json:"createTime"` // 创建时间
}

func (s DictTableResponse) MarshalJSON() ([]byte, error) {
	type temp DictTableResponse
	return json.Marshal(&struct {
		temp
		CreateTime utils.DateTime `json:"createTime"`
	}{
		temp:       (temp)(s),
		CreateTime: utils.DateTime(*s.CreateTime),
	})
}
