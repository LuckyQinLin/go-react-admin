export interface WsState {
    connected: boolean;
    token?: string;
    is_login: boolean;
}

export const defaultWsState: WsState = {
    connected: false,
    is_login: false
}



