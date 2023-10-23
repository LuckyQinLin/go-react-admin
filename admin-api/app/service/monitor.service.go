package service

import (
	"admin-api/app/models/response"
	"admin-api/core"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"time"
)

var Monitor = new(MonitorService)

type MonitorService struct{}

// GetServerRate 获取服务器占用情况
//func (m *MonitorService) GetServerRate() (response.ServerUsageResponse, *response.BusinessError) {
//	percent, _ := cpu.Percent(time.Second, false)
//	memory, _ := mem.VirtualMemory()
//	parts, _ := disk.Partitions(true)
//	disk, _ := disk.Usage(parts[0].Mountpoint)
//	return response.ServerUsageResponse{
//		Cpu:    percent[0],
//		Memory: memory.UsedPercent,
//		Disk:   disk.UsedPercent,
//	}, nil
//}

// GetCpuInfo 获取CPU占用率
func (m *MonitorService) GetCpuInfo(interval int64) (*response.CpuUsageResponse, *response.BusinessError) {
	var (
		timeValue time.Time
		percent   []float64
		err       error
	)
	if percent, err = cpu.Percent(time.Duration(interval)*time.Second, false); err != nil {
		return nil, response.CustomBusinessError(response.Failed, "获取服务器CPU占用率出错")
	}
	timeValue = time.Now()
	return &response.CpuUsageResponse{
		Time: &timeValue,
		Num:  percent[0],
	}, nil
}

// GetServerInfo 获取系统信息
func (m *MonitorService) GetServerInfo() (*response.ServerInfoResponse, *response.BusinessError) {
	var (
		info *host.InfoStat
		err  error
	)
	if info, err = host.Info(); err != nil {
		core.Log.Error("获取服务器信息出错:%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取服务器信息出错")
	}
	return &response.ServerInfoResponse{
		Hostname:        info.Hostname,
		Uptime:          time.Now().Add(time.Duration(info.Uptime)).Unix(),
		Procs:           info.Procs,
		OS:              info.OS,
		Platform:        info.Platform,
		PlatformVersion: info.PlatformVersion,
		KernelVersion:   info.KernelVersion,
		KernelArch:      info.KernelArch,
	}, nil
}
