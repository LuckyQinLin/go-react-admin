package response

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
		Uptime          uint64 `json:"runTime"`         // 运行时间
		Procs           uint64 `json:"procs"`           // 进程数
		OS              string `json:"os"`              // 操作系统名称
		Platform        string `json:"platform"`        // 操作系统平台
		PlatformVersion string `json:"platformVersion"` // 操作系统版本
		KernelVersion   string `json:"kernelVersion"`   // 内核版本
		KernelArch      string `json:"kernelArch"`      // 内核架构
	}
)
