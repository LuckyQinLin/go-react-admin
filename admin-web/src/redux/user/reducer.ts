import {CHANGE_LONGIN_STATUS, CHANGE_MENU_STATUS, CLEAN_USER_STORE, UserTypes} from "./action";
// import {MenuProps} from "antd";
// import {IRouteObject} from "@/router/modules.ts";
import {UserRouterProp} from "@/pages/system/menu/modules.ts";

// type MenuItem = Required<MenuProps>['items'][number];

export interface UserState {
	token?: string;
	expireTime?: number; // 到期时间
	user?: UserProp; // 用户信息
	roles?: string[]; // 角色
	permissions?: string[]; // 资源
	userRouter?: UserRouterProp[]; // 菜单
}

export interface PermInfo {
	id: number;
	name: string;
	code: string;
	icon?: string;
	types: number; // 1:目录 2:菜单 3:按钮
	level: number;
	parentId: number;
	path?: string;
	orderNo: number;
	children?: PermInfo[];
}

export interface DeptProp {
	deptId: number;
	parentId: number;
	deptName: string;
	leader: string;
	ancestors: string;
	orderNum: number;
	status: number;
}

export interface RoleProp {
	roleId: number;
	roleName: string;
	roleCode: string;
}

export interface PostProp {
	postId: number;
	postName: string;
	postCode: string;
}

export interface UserProp {
	admin: boolean;   // 是否admin
	avatar: string;   // 头像
	userId: number;   // 用户ID
	userName: string; // 用户名称
	sex: number;      // 性别
	phone: string;    // 手机号
	nickName: string; // 昵称
	email: string;    // 邮箱
	deptId: number;   // 部门ID
	dept: DeptProp;   // 部门信息
	roles: RoleProp[];// 角色
	posts: PostProp[];// 岗位
}

export const defaultUserState: UserState = {}

export default function userReducer(state = defaultUserState, action: UserTypes) {
	switch (action.type) {
		case CHANGE_LONGIN_STATUS:
			return {
				...state,
				...action.payload
			};
		case CLEAN_USER_STORE:
			// return {
			// 	...state,
			// 	token: undefined
			// };
			return defaultUserState
		case CHANGE_MENU_STATUS:
			return {
				...state,
				...action.payload
			}
		default:
			return state;
	}
}
