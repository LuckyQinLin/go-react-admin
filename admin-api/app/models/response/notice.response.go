package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

// NoticeInfoResponse 通知详情
type NoticeInfoResponse struct {
	NoticeId      int64  `json:"noticeId"`      // 主键
	NoticeTitle   string `json:"noticeTitle"`   // 公告标题
	NoticeType    int    `json:"noticeType"`    // 公告类型(1通知 2公告)
	NoticeContent string `json:"noticeContent"` // 公告内容
	Status        int    `json:"status"`        // 部门状态
}

// NoticeTableResponse 通知表属性
type NoticeTableResponse struct {
	NoticeId    int64      `json:"noticeId"`    // 主键
	NoticeTitle string     `json:"noticeTitle"` // 公告标题
	NoticeType  int        `json:"noticeType"`  // 公告类型(1通知 2公告)
	Status      int        `json:"status"`      // 部门状态
	CreateBy    string     `json:"createBy"`    // 创建者
	CreateTime  *time.Time `json:"createTime"`  // 创建时间
}

func (s NoticeTableResponse) MarshalJSON() ([]byte, error) {
	type temp NoticeTableResponse
	return json.Marshal(&struct {
		temp
		CreateTime utils.DateTime `json:"createTime"`
	}{
		temp:       (temp)(s),
		CreateTime: utils.DateTime(*s.CreateTime),
	})
}
