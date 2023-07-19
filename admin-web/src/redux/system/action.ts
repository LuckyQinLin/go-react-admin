import {defaultSystemState, SystemState} from "@/redux/system/reducer";

export const SET_MENU_KEY = "set_menu_key";

export const RESET_SYSTEM = "reset_system";

export interface ChangeMenuAction {
    type: typeof SET_MENU_KEY,
    payload: SystemState
}

export interface RestSystemAction {
    type: typeof RESET_SYSTEM,
    payload: SystemState
}

export type SystemTypes = ChangeMenuAction | RestSystemAction;

export const changeMenuActionCreator = (state: SystemState): ChangeMenuAction => {
    return {
        type: SET_MENU_KEY,
        payload: state
    }
}

export const resetSystemActionCreator = (): ChangeMenuAction => {
    return {
        type: SET_MENU_KEY,
        payload: defaultSystemState
    }
}
