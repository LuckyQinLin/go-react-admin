package service

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/core"
	"admin-api/internal/gorm"
)

var Logger = new(LoggerService)

type LoggerService struct{}

// VisitPage 访问日志分页查询
func (s LoggerService) VisitPage(param *request.VisitLogRequest) (*response.PageData, *response.BusinessError) {
	var (
		buildCondition = func(param *request.VisitLogRequest) func(db *gorm.DB) *gorm.DB {
			return func(db *gorm.DB) *gorm.DB {
				db.Model(&entity.Visit{})
				if param.Status != 0 {
					db.Where("status = ?", param.Status)
				}
				if param.UserName != "" {
					db.Where("user_name like concat('%', ?, '%')", param.UserName)
				}
				if param.Address != "" {
					db.Where("ip_addr like concat('%', ?, '%')", param.Address)
				}
				if param.StartTime == nil && param.EndTime != nil {
					db.Where("login_time <= ?", param.EndTime)
				}
				if param.StartTime != nil && param.EndTime == nil {
					db.Where("login_time >= ?", param.StartTime)
				}
				if param.StartTime != nil && param.EndTime != nil {
					db.Where("login_time between ? and ?", param.StartTime, param.EndTime)
				}
				return db
			}
		}
		list  []response.VisitLogResponse
		total int64
		err   error
	)
	if err = core.DB.Scopes(buildCondition(param)).Count(&total).Error; err != nil {
		core.Log.Error("统计访问数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取访问数据失败")
	}
	if err = core.DB.Scopes(buildCondition(param)).
		Select("visit_id as id, user_name, ip_addr as ip, login_location as address, browser, os, status, msg, login_time").
		Find(&list).
		Error; err != nil {
		core.Log.Error("查询访问数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取访问数据失败")
	}
	return &response.PageData{
		Total: total,
		Page:  param.Page,
		Size:  param.Size,
		Data:  list,
	}, nil
}

// OperatePage 操作日志分页查询
func (s LoggerService) OperatePage(param *request.OperateLogRequest) (*response.PageData, *response.BusinessError) {
	var (
		buildCondition = func(param *request.OperateLogRequest) func(db *gorm.DB) *gorm.DB {
			return func(db *gorm.DB) *gorm.DB {
				db.Model(&entity.Operate{})
				if param.Status != 0 {
					db.Where("status = ?", param.Status)
				}
				if param.UserName != "" {
					db.Where("oper_name like concat('%', ?, '%')", param.UserName)
				}
				if param.OperateType != -1 {
					db.Where("operator_type = ?", param.OperateType)
				}
				if param.BusinessType != -1 {
					db.Where("business_type = ?", param.BusinessType)
				}
				if param.StartTime == nil && param.EndTime != nil {
					db.Where("oper_time <= ?", param.EndTime)
				}
				if param.StartTime != nil && param.EndTime == nil {
					db.Where("oper_time >= ?", param.StartTime)
				}
				if param.StartTime != nil && param.EndTime != nil {
					db.Where("oper_time between ? and ?", param.StartTime, param.EndTime)
				}
				return db
			}
		}
		list  []response.OperateLogResponse
		total int64
		err   error
	)
	if err = core.DB.Scopes(buildCondition(param)).Count(&total).Error; err != nil {
		core.Log.Error("统计操作数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取操作数据失败")
	}
	if err = core.DB.Scopes(buildCondition(param)).
		Select("oper_id as id, title, business_type, operator_type, oper_ip as ip, oper_location as address, status, oper_time, cost_time").
		Order("oper_time").
		Find(&list).
		Error; err != nil {
		core.Log.Error("查询操作数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取操作数据失败")
	}
	return &response.PageData{
		Total: total,
		Page:  param.Page,
		Size:  param.Size,
		Data:  list,
	}, nil
}
