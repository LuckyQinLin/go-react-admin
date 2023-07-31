import {https, PageData} from "@/utils/request.ts";
import {ConfigCreateFormProp, ConfigPageProp, ConfigPageQueryProp, ConfigUpdateFormProp} from "@/pages/system/param/modules.ts";

// configPage 字典分页查询
export const configPage = (data: ConfigPageQueryProp): Promise<PageData<ConfigPageProp>> => {
    return https.request({
        url: '/config/page',
        method: 'post',
        data: data
    })
}

// configCreate 创建字典
export const configCreate = (data: ConfigCreateFormProp): Promise<boolean> => {
    return https.request({
        url: '/config/create',
        method: 'post',
        data: data
    })
}

// configUpdate 修改字典
export const configUpdate = (data: ConfigUpdateFormProp): Promise<boolean> => {
    return https.request({
        url: '/config/update',
        method: 'post',
        data: data
    })
}

// configInfo 字典详情
export const configInfo = (configId: number): Promise<ConfigUpdateFormProp> => {
    return https.request({
        url: '/config/info',
        method: 'get',
        params: {configId: configId}
    })
}

// configDelete 字典删除
export const configDelete = (ids: number[]): Promise<boolean> => {
    return https.request({
        url: '/config/delete',
        method: 'post',
        data: {ids: ids}
    })
}