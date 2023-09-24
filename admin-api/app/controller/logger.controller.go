package controller

import (
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/service"
	"gitee.com/molonglove/goboot/gin"
	"net/http"
)

var Logger = new(LoggerController)

// LoggerController 日志
type LoggerController struct{}

// VisitPage 访问日志
func (l *LoggerController) VisitPage(ctx *gin.Context) {
	var (
		param     request.VisitLogRequest
		result    *response.PageData
		customErr *response.BusinessError
		err       error
	)
	if err = ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Logger.VisitPage(&param); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// OperatePage 操作日志查询
func (l *LoggerController) OperatePage(ctx *gin.Context) {
	var (
		param     request.OperateLogRequest
		result    *response.PageData
		customErr *response.BusinessError
		err       error
	)
	if err = ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Logger.OperatePage(&param); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}
