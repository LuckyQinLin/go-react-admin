import {Button, Drawer, Form, Input, message, Radio, Space, Spin} from "antd";
import React, {useEffect} from "react";
import {useRequest} from "ahooks";
import {ConfigCreateDrawerProp, ConfigCreateFormProp} from "@/pages/system/param/modules.ts";
import {configCreate} from "@/api/config.ts";

const ParamCreateDrawer: React.FC<ConfigCreateDrawerProp> = ({visible, close}) => {

    const [form] = Form.useForm<ConfigCreateFormProp>();

    const {loading, run} = useRequest(configCreate, {
        manual: true,
        onSuccess: (_)=> {
            message.success('创建参数成功');
            close(true);
        }
    })

    const submitForm = () => form.validateFields().then(value => {
        run(value)
    });

    useEffect(() => {
        if (visible) {
            form.setFieldsValue({configType: 1})
        }
    }, [visible])

    return <Drawer
        width={500}
        title="创建参数"
        placement="right"
        onClose={() => close(false)}
        open={visible}
        extra={
            <Space>
                <Button type="primary" danger onClick={()=> close(false)}>取消</Button>
                <Button type="primary" onClick={submitForm}>保存</Button>
            </Space>
        }
    >
        <Spin tip="加载中......" spinning={loading}>
            <Form labelCol={{ span: 5 }}
                  wrapperCol={{ span: 19 }}
                  form={form}
                  layout="horizontal"
                  name="form_in_modal"
            >
                <Form.Item name="configName" label="参数名称" rules={[{ required: true, message: '请输入参数名称' }]}>
                    <Input placeholder="请输入参数名称" />
                </Form.Item>
                <Form.Item name="configKey" label="参数键名" rules={[{ required: true, message: '请输入参数键名' }]}>
                    <Input placeholder="请输入参数键名" />
                </Form.Item>
                <Form.Item name="configValue" label="参数键值" rules={[{ required: true, message: '请输入参数键值' }]}>
                    <Input placeholder="请输入参数键值" />
                </Form.Item>
                <Form.Item name="configType" label="系统内置">
                    <Radio.Group>
                        <Radio value={1}>是</Radio>
                        <Radio value={0}>否</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="remark" label="备注">
                    <Input.TextArea placeholder="请输入备注" />
                </Form.Item>
            </Form>
        </Spin>
    </Drawer>
}

export default ParamCreateDrawer;