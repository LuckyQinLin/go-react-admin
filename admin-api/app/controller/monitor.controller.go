package controller

import (
	"admin-api/app/models/response"
	"admin-api/app/service"
	"gitee.com/molonglove/goboot/gin"
	"net/http"
)

// https://github.com/shirou/gopsutil
var Monitor = new(MonitorController)

// MonitorController 监控控制器
type MonitorController struct {
	BaseController
}

// GetServerRate 获取服务器内存、CPU和磁盘的使用率
//func (m *MonitorController) GetServerRate(ctx *gin.Context) {
//	var (
//		result    response.ServerUsageResponse
//		customErr *response.BusinessError
//	)
//	if result, customErr = service.Monitor.GetServerRate(); customErr != nil {
//		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
//		return
//	}
//	ctx.JSON(http.StatusOK, response.Ok(result))
//}

// GetCPUInfo 获取服务器CPU信息
func (m *MonitorController) GetCPUInfo(ctx *gin.Context) {
	var (
		result    *response.CpuUsageResponse
		customErr *response.BusinessError
	)
	if result, customErr = service.Monitor.GetCpuInfo(1); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// GetMemInfo 获取服务器内存信息
func (m *MonitorController) GetMemInfo(ctx *gin.Context) {
	var (
		result    *response.MemUsageResponse
		customErr *response.BusinessError
	)
	if result, customErr = service.Monitor.GetMemoryInfo(); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// GetServerInfo 获取服务器信息
func (m *MonitorController) GetServerInfo(ctx *gin.Context) {
	var (
		result    *response.ServerInfoResponse
		customErr *response.BusinessError
	)
	if result, customErr = service.Monitor.GetServerInfo(); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}
