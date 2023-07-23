import {RoleCreateFormProp, RolePageProp, RolePageQueryProp} from "@/pages/system/role/modules.ts";
import {PageData} from "@/utils/request.ts";
import {https} from "@/utils/request.ts";

// rolePage 角色分页查询
export const rolePage = (data: RolePageQueryProp): Promise<PageData<RolePageProp>> => {
    return https.request({
        url: '/role/page',
        method: 'post',
        data: data
    })
}

// roleCreate 创建角色
export const roleCreate = (data: RoleCreateFormProp): Promise<boolean> => {
    return https.request({
        url: '/role/create',
        method: 'post',
        data: data
    })
}