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

var Dict = new(DictController)

type DictController struct {
	BaseController
}

// DictCreate 字典创建
func (d *DictController) DictCreate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.DictCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = d.Parse(ctx, "字典创建", vo.Add, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Dict.Create(&param); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("字典创建成功"))
}

// DictUpdate 字典更新
func (d *DictController) DictUpdate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.DictUpdateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = d.Parse(ctx, "字典修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Dict.Update(&param); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("字典修改成功"))
}

// DictDelete 字典删除
func (d *DictController) DictDelete(ctx *gin.Context) {
	var (
		err       error
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.DictDeleteRequest
		customErr *response.BusinessError
	)
	claims, operate = d.Parse(ctx, "字典删除", vo.Delete, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	operate.ParamToJson(param)
	if customErr = service.Dict.Delete(&param, claims.Username); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("删除字典数据成功"))
}

// DictPage 字典分页
func (d *DictController) DictPage(ctx *gin.Context) {
	var (
		param     request.DictPageRequest
		result    *response.PageData
		customErr *response.BusinessError
		err       error
	)
	if err = ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Dict.Page(&param); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// DictInfo 字典详情
func (d *DictController) DictInfo(ctx *gin.Context) {
	var (
		err       error
		dictId    int64
		info      *response.DictInfoResponse
		customErr *response.BusinessError
	)
	if dictId, err = ctx.QueryInt64("dictId"); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if info, customErr = service.Dict.Info(dictId); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(info))
}
