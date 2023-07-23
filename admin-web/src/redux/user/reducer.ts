import {CHANGE_LONGIN_STATUS, CLEAN_USER_STORE, UserTypes} from "./action";

export interface UserState {
	status: boolean;
	userId?: number; // 用户ID
	sex?: number; // 性别
	username?: string; // 用户名称
	nickName?: string; // 用户昵称
	avatar?: string;// 头像
	deptId?: number; // 部门ID
	phone?: string; // 手机号
	email?: string; // 邮箱
	remark?: string; // 备注
	token?: string;
	expireTime?: number; // 到期时间
	roles?: RoleInfo[]; // 角色信息
	perms?: PermInfo[]; // 资源
}

export interface RoleInfo {
	id: number;
	code: string;
	name: string;
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

export const defaultUserState: UserState = {status: false}

export default function userReducer(state = defaultUserState, action: UserTypes) {
	switch (action.type) {
		case CHANGE_LONGIN_STATUS:
			return {
				...state,
				...action.payload
			};
		case CLEAN_USER_STORE:
			return {
				...state,
				...action.payload,
			}
		default:
			return state;
	}
}
