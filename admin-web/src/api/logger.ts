// rolePage 角色分页查询
import {
    LoginLoggerQueryProp,
    LoginLoggerTableProp,
    OperateLoggerQueryProp,
    OperateLoggerTableProp
} from "@/pages/logger/modules.ts";
import {PageData} from "@/utils/request.ts";
import {https} from "@/utils/request.ts";

// loginLoggerPage 访问日志
export const loginLoggerPage = (data: LoginLoggerQueryProp): Promise<PageData<LoginLoggerTableProp>> => {
    return https.request({
        url: '/logger/visit',
        method: 'post',
        data: data
    })
}

// operateLoggerPage 操作日志
export const operateLoggerPage = (data: OperateLoggerQueryProp): Promise<PageData<OperateLoggerTableProp>> => {
    return https.request({
        url: '/logger/operate',
        method: 'post',
        data: data
    })
}