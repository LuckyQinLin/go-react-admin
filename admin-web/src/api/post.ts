import {https, PageData} from "@/utils/request.ts";
import {PostCreateFormProp, PostPageProp, PostPageQueryProp, PostUpdateFormProp} from "@/pages/system/post/modules.ts";

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