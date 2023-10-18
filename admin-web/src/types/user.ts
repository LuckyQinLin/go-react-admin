import Post from "@/types/post.ts";
import Role from "@/types/role.ts";
import Dept from "@/types/dept.ts";
import Menu from "@/types/menu.ts";

namespace User {
    // 用户详情
    export interface UserInfoProp {
        admin: boolean;   // 是否admin
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
    }

    // 用户权限信息
    export interface UserPermissionProp {
        menus: Menu.MenuItemProp[]; // 菜单数据
        buttons: string[];          // 按钮
        paths: string[];            // 路径

    }
}

export default User
