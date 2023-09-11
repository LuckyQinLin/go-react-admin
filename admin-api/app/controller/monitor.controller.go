package controller

// https://github.com/shirou/gopsutil
var Monitor = new(MonitorController)

// MonitorController 监控控制器
type MonitorController struct {
	BaseController
}
