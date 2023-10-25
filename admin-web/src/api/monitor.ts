import {https} from "@/utils/request.ts";
import Monitor from "@/types/monitor.ts";

// serverInfo 获取服务器信息
export const cpuInfo = (): Promise<Monitor.CpuInfoResponse> => {
	return https.request({
		url: '/monitor/cpuInfo',
		method: 'get'
	})
}

// serverInfo 获取服务器信息
export const memInfo = (): Promise<Monitor.MemUsageResponse> => {
	return https.request({
		url: '/monitor/memInfo',
		method: 'get'
	})
}

// serverInfo 获取服务器信息
export const serverInfo = (): Promise<Monitor.ServerInfoResponse> => {
	return https.request({
		url: '/monitor/serverInfo',
		method: 'get'
	})
}

