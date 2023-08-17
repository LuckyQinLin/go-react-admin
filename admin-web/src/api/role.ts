import {
    RoleCreateFormProp,
    RoleInfoProp,
    RolePageProp,
    RolePageQueryProp,
    RoleUpdateFormProp
} from "@/pages/system/role/modules.ts";
import {PageData} from "@/utils/request.ts";
import {https} from "@/utils/request.ts";

// postList 全部岗位
export const roleList = (): Promise<RoleInfoProp[]> => {
    return https.request({
        url: '/role/all',
        method: 'get'
    })
}

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

// roleInfo 角色信息
export const roleInfo = (id: number): Promise<RoleUpdateFormProp> => {
    return https.request({
        url: '/role/info',
        method: 'get',
        params: {id: id}
    })
}

// roleCreate 修改角色
export const roleUpdate = (data: RoleUpdateFormProp): Promise<boolean> => {
    return https.request({
        url: '/role/update',
        method: 'post',
        data: data
    })
}

// roleDelete 删除角色
export const roleDelete = (ids: number[]): Promise<boolean> => {
    return https.request({
        url: '/role/delete',
        method: 'post',
        data: {ids: ids}
    })
}

// roleByUserId 获取用户的角色
export const roleByUserId = (userId: number): Promise<number[]> => {
    return https.request({
        url: '/role/user',
        method: 'get',
        params: {userId: userId}
    })
}

// roleByUserId 获取用户的角色
export const getRoleUser = (roleId: number): Promise<number[]> => {
    return https.request({
        url: '/role/getUser',
        method: 'get',
        params: {roleId: roleId}
    })
}

// roleAllocateUser 角色分配用户
export const roleAllocateUser = (roleId: number, userIds: number[]): Promise<boolean> => {
    return https.request({
        url: '/role/allocateUser',
        method: 'post',
        params: {roleId: roleId, userIds: userIds}
    })
}