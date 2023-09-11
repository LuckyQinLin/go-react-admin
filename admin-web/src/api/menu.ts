import {https} from "@/utils/request.ts";
import {
    MenuCreateFormProp,
    MenuTableTreeProp,
    MenuTableTreeQueryProp,
    MenuTreeProp, MenuUpdateFormProp, UserRouterProp
} from "@/pages/system/menu/modules.ts";

// menuTree 菜单树
export const menuTable = (data: MenuTableTreeQueryProp): Promise<MenuTableTreeProp[]> => {
    return https.request({
        url: '/menu/table',
        method: 'post',
        data: data
    })
}

// menuTree 菜单树
export const menuTree = (): Promise<MenuTreeProp[]> => {
    return https.request({
        url: '/menu/tree',
        method: 'get',
    })
}

// menuTree 菜单创建
export const menuCreate = (data: MenuCreateFormProp): Promise<string> => {
    return https.request({
        url: '/menu/create',
        method: 'post',
        data: data
    })
}

// menuTree 菜单修改
export const menuUpdate = (data: MenuUpdateFormProp): Promise<string> => {
    return https.request({
        url: '/menu/update',
        method: 'post',
        data: data
    })
}

// menuTree 菜单详情
export const menuInfo = (menuId: number): Promise<MenuUpdateFormProp> => {
    return https.request({
        url: '/menu/info',
        method: 'get',
        params: {menuId: menuId}
    })
}

// menuDelete 菜单删除
export const menuDelete = (menuId: number): Promise<string> => {
    return https.request({
        url: '/menu/delete',
        method: 'get',
        params: {menuId: menuId}
    })
}

// menuDelete 菜单删除
export const userRouter = (roleId?: number): Promise<UserRouterProp[]> => {
    return https.request({
        url: '/menu/router',
        method: 'get',
        params: {roleId: roleId}
    })
}