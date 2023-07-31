package request

// ConfigCreateRequest 参数创建
type ConfigCreateRequest struct {
	ConfigName  string `json:"configName" binding:"required"`  // 参数名称
	ConfigKey   string `json:"configKey" binding:"required"`   // 参数键名
	ConfigValue string `json:"configValue" binding:"required"` // 参数键值
	ConfigType  int    `json:"configType"`                     // 参数类型
	Remark      string `json:"remark"`                         // 备注
	UserName    string `json:"userName"`
}

// ConfigUpdateRequest 参数更新
type ConfigUpdateRequest struct {
	ConfigCreateRequest
	ConfigId int64 `json:"configId" binding:"required"` // 主键
}

// ConfigPageRequest 参数分页查询
type ConfigPageRequest struct {
	CommonPage
	ConfigName string `json:"configName"` // 参数名称
	ConfigKey  string `json:"configKey"`  // 参数键名
	ConfigType int    `json:"configType"` // 参数类型
}

// ConfigDeleteRequest 参数删除
type ConfigDeleteRequest struct {
	Ids      []int64 `json:"ids"`      // 删除参数的IDs
	UserName string  `json:"userName"` // 用户名
}
