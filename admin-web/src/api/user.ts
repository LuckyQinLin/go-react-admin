import {https, PageData} from "@/utils/request.ts";
import {LoginCaptchaProp, RegisterFormProp} from "@/pages/login/modules.ts";
import {
    UserCreateFormProp, UserLoginInfoProp,
    UserPageQueryProp,
    UserTableProp,
    UserUpdateFormProp
} from "@/pages/system/user/modules.ts";
import {User} from "@/types";

// userLogin 用户登录
export const userLogin = (data: User.LoginFormProp): Promise<User.UserLoginResponse> => {
    return https.request({
        url: '/user/login',
        method: 'post',
        data: data
    }, {isShowSuccessMessage: true})
}

// captchaImage 验证码
export const captchaImage = (): Promise<LoginCaptchaProp> => {
    return https.request({
        url: '/user/captchaImage',
        method: 'get',
    }, {})
}

// sendCaptcha 发送邮箱验证码
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

// userInfo 获取用户信息
export const userInfo = (userId: number): Promise<UserUpdateFormProp> => {
    return https.request({
        url: '/user/info',
        method: 'get',
        params: {userId: userId}
    })
}

// userPage 用户分页查询
export const userPage = (data: UserPageQueryProp): Promise<PageData<UserTableProp>> => {
    return https.request({
        url: '/user/page',
        method: 'post',
        data: data
    })
}

// userCreate 用户创建
export const userCreate = (data: UserCreateFormProp): Promise<string> => {
    return https.request({
        url: '/user/create',
        method: 'post',
        data: data
    })
}

// userCreate 用户修改
export const userUpdate = (data: UserUpdateFormProp): Promise<string> => {
    return https.request({
        url: '/user/update',
        method: 'post',
        data: data
    })
}

// userDelete 用户删除
export const userDelete = (ids: number[]): Promise<string> => {
    return https.request({
        url: '/user/delete',
        method: 'post',
        data: {ids: ids}
    })
}

// userDelete 用户状态
export const userStatus = (userId: number, status: boolean): Promise<string> => {
    return https.request({
        url: '/user/status',
        method: 'post',
        data: {userId: userId, status: status ? 1 : 0}
    })
}

// userRole 用户分配角色
export const userRole = (userId: number, ids: number[]): Promise<string> => {
    return https.request({
        url: '/user/role',
        method: 'post',
        data: {userId: userId, ids: ids}
    })
}

// userLoginInfo 获取用户登录信息
export const userLoginInfo = (): Promise<User.UserInfoProp> => {
    return https.request({
        url: '/user/loginInfo',
        method: 'get',
    })
}
