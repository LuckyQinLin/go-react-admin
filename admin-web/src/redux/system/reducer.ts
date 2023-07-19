import {SystemConfig} from '@/config'
import {RESET_SYSTEM, SET_MENU_KEY, SystemTypes} from "@/redux/system/action";

export interface SystemState {
    menuKey: string;
}

export const defaultSystemState: SystemState = {
    menuKey: SystemConfig.DEFAULT_MENU_KEY
}

export default function systemReducer(state = defaultSystemState, action: SystemTypes) {
    switch (action.type) {
        case SET_MENU_KEY:
            return {
                ...state,
                ...action.payload
            };
        case RESET_SYSTEM:
            return {
                ...state,
                ...action.payload,
            }
        default:
            return state;
    }
}
