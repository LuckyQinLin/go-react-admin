package request

// NoticeCreateRequest 通知创建
type NoticeCreateRequest struct {
	NoticeTitle   string `json:"noticeTitle"  binding:"required"`   // 公告标题
	NoticeType    int    `json:"noticeType"  binding:"required"`    // 公告类型(1通知 2公告)
	NoticeContent string `json:"noticeContent"  binding:"required"` // 公告内容
	Status        int    `json:"status"`                            // 部门状态
	UserName      string `json:"userName"`
}

// NoticeUpdateRequest 通知更新
type NoticeUpdateRequest struct {
	NoticeCreateRequest
	NoticeId int64 `json:"noticeId" binding:"required"` // 主键
}

// NoticePageRequest 通知分页查询
type NoticePageRequest struct {
	CommonPage
	NoticeTitle string `json:"noticeTitle"` // 公告标题
	UserName    string `json:"UserName"`    // 操作人
	NoticeType  int    `json:"configType"`  // 通知类型
}

// NoticeDeleteRequest 通知删除
type NoticeDeleteRequest struct {
	Ids      []int64 `json:"ids"`      // 删除通知的IDs
	UserName string  `json:"userName"` // 用户名
}
