import {https, PageData} from "@/utils/request.ts";
import {NoticeCreateFormProp, NoticePageProp, NoticePageQueryProp, NoticeUpdateFormProp} from "@/pages/system/inform/modules.ts";

// noticePage 字典分页查询
export const noticePage = (data: NoticePageQueryProp): Promise<PageData<NoticePageProp>> => {
    return https.request({
        url: '/notice/page',
        method: 'post',
        data: data
    })
}

// noticeCreate 创建字典
export const noticeCreate = (data: NoticeCreateFormProp): Promise<boolean> => {
    return https.request({
        url: '/notice/create',
        method: 'post',
        data: data
    })
}

// noticeUpdate 修改字典
export const noticeUpdate = (data: NoticeUpdateFormProp): Promise<boolean> => {
    return https.request({
        url: '/notice/update',
        method: 'post',
        data: data
    })
}

// noticeInfo 字典详情
export const noticeInfo = (noticeId: number): Promise<NoticeUpdateFormProp> => {
    return https.request({
        url: '/notice/info',
        method: 'get',
        params: {noticeId: noticeId}
    })
}

// noticeDelete 字典删除
export const noticeDelete = (ids: number[]): Promise<boolean> => {
    return https.request({
        url: '/notice/delete',
        method: 'post',
        data: {ids: ids}
    })
}