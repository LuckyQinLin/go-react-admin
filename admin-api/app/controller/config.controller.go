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

var Config = new(ConfigController)

type ConfigController struct{ BaseController }

// ConfigCreate 参数创建
func (d *ConfigController) ConfigCreate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.ConfigCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = d.Parse(ctx, "参数创建", vo.Add, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Config.Create(&param); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("参数创建成功"))
}

// ConfigUpdate 参数更新
func (d *ConfigController) ConfigUpdate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.ConfigUpdateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = d.Parse(ctx, "参数修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Config.Update(&param); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("参数修改成功"))
}

// ConfigDelete 参数删除
func (d *ConfigController) ConfigDelete(ctx *gin.Context) {
	var (
		err       error
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.ConfigDeleteRequest
		customErr *response.BusinessError
	)
	claims, operate = d.Parse(ctx, "参数删除", vo.Delete, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	operate.ParamToJson(param)
	if customErr = service.Config.Delete(&param, claims.Username); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("删除参数数据成功"))
}

// ConfigPage 参数分页
func (d *ConfigController) ConfigPage(ctx *gin.Context) {
	var (
		param     request.ConfigPageRequest
		result    *response.PageData
		customErr *response.BusinessError
		err       error
	)
	if err = ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Config.Page(&param); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// ConfigInfo 参数详情
func (d *ConfigController) ConfigInfo(ctx *gin.Context) {
	var (
		err       error
		dictId    int64
		info      *response.ConfigInfoResponse
		customErr *response.BusinessError
	)
	if dictId, err = ctx.QueryInt64("dictId"); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if info, customErr = service.Config.Info(dictId); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(info))
}
