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
    roleName: string; // 角色名称
    roleKey: string; // 权限字符
    roleSort: number; // 角色排序
    status: number; // 状态
    menuIds?: number[]; // 授权菜单ids
    remark?: string; // 备注
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