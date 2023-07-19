
// 下载资源
import {SoftwareState} from "@/redux/software/reducer";
import {changeMenuActionCreator} from "@/redux/system/action";

export const DOWNLOAD_SOFTWARE = "download_software";
// 下载完成
export const DOWNLOAD_END = "download_end";

export type SoftwareTypes = DownloadSoftwareAction | DownloadEndAction;

// 下载资源
export interface DownloadSoftwareAction {
    type: typeof DOWNLOAD_SOFTWARE,
    payload: SoftwareState,
}

export interface DownloadEndAction {
    type: typeof DOWNLOAD_END,
    payload: SoftwareState,
}


export const downloadSoftwareActionCreator = (state: SoftwareState): DownloadSoftwareAction => {
    return {
        type: DOWNLOAD_SOFTWARE,
        payload: state,
    }
}

export const downloadEndActionCreator = (state: SoftwareState): DownloadEndAction => {
    return {
        type: DOWNLOAD_END,
        payload: state,
    }
}
