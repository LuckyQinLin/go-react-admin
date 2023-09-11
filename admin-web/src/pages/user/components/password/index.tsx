import {Button, Form, Input} from "antd";
import {UpdatePasswordProp} from "@/pages/user/modules.ts";

const PasswordForm = () => {

    const buttonItemLayout = { wrapperCol: { span: 14, offset: 3 } };
    const [form] = Form.useForm<UpdatePasswordProp>();

    return <Form labelCol={{ span: 3 }}
                 wrapperCol={{ span: 21 }}
                 form={form}
                 layout="horizontal"
                 name="form_in_modal">
        <Form.Item name="oldPassword" label="旧密码" rules={[{ required: true, message: '请输入旧密码' }]}>
            <Input.Password style={{width: 400}} placeholder="请输入旧密码" />
        </Form.Item>
        <Form.Item name="newPassword" label="新密码" rules={[{ required: true, message: '请输入新密码' }]}>
            <Input.Password style={{width: 400}} placeholder="请输入新密码" />
        </Form.Item>
        <Form.Item name="confirmPassword" label="确认密码" rules={[{ required: true, message: '请再次输入新密码' }]}>
            <Input.Password style={{width: 400}} placeholder="请再次输入新密码" />
        </Form.Item>
        <Form.Item {...buttonItemLayout}>
            <Button type="primary" danger style={{marginRight: 10}}>重置</Button>
            <Button type="primary">保存</Button>
        </Form.Item>
    </Form>
}

export default PasswordForm;