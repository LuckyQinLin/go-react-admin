import {CHANGE_LONGIN_STATUS, CLEAN_USER_STORE, UserTypes} from "./action";

export interface UserState {
	status?: boolean;
	token?: string;
	roleType?: number;
	username?: string;
	phone?: string;
	email?: string;
	userId?: string;
	roles?: RoleInfo[];
	perms?: PermInfo[];
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

export const defaultUserState: UserState = {
	status: false,
	token: undefined,
	roleType: undefined,
	username: undefined,
	phone: undefined,
	email: undefined,
	userId: undefined
}

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
