export interface RegisterFormProp {
    username: string;
    email: string;
    captcha: string;
    password: string;
    confirm: string;
}

export interface LoginCaptchaProp {
    uuid: string; // uuid
    image: string; // 图片
    expireTime: number;// 过期时间
}