package service

import (
	"admin-api/app/dao"
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/core"
	"gitee.com/molonglove/goboot/gorm"
	"time"
)

var Notice = new(NoticeService)

type NoticeService struct{}

// Create 通知创建
func (d *NoticeService) Create(param *request.NoticeCreateRequest) *response.BusinessError {
	var (
		dict      entity.Notice
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测岗位名称是否唯一
	condition = core.DB.Where("notice_title = ?", param.NoticeType)
	if exist, err = dao.Notice.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的通知名称")
		return response.CustomBusinessError(response.Failed, "存在相同的通知名称")
	}
	// 保存菜单数据
	now = time.Now()
	dict = entity.Notice{
		NoticeTitle:   param.NoticeTitle,
		NoticeContent: param.NoticeContent,
		NoticeType:    param.NoticeType,
		Status:        param.NoticeType,
		CreateBy:      param.UserName,
		CreateTime:    &now,
	}
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Notice.Create(tx, &dict)
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "创建通知失败")
	}
	core.Log.Info("创建通知[%d:%s]成功", dict.NoticeId, dict.NoticeTitle)
	return nil
}

// Update 通知更新
func (d *NoticeService) Update(param *request.NoticeUpdateRequest) *response.BusinessError {
	var (
		old       entity.Notice
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测岗位名称是否唯一
	condition = core.DB.Where("notice_title = ? and notice_id != ?", param.NoticeTitle, param.NoticeId)
	if exist, err = dao.Notice.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的通知名称")
		return response.CustomBusinessError(response.Failed, "存在相同的通知名称")
	}
	if old, err = dao.Notice.GetById(param.NoticeId); err != nil {
		core.Log.Error("当前通知不存在：%s", err.Error())
		return response.CustomBusinessError(response.Failed, "当前通知不存在")
	}
	// 保存部门数据
	now = time.Now()
	old.NoticeTitle = param.NoticeTitle
	old.NoticeContent = param.NoticeContent
	old.NoticeType = param.NoticeType
	old.Status = param.Status
	old.UpdateBy = param.UserName
	old.UpdateTime = &now
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Notice.UpdateById(tx, &old)
	}); err != nil {
		core.Log.Error("更新通知[%d]:%s", param.NoticeId, err.Error())
		return response.CustomBusinessError(response.Failed, "更新通知失败")
	}
	return nil
}

// Info 通知详情
func (d *NoticeService) Info(noticeId int64) (*response.NoticeInfoResponse, *response.BusinessError) {
	var (
		notice entity.Notice
		err    error
	)
	if notice, err = dao.Notice.GetById(noticeId); err != nil {
		return nil, response.CustomBusinessError(response.Failed, "当前通知不存在")
	}
	return &response.NoticeInfoResponse{
		NoticeId:      notice.NoticeId,
		NoticeTitle:   notice.NoticeTitle,
		NoticeContent: notice.NoticeContent,
		NoticeType:    notice.NoticeType,
		Status:        notice.Status,
	}, nil
}

// Page 通知分页
func (d *NoticeService) Page(param *request.NoticePageRequest) (*response.PageData, *response.BusinessError) {
	var (
		buildCondition = func(param *request.NoticePageRequest) func(db *gorm.DB) *gorm.DB {
			return func(db *gorm.DB) *gorm.DB {
				db.Model(&entity.Notice{})
				if param.NoticeType != 0 {
					db.Where("notice_type = ?", param.NoticeType)
				}
				if param.NoticeTitle != "" {
					db.Where("notice_title like concat('%', ?, '%')", param.NoticeTitle)
				}
				if param.UserName != "" {
					db.Where("create_by like concat('%', ?, '%')", param.UserName)
				}
				return db
			}
		}
		list  []response.NoticeTableResponse
		total int64
		err   error
	)
	if err = core.DB.Scopes(buildCondition(param)).Count(&total).Error; err != nil {
		core.Log.Error("统计通知数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取通知数据失败")
	}
	if err = core.DB.Scopes(buildCondition(param)).
		Find(&list).
		Error; err != nil {
		core.Log.Error("查询通知数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取通知数据失败")
	}
	return &response.PageData{
		Total: total,
		Page:  param.Page,
		Size:  param.Size,
		Data:  list,
	}, nil
}

// Delete 通知删除
// gitee.com/molonglove/goboot
// gitee.com/molonglove/goboot
func (d *NoticeService) Delete(param *request.NoticeDeleteRequest, username string) *response.BusinessError {
	var (
		err error
	)
	if err = core.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = dao.Notice.Delete(tx, param.Ids...); err != nil {
			core.Log.Error("删除通知失败：%s", err.Error())
			return
		}
		return
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "删除通知失败")
	}
	return nil
}
