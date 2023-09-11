// UserPageQueryProp 用户分页查询
import {RuleObject} from "rc-field-form/lib/interface";

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
    password: string;   // 密码
    deptId:   number;   // 部门ID
    phone?:   string;   // 手机号
    email?:   string;   // 邮箱
    sex:      number;   // 性别
    status:   number;   // 状态
    postId?:  number[]; // 岗位
    roleId?:  number[]; // 角色
    remark?:  string;   // 备注
}

// UserUpdateFormProp 用户创建表单
interface UserUpdateFormProp {
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

// UserDrawerProp 用户抽屉属性
interface UserDrawerProp {
    createVisible: boolean;
    updateVisible: boolean;
    roleVisible: boolean;
    userId?: number;
}

// 验证手机号码
const validateMobile = (_: RuleObject, value: string, callback: (error?: string) => void) => {
    if (value !== "") {
        const phoneReg = /^1[3456789]\d{9}$/;
        if (!phoneReg.test(value)) {
            callback('手机号码格式不正确，请重新输入');
        }
    }
    callback();
};
// 验证邮箱
const validateEmail = (_: RuleObject, value: string, callback: (error?: string) => void) => {
    debugger;
    if (value !== "") {
        const mailReg = /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$/;
        if (!mailReg.test(value)) {
            // return Promise.reject('邮箱格式不正确，请重新输入');
            callback('邮箱格式不正确，请重新输入');
        }
    }
    callback();
};

interface UserRoleDrawerProp {
    userId?: number;
    visible: boolean;
    close: () => void;
}

// UserLoginInfoProp 用户登录信息
interface UserLoginInfoProp {
    user: UserProp;
    roles: string[];
    permissions: string[];
}

interface DeptProp {
    deptId: number;
    parentId: number;
    deptName: string;
    leader: string;
    ancestors: string;
    orderNum: number;
    status: number;
}

interface RoleProp {
    roleId: number;
    roleName: string;
    roleCode: string;
}

interface PostProp {
    postId: number;
    postName: string;
    postCode: string;
}

interface UserProp {
    admin: boolean;
    avatar: string;
    userId: number;
    userName: string;
    sex: number;
    phone: string;
    nickName: string;
    email: string;
    deptId: number;
    dept: DeptProp;
    roles: RoleProp[];
    posts: PostProp[];
}


export {
    validateMobile,
    validateEmail
}

export type {
    UserPageQueryProp,
    UserCreateDrawerProp,
    UserUpdateDrawerProp,
    UserCreateFormProp,
    UserUpdateFormProp,
    UserTableProp,
    UserDrawerProp,
    UserRoleDrawerProp,
    UserLoginInfoProp
}