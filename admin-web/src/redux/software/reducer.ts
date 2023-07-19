import {DOWNLOAD_END, DOWNLOAD_SOFTWARE, SoftwareTypes} from "@/redux/software/action";


export interface SoftwareState {
    id?: number;
    is_result?: boolean;
}

export const defaultSoftwareState: SoftwareState = {
    id: undefined,
    is_result: false
}

export default function softwareReducer(state = defaultSoftwareState, action: SoftwareTypes) {
    switch (action.type) {
        case DOWNLOAD_SOFTWARE:
            return {
                ...state,
                ...action.payload
            };
        case DOWNLOAD_END:
            return {
                ...state,
                ...action.payload,
            }
        default:
            return state;
    }
}
