package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

// ConfigInfoResponse 字典详情
type ConfigInfoResponse struct {
	ConfigId    int64  `json:"dictId"`      // 主键
	ConfigName  string `json:"configName"`  // 参数名称
	ConfigKey   string `json:"configKey"`   // 参数键名
	ConfigValue string `json:"configValue"` // 参数键值
	ConfigType  int    `json:"configType"`  // 参数类型
	Remark      string `json:"remark"`      // 备注
}

// ConfigTableResponse 字典表属性
type ConfigTableResponse struct {
	ConfigId    int64      `json:"dictId"`      // 主键
	ConfigName  string     `json:"configName"`  // 参数名称
	ConfigKey   string     `json:"configKey"`   // 参数键名
	ConfigValue string     `json:"configValue"` // 参数键值
	ConfigType  int        `json:"configType"`  // 参数类型
	Remark      string     `json:"remark"`      // 备注
	CreateTime  *time.Time `json:"createTime"`  // 创建时间
}

func (s ConfigTableResponse) MarshalJSON() ([]byte, error) {
	type temp ConfigTableResponse
	return json.Marshal(&struct {
		temp
		CreateTime utils.DateTime `json:"createTime"`
	}{
		temp:       (temp)(s),
		CreateTime: utils.DateTime(*s.CreateTime),
	})
}
