package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

// PostPageResponse 岗位分页返回
type PostPageResponse struct {
	PostId     int64      `json:"postId"`     // 岗位ID
	PostCode   string     `json:"postCode"`   // 岗位编码
	PostName   string     `json:"postName"`   // 岗位名称
	PostSort   int        `json:"postSort"`   // 岗位顺序
	Status     int        `json:"status"`     // 状态
	CreateTime *time.Time `json:"createTime"` // 创建时间
}

func (s PostPageResponse) MarshalJSON() ([]byte, error) {
	type temp PostPageResponse
	return json.Marshal(&struct {
		temp
		CreateTime utils.DateTime `json:"createTime"`
	}{
		temp:       (temp)(s),
		CreateTime: utils.DateTime(*s.CreateTime),
	})
}

// PostInfoResponse 岗位详情
type PostInfoResponse struct {
	PostId   int64  `json:"postId"`   // 岗位ID
	PostCode string `json:"postCode"` // 岗位编码
	PostName string `json:"postName"` // 岗位名称
	PostSort int    `json:"postSort"` // 岗位顺序
	Status   int    `json:"status"`   // 状态
	Remark   string `json:"remark"`   // 备注
}

// PostListResponse 岗位信息
type PostListResponse struct {
	PostId   int64  `json:"value"`
	PostName string `json:"label"`
}
