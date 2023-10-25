namespace Monitor {
	export interface ServerInfoResponse {
		hostname: string;
		runTime: number; // 运行时间，使用 number 表示
		procs: number; // 进程数，使用 number 表示
		os: string; // 操作系统名称
		platform: string; // 操作系统平台
		platformVersion: string; // 操作系统版本
		kernelVersion: string; // 内核版本
		kernelArch: string; // 内核架构
		virtualizationSystem: string; // 虚拟化系统
		virtualizationRole: string; // 指示主机是虚拟机中的“guest”（客户机）还是虚拟机宿主机（"host"）
	}

	// CpuInfoResponse cpu占用率
	export interface CpuInfoResponse {
		time: string; // 时间
		num: number; // 数值
	}

	// MemUsageResponse 内存占用
	export interface MemUsageResponse {
		time: string;
		total: number;
		used: number;
		free: number;
		percent: number;
	}

	// MemAreaField 内存面积图属性
	export interface MemAreaField {
		data: number;
		time: string;
		types: string;
	}
}

export default Monitor;
