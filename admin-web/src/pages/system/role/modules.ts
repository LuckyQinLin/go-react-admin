interface RolePageQueryProp {
    page: number;
    size: number;
    name?: string;
    status?: number;
    startTime?: number;
    endTime?: number;
}

interface RoleInfoProp {
    value: number;
    label: string;
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

// RoleUpdateDrawerProp 角色更新
interface RoleUpdateDrawerProp {
    roleId?: number;
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

// RoleUpdateFormProp 角色修改表单属性
interface RoleUpdateFormProp extends RoleCreateFormProp{
    roleId: number; // 角色ID
}

// RoleDrawerProp 角色抽屉属性
interface RoleDrawerProp {
    createVisible: boolean;
    updateVisible: boolean;
    roleId?: number;
}

// AllocateUserDrawerProp 分配用户
interface AllocateUserDrawerProp {
    roleId: number;
    visible: boolean;
    close: () => void;
}

// AllocatePermDrawerProp 分配资源
interface AllocatePermDrawerProp {
    visible: boolean;
    roleId: number;
    close: () => void;
}


export type {
    RoleInfoProp,
    RolePageQueryProp,
    RolePageProp,
    RoleCreateDrawerProp,
    RoleUpdateDrawerProp,
    RoleCreateFormProp,
    RoleUpdateFormProp,
    RoleDrawerProp,
    AllocateUserDrawerProp,
    AllocatePermDrawerProp
}