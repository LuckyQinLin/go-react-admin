import {https} from "@/utils/request.ts";
import {
    DeptCreateFormProp,
    DeptTableQueryProp,
    DeptTableTreeProp,
    DeptTreeProp,
    DeptUpdateFormProp
} from "@/pages/system/dept/modules.ts";

// menuTree 部门树
export const deptTree = (): Promise<DeptTreeProp[]> => {
    return https.request({
        url: '/dept/tree',
        method: 'get',
    })
}

// menuTree 部门创建
export const deptCreate = (data: DeptCreateFormProp): Promise<string> => {
    return https.request({
        url: '/dept/create',
        method: 'post',
        data: data
    })
}

// menuTree 部门修改
export const deptUpdate = (data: DeptUpdateFormProp): Promise<string> => {
    return https.request({
        url: '/dept/update',
        method: 'post',
        data: data
    })
}

// menuTree 部门详情
export const deptInfo = (deptId: number): Promise<DeptUpdateFormProp> => {
    return https.request({
        url: '/dept/info',
        method: 'get',
        params: {deptId: deptId}
    })
}

// menuTree 部门树
export const deptTable = (data: DeptTableQueryProp): Promise<DeptTableTreeProp[]> => {
    return https.request({
        url: '/dept/table',
        method: 'post',
        data: data
    })
}

// deptDelete 部门删除
export const deptDelete = (deptId: number): Promise<string> => {
    return https.request({
        url: '/dept/delete',
        method: 'get',
        params: {deptId: deptId}
    })
}