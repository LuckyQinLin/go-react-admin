import {https} from "@/utils/request.ts";
import {MenuTreeProp} from "@/pages/system/menu/modules.ts";


// menuTree 菜单树
export const menuTree = (): Promise<MenuTreeProp[]> => {
    return https.request({
        url: '/menu/tree',
        method: 'get',
    })
}