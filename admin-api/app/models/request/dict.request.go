package request

// DictCreateRequest 字典创建
type DictCreateRequest struct {
	DictName string `json:"dictName" binding:"required"` // 字典名称
	DictType string `json:"dictType" binding:"required"` // 字典类型
	Status   int    `json:"status"`                      // 字典状态
	Remark   string `json:"remark"`                      // 备注
	UserName string `json:"userName"`
}

// DictUpdateRequest 字典更新
type DictUpdateRequest struct {
	DictCreateRequest
	DictId int64 `json:"dictId" binding:"required"`
}

// DictDeleteRequest 字典删除
type DictDeleteRequest struct {
	Ids      []int64 `json:"ids"`      // 删除字典的IDs
	UserName string  `json:"userName"` // 用户名
}

// DictPageRequest 字典分页查询
type DictPageRequest struct {
	CommonPage
	DictName string `json:"dictName"` // 字典名称
	DictType string `json:"dictType"` // 字典类型
	Status   int    `json:"status"`   // 字典状态
}
