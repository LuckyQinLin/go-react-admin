package request

// PostPageRequest 岗位分页查询
type PostPageRequest struct {
	CommonPage
	PostCode string `json:"postCode"` // 岗位编码
	PostName string `json:"postName"` // 岗位名称
	Status   int    `json:"status"`   // 状态
}

// PostCreateRequest 岗位创建
type PostCreateRequest struct {
	PostCode string `json:"postCode"` // 岗位编码
	PostName string `json:"postName"` // 岗位名称
	PostSort int    `json:"postSort"` // 岗位顺序
	Status   int    `json:"status"`   // 状态
	Remark   string `json:"remark"`   // 备注
	UserName string `json:"userName"`
}

// PostUpdateRequest 岗位修改
type PostUpdateRequest struct {
	PostCreateRequest
	PostId int64 `json:"postId"` // 岗位修改
}

// PostDeleteRequest 岗位删除
type PostDeleteRequest struct {
	Ids      []int64 `json:"ids" binding:"required"`
	UserName string  `json:"userName"`
}
