import React, {useEffect, useState} from "react";
import Logo from '@/assets/images/account-logo.png';
import './index.less';
import {Button, Checkbox, Form, Input, message, Tabs, TabsProps} from "antd";
import {
    GithubOutlined,
    LockOutlined,
    MailOutlined,
    QqOutlined,
    SafetyOutlined,
    UserOutlined,
    VerifiedOutlined,
    WechatOutlined
} from "@ant-design/icons";
import LoginBg from '@/assets/images/login-bg-1.svg';
import {useNavigate} from "react-router-dom";
import {captchaImage, sendCaptcha, userRegister} from "@/api/user";
import {LoginCaptchaProp, RegisterFormProp} from "@/pages/login/modules";
import {HOME_PAGE} from "@/constant/setting.ts";
import {User} from "@/types";
import useStore from "@/store/store.ts";


export interface LoginFormProp {
    username: string;
    password: string;
    captcha: string;
    uuid: string;
}

interface LoginProp {
    submit: (data: User.LoginFormProp) => void;
    changeRegister: () => void;
}

const LoginForm: React.FC<LoginProp> = ({submit, changeRegister}) => {

    const [form] = Form.useForm<User.LoginFormProp>();

    const onSubmit = () => {
        form.validateFields().then(value => submit({...value, uuid: captchaProp.uuid}))
    }

    const [captchaProp, setCaptchaProp] = useState<LoginCaptchaProp>({} as LoginCaptchaProp);

    const toRegister = () => {
        form.resetFields();
        changeRegister()
    }

    const loadCaptchaImage = async () => {
        const result = await captchaImage();
        setCaptchaProp(result);
    }

    useEffect(() => {
        loadCaptchaImage();
    }, [])

    return <Form className="account-login-form" form={form}>
        <Form.Item name="username" rules={[{ required: true, message: '请输入账号' }]}>
            <Input size="large" prefix={<UserOutlined />} placeholder="账号"/>
        </Form.Item>
        <Form.Item name="password" rules={[{ required: true, message: '请输入密码' }]}>
            <Input.Password size="large" prefix={<LockOutlined />} placeholder="密码"/>
        </Form.Item>
        <Form.Item name="captcha" rules={[{ required: true, message: '请输入验证码' }]}>
            <div className="account-login-form-captcha">
                <Input className="alfc-left" size="large" prefix={<SafetyOutlined />} placeholder="验证码"/>
                <img onClick={() => loadCaptchaImage()} className="alfc-right" src={captchaProp.image} />
            </div>
        </Form.Item>
        <Form.Item className="account-login-item-1">
            <Checkbox>自动登录</Checkbox>
            <Button type="link" style={{padding: '4px 0px'}}>忘记密码</Button>
        </Form.Item>
        <Form.Item>
            <Button type="primary" block size="large" onClick={onSubmit}>登 录</Button>
        </Form.Item>
        <Form.Item className="account-login-item-1">
            <div className="account-login-method">
                <div>其他登录方式</div>
                <WechatOutlined style={{fontSize: 20, marginRight: 5, marginLeft: 10}} />
                <QqOutlined style={{fontSize: 20, marginRight: 5}} />
                <GithubOutlined style={{fontSize: 20}} />
            </div>
            <Button type="link" style={{padding: '4px 0px'}} onClick={toRegister}>注册账号</Button>
        </Form.Item>
    </Form>
}


interface RegisterProp {
    submit: (value: RegisterFormProp) => void;
    changeRegister: () => void;
}

const RegisterForm: React.FC<RegisterProp> = ({submit, changeRegister}) => {

    const [form] = Form.useForm<RegisterFormProp>();

    const onSubmit = () => {
        form.validateFields().then(value => submit(value))
    }

    // sendCaptcha 发送验证码
    const getCaptcha = async () => {
        const value = form.getFieldValue("email");
        await sendCaptcha(value)
    }

    const toLogin = () => {
        form.resetFields();
        changeRegister()
    }

    return <Form className="account-register-form" form={form}>
        <Form.Item name="username" rules={[{ required: true, message: '请输入用户名' }]}>
            <Input size="large" prefix={<UserOutlined />} placeholder="用户名"/>
        </Form.Item>
        <Form.Item name="email" rules={[
            { type: 'email', message: '邮箱格式不正确' },
            { required: true, message: '请输入邮箱' },
        ]}>
            <div className="account-register-email">
                <Input size="large" prefix={<MailOutlined />} placeholder="邮箱"/>
                <Button size="large"
                        type="primary"
                        style={{marginLeft: 10}}
                        onClick={getCaptcha}
                >获取验证码</Button>
            </div>
        </Form.Item>
        <Form.Item name="captcha" rules={[{ required: true, message: '请输入验证码' }]}>
            <Input size="large" prefix={<VerifiedOutlined />} placeholder="验证码"/>
        </Form.Item>
        <Form.Item name="password" rules={[{ required: true, message: '请输入密码' }]}>
            <Input.Password size="large" prefix={<LockOutlined />} placeholder="密码"/>
        </Form.Item>
        <Form.Item
            name="confirm"
            rules={[
                { required: true, message: '确认密码' },
                ({ getFieldValue }) => ({
                    validator(_, value) {
                        if (!value || getFieldValue('password') === value) {
                            return Promise.resolve();
                        }
                        return Promise.reject(new Error('两次密码不一致'));
                    },
                }),
            ]}>
            <Input.Password size="large" prefix={<LockOutlined />} placeholder="确认密码"/>
        </Form.Item>
        <Form.Item>
            <Button type="primary" block size="large" onClick={onSubmit}>登 录</Button>
        </Form.Item>
        <Form.Item>
            <Button block size="large" onClick={toLogin}>返 回</Button>
        </Form.Item>
    </Form>
}

const LoginPage: React.FC = () => {

    const userLoginFetch = useStore((state) => state.userLoginFetch)


    const navigate = useNavigate();


    const loginHandle = async (value: LoginFormProp) => {
        try {
            // setLoading(true);
            userLoginFetch(value)
                .then(() => {navigate(HOME_PAGE)})
                .catch(err => {console.log(err.message)});
        } finally {
            // setLoading(false)
        }

    }

    const registerHandle = async (value: RegisterFormProp) => {
        await userRegister(value)
        message.info('注册成功')
        setActiveKey('1')
    }

    const items: TabsProps['items'] = [
        {
            key: '1',
            label: `登录`,
            children: <LoginForm submit={loginHandle} changeRegister={() => setActiveKey('2')} />,
        },
        {
            key: '2',
            label: `注册`,
            children: <RegisterForm submit={registerHandle} changeRegister={() => setActiveKey('1')} />,
        }
    ];

    const [activeKey, setActiveKey] = useState<string>('1');

    return <div className="account-root">
        {/*<Spin spinning={loading} >*/}
            <div className="account-root-item">
                <div className="account-root-item-img">
                    <img src={LoginBg} alt="" />
                    {/*<img src="https://pro.naiveadmin.com/assets/login-bg.be83cd61.svg" alt="" />*/}
                </div>
            </div>
            <div className="account-root-item root-right-item">
                <div className="account-form">
                    <div className="account-top">
                        <div className="account-top-logo">
                            <img src={Logo} alt="" />
                        </div>
                        <div className="account-top-desc">一款通用的后台管理系统</div>
                    </div>
                    <Tabs className="account-tabs" activeKey={activeKey} onChange={setActiveKey} centered items={items} />
                </div>
            </div>
        {/*</Spin>*/}
    </div>
}

export default LoginPage;
