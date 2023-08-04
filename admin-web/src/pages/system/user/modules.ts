// UserPageQueryProp 用户分页查询
interface UserPageQueryProp {
    page:      number;
    size:      number;
    deptId?:   number; // 部门
    status?:   number; // 用户状态
    userName?: string; // 用户名称
    phone?:    string; // 手机号
}

// UserTableProp 用户表属性
interface UserTableProp {
    userId:     number; // 用户ID
    userName:   string; // 用户名称
    nickName:   string; // 昵称
    deptName:   string; // 部门名称
    phone:      string; // 手机号
    status:     number; // 状态
    createTime: string; // 创建时间
}

// UserCreateDrawerProp 用户创建抽屉属性
interface UserCreateDrawerProp {
    visible: boolean;
    deptId?: number;
    close: (isLoad: boolean) => void;
}

// UserUpdateDrawerProp 用户更新抽屉属性
interface UserUpdateDrawerProp {
    userId?: number;
    visible: boolean;
    close: (isLoad: boolean) => void;
}

// UserCreateFormProp 用户创建表单
interface UserCreateFormProp {
    userName: string;   // 用户名称
    nickName: string;   // 用户昵称
    deptId:   number;   // 部门ID
    phone?:   string;   // 手机号
    email?:   string;   // 邮箱
    sex:      number;   // 性别
    status:   number;   // 状态
    postId?:  number[]; // 岗位
    roleId?:  number[]; // 角色
    remark?:  string;   // 备注
}

export type {
    UserPageQueryProp,
    UserCreateDrawerProp,
    UserUpdateDrawerProp,
    UserCreateFormProp,
    UserTableProp
}