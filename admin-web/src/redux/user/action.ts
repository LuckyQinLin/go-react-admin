import {defaultUserState, UserState} from "./reducer";

export const CHANGE_LONGIN_STATUS = "change_login_status";
export const CLEAN_USER_STORE = "clean_user_store";
export const CHANGE_MENU_STATUS = "change_menu_status";

interface ChangLoginAction {
	type: typeof CHANGE_LONGIN_STATUS;
	payload: UserState;
}

interface CleanUserAction {
	type: typeof CLEAN_USER_STORE;
	payload: UserState;
}

interface ChangeMenuAction {
	type: typeof CHANGE_MENU_STATUS;
	payload: UserState;
}


export type UserTypes = ChangLoginAction | CleanUserAction | ChangeMenuAction;

export const changeLoginStatusActionCreator = (status: UserState): ChangLoginAction => {
	return {
		type: CHANGE_LONGIN_STATUS,
		payload: status
	}
}

export const changeMenuStatusActionCreator = (status: UserState): ChangeMenuAction => {
	return {
		type: CHANGE_MENU_STATUS,
		payload: status
	}
}

export const cleanUserStoreActionCreator = (): CleanUserAction => {
	return {
		type: CLEAN_USER_STORE,
		payload: defaultUserState
	}
}
