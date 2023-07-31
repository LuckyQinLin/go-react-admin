import {https, PageData} from "@/utils/request.ts";
import {DictCreateFormProp, DictPageProp, DictPageQueryProp, DictUpdateFormProp} from "@/pages/system/dict/modules.ts";

// dictPage 字典分页查询
export const dictPage = (data: DictPageQueryProp): Promise<PageData<DictPageProp>> => {
    return https.request({
        url: '/dict/page',
        method: 'post',
        data: data
    })
}

// dictCreate 创建字典
export const dictCreate = (data: DictCreateFormProp): Promise<boolean> => {
    return https.request({
        url: '/dict/create',
        method: 'post',
        data: data
    })
}

// dictUpdate 修改字典
export const dictUpdate = (data: DictUpdateFormProp): Promise<boolean> => {
    return https.request({
        url: '/dict/update',
        method: 'post',
        data: data
    })
}

// dictInfo 字典详情
export const dictInfo = (dictId: number): Promise<DictUpdateFormProp> => {
    return https.request({
        url: '/dict/info',
        method: 'get',
        params: {dictId: dictId}
    })
}

// dictDelete 字典删除
export const dictDelete = (ids: number[]): Promise<boolean> => {
    return https.request({
        url: '/dict/delete',
        method: 'post',
        data: {ids: ids}
    })
}