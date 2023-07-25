import {https} from "@/utils/request.ts";
import {MenuTableTreeProp, MenuTableTreeQueryProp, MenuTreeProp} from "@/pages/system/menu/modules.ts";

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