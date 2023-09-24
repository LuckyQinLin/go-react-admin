package controller

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/app/service"
	"gitee.com/molonglove/goboot/gin"
	"net/http"
)

var Notice = new(NoticeController)

type NoticeController struct{ BaseController }

// NoticeCreate 通知创建
func (d *NoticeController) NoticeCreate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.NoticeCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = d.Parse(ctx, "通知创建", vo.Add, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Notice.Create(&param); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("通知创建成功"))
}

// NoticeUpdate 通知更新
func (d *NoticeController) NoticeUpdate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.NoticeUpdateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = d.Parse(ctx, "通知修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Notice.Update(&param); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("通知修改成功"))
}

// NoticeDelete 通知删除
func (d *NoticeController) NoticeDelete(ctx *gin.Context) {
	var (
		err       error
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.NoticeDeleteRequest
		customErr *response.BusinessError
	)
	claims, operate = d.Parse(ctx, "通知删除", vo.Delete, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	operate.ParamToJson(param)
	if customErr = service.Notice.Delete(&param, claims.Username); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("删除通知数据成功"))
}

// NoticePage 通知分页
func (d *NoticeController) NoticePage(ctx *gin.Context) {
	var (
		param     request.NoticePageRequest
		result    *response.PageData
		customErr *response.BusinessError
		err       error
	)
	if err = ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Notice.Page(&param); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// NoticeInfo 通知详情
func (d *NoticeController) NoticeInfo(ctx *gin.Context) {
	var (
		err       error
		dictId    int64
		info      *response.NoticeInfoResponse
		customErr *response.BusinessError
	)
	if dictId, err = ctx.QueryInt64("dictId"); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if info, customErr = service.Notice.Info(dictId); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(info))
}
