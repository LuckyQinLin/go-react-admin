import {Button, Form, Input, Radio} from "antd";
import {UpdateBaseInfoProp} from "@/pages/user/modules.ts";

const BaseInfoForm = () => {

    const buttonItemLayout = { wrapperCol: { span: 14, offset: 3 } };
    const [form] = Form.useForm<UpdateBaseInfoProp>();

    return <Form labelCol={{ span: 3 }}
                 wrapperCol={{ span: 21 }}
                 form={form}
                 layout="horizontal"
                 name="form_in_modal">
        <Form.Item name="nickName" label="用户昵称" rules={[{ required: true, message: '请输入用户昵称' }]}>
            <Input style={{width: 400}} placeholder="请输入用户昵称" />
        </Form.Item>
        <Form.Item name="phone" label="手机号码" rules={[{ required: true, message: '请输入手机号码' }]}>
            <Input style={{width: 400}} placeholder="请输入手机号码" />
        </Form.Item>
        <Form.Item name="email" label="邮箱地址" rules={[{ required: true, message: '请输入邮箱地址' }]}>
            <Input style={{width: 400}} placeholder="请输入邮箱地址" />
        </Form.Item>
        <Form.Item name="sex" label="性别" rules={[{ required: true, message: '请选择性别' }]}>
            <Radio.Group>
                <Radio value={1}>男</Radio>
                <Radio value={0}>女</Radio>
            </Radio.Group>
        </Form.Item>
        <Form.Item {...buttonItemLayout}>
            <Button type="primary" danger style={{marginRight: 10}}>重置</Button>
            <Button type="primary">保存</Button>
        </Form.Item>
    </Form>
}

export default BaseInfoForm;