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

export type {
    UserPageQueryProp,
    UserTableProp
}