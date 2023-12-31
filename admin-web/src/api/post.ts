import {https, PageData} from "@/utils/request.ts";
import {
    PostCreateFormProp,
    PostInfoProp,
    PostPageProp,
    PostPageQueryProp,
    PostUpdateFormProp
} from "@/pages/system/post/modules.ts";

// postList 全部岗位
export const postList = (): Promise<PostInfoProp[]> => {
    return https.request({
        url: '/post/all',
        method: 'get'
    })
}

// postPage 岗位分页查询
export const postPage = (data: PostPageQueryProp): Promise<PageData<PostPageProp>> => {
    return https.request({
        url: '/post/page',
        method: 'post',
        data: data
    })
}

// postCreate 创建岗位
export const postCreate = (data: PostCreateFormProp): Promise<boolean> => {
    return https.request({
        url: '/post/create',
        method: 'post',
        data: data
    })
}

// postUpdate 修改岗位
export const postUpdate = (data: PostUpdateFormProp): Promise<boolean> => {
    return https.request({
        url: '/post/update',
        method: 'post',
        data: data
    })
}

// postInfo 岗位详情
export const postInfo = (postId: number): Promise<PostUpdateFormProp> => {
    return https.request({
        url: '/post/info',
        method: 'get',
        params: {postId: postId}
    })
}

// postDelete 岗位删除
export const postDelete = (ids: number[]): Promise<boolean> => {
    return https.request({
        url: '/post/delete',
        method: 'post',
        data: {ids: ids}
    })
}