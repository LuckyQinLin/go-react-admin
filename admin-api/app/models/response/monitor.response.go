package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

type (
	// ServerUsageResponse 服务器占用率
	ServerUsageResponse struct {
		Cpu    float64 // cpu占用
		Memory float64 // 内存占用
		Disk   float64 // 磁盘占用
	}

	// ServerInfoResponse 服务器信息
	ServerInfoResponse struct {
		Hostname        string `json:"hostname"`
		Uptime          int64  `json:"runTime"`         // 运行时间
		Procs           uint64 `json:"procs"`           // 进程数
		OS              string `json:"os"`              // 操作系统名称
		Platform        string `json:"platform"`        // 操作系统平台
		PlatformVersion string `json:"platformVersion"` // 操作系统版本
		KernelVersion   string `json:"kernelVersion"`   // 内核版本
		KernelArch      string `json:"kernelArch"`      // 内核架构
	}

	// CpuUsageResponse CPU占用
	CpuUsageResponse struct {
		Time *time.Time `json:"time"`
		Num  float64    `json:"num"`
	}

	// MemUsageResponse 内存占用
	MemUsageResponse struct {
		Time    *time.Time `json:"time"`    // 时间
		Total   uint64     `json:"total"`   // 总内存
		Used    uint64     `json:"used"`    // 使用
		Free    uint64     `json:"free"`    // 空闲
		Percent float64    `json:"percent"` // 占比
	}
)

func (s CpuUsageResponse) MarshalJSON() ([]byte, error) {
	type temp CpuUsageResponse
	return json.Marshal(&struct {
		temp
		Time utils.DateTime1 `json:"time"`
	}{
		temp: (temp)(s),
		Time: utils.DateTime1(*s.Time),
	})
}

func (s MemUsageResponse) MarshalJSON() ([]byte, error) {
	type temp MemUsageResponse
	return json.Marshal(&struct {
		temp
		Time utils.DateTime1 `json:"time"`
	}{
		temp: (temp)(s),
		Time: utils.DateTime1(*s.Time),
	})
}
