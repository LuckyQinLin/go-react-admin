package controller

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/app/service"
	"admin-api/internal/gin"
	"admin-api/utils"
	"github.com/goccy/go-json"
	"net/http"
	"runtime"
	"time"
)

type BaseController struct{}

func (b *BaseController) FunctionName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// Parse 解析前置数据
func (b *BaseController) Parse(c *gin.Context, title string, businessType vo.BusinessType, param any) (*vo.UserClaims, *entity.Operate) {
	var (
		value  any
		ip     string
		claims *vo.UserClaims
		bytes  []byte
		now    time.Time
	)
	value, _ = c.Get(vo.ClaimsInfo)
	claims = value.(*vo.UserClaims)
	ip = c.ClientIP()

	bytes, _ = json.Marshal(param)
	now = time.Now()
	return claims, &entity.Operate{
		Title:         title,
		BusinessType:  int(businessType),
		Method:        b.FunctionName(),
		RequestMethod: c.Request.Method,
		OperatorType:  1,
		OperName:      claims.Username,
		DeptName:      claims.DeptName,
		OperUrl:       c.Request.URL.Path,
		OperIp:        ip,
		OperParam:     string(bytes),
		OperLocation:  utils.IpAddress(ip),
		OperTime:      &now,
	}
}

func (b *BaseController) result(ctx *gin.Context, oper *entity.Operate, isSuccess bool, msg response.Message) {
	bytes, _ := json.Marshal(msg)
	if isSuccess {
		oper.Status = 1
		oper.JsonResult = string(bytes)
	} else {
		oper.Status = 0
		oper.ErrorMsg = string(bytes)
	}
	oper.CostTime = time.Now().Sub(*oper.OperTime).Microseconds()
	service.Operate.Push(oper)
	ctx.JSON(http.StatusOK, msg)
	return
}

func (b *BaseController) Failed(ctx *gin.Context, oper *entity.Operate, msg response.Message) {
	b.result(ctx, oper, false, msg)
}

func (b *BaseController) Success(ctx *gin.Context, oper *entity.Operate, msg response.Message) {
	b.result(ctx, oper, true, msg)
}
