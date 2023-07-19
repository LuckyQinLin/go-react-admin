import {defaultUserState, UserState} from "./reducer";

export const CHANGE_LONGIN_STATUS = "change_login_status";
export const CLEAN_USER_STORE = "clean_user_store";

interface ChangLoginAction {
	type: typeof CHANGE_LONGIN_STATUS;
	payload: UserState;
}

interface CleanUserAction {
	type: typeof CLEAN_USER_STORE;
	payload: UserState;
}


export type UserTypes = ChangLoginAction | CleanUserAction;

export const changeLoginStatusActionCreator = (status: UserState): ChangLoginAction => {
	return {
		type: CHANGE_LONGIN_STATUS,
		payload: status
	}
}

export const cleanUserStoreActionCreator = (): CleanUserAction => {
	return {
		type: CLEAN_USER_STORE,
		payload: defaultUserState
	}
}
