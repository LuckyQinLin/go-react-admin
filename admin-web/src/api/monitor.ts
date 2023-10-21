import {https} from "@/utils/request.ts";
import Monitor from "@/types/monitor.ts";

// serverInfo 获取服务器信息
export const serverInfo = (): Promise<Monitor.ServerInfoResponse> => {
	return https.request({
		url: '/monitor/serverInfo',
		method: 'get'
	})
}