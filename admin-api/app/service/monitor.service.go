package service

import (
	"admin-api/app/models/response"
	"admin-api/core"
	"github.com/shirou/gopsutil/host"
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
		Uptime:          info.Uptime,
		Procs:           info.Procs,
		OS:              info.OS,
		Platform:        info.Platform,
		PlatformVersion: info.PlatformVersion,
		KernelVersion:   info.KernelVersion,
		KernelArch:      info.KernelArch,
	}, nil
}
