import {LoginFormProp} from "@/pages/login";
import {UserState} from "@/redux/user/reducer.ts";
import {https} from "@/utils/request.ts";
import {RegisterFormProp} from "@/pages/login/modules.ts";

export const userLogin = (data: LoginFormProp): Promise<UserState> => {
    return https.request({
        url: '/user/login',
        method: 'post',
        data: data
    }, {})
}

// sendCaptcha 发送验证码
export const sendCaptcha = (data: string): Promise<string> => {
    return https.request({
        url: '/user/captcha',
        method: 'get',
        params: {email: data}
    })
}

// userRegister 用户注册
export const userRegister = (data: RegisterFormProp): Promise<boolean> => {
    return https.request({
        url: '/user/register',
        method: 'post',
        data: data
    })
}