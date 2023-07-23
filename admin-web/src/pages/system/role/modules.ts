interface RolePageQueryProp {
    page: number;
    size: number;
    name?: string;
    status?: number;
    startTime?: number;
    endTime?: number;
}

// RolePageProp 角色分页查询返回属性
interface RolePageProp {
    roleId: number;
    roleName: string;
    roleKey: string;
    roleSort: number;
    status: number;
    createTime: string;
}

// RoleCreateDrawerProp 角色创建属性
interface RoleCreateDrawerProp {
    visible: boolean;
    close: (isLoad: boolean) => void;
}

// RoleCreateFormProp 角色创建表单属性
interface RoleCreateFormProp {

}

// RoleDrawerProp 角色抽屉属性
interface RoleDrawerProp {
    createVisible: boolean;
}

export type {
    RolePageQueryProp,
    RolePageProp,
    RoleCreateDrawerProp,
    RoleCreateFormProp,
    RoleDrawerProp
}