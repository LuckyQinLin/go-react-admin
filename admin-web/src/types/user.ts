import Post from "@/types/post.ts";
import Role from "@/types/role.ts";
import Dept from "@/types/dept.ts";
import Menu from "@/types/menu.ts";
import Menus from "@/types/menu.ts";

namespace User {

    export interface UserStoreProp {
        tabViewKey: string;
        tabViews: Menus.TabViewProp[];
        loginProp?: UserLoginResponse;
        userProp?: UserInfoProp;
        setLoginProp: (data: UserLoginResponse) => void;
        setUserProp: (data: UserInfoProp) => void;
        userLoginFetch: (data: LoginFormProp) => Promise<void>;
        useInfoFetch: () => Promise<void>;
        addTabView: (data: Menus.TabViewProp) => void;
        removeTabView: (key: string, isNegation?: boolean) => void;
        setTabViewKey: (key: string) => void;
        closeTabViewAll: () => void;
    }

    // LoginFormProp 用户登录表单
    export interface LoginFormProp {
        username: string;
        password: string;
        captcha: string;
        uuid: string;
    }

    // 用户详情
    export interface UserInfoProp {
        isSuper: boolean;   // 是否admin
        avatar: string;   // 头像
        userId: number;   // 用户ID
        userName: string; // 用户名称
        sex: number;      // 性别
        phone: string;    // 手机号
        nickName: string; // 昵称
        email: string;    // 邮箱
        deptId: number;   // 部门ID
        dept: Dept.UserDeptProp;   // 部门信息
        roles: Role.UserRoleProp[];// 角色
        posts: Post.UserPostProp[];// 岗位
        operates: string[]; // 操作
    }

    // 用户权限信息
    export interface UserPermissionProp {
        menus: Menu.MenuItemProp[]; // 菜单数据
        buttons: string[];          // 按钮
        paths: string[];            // 路径

    }

    // UserLoginResponse 用户登录返回
    export interface UserLoginResponse {
        token: string;      // token信息
        expireTime: number; // 到期时间
    }

}

export default User
