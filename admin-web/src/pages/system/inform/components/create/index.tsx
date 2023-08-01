import {Button, Drawer, Form, Input, message, Radio, Space, Spin} from "antd";
import React, {useEffect} from "react";
import {useRequest} from "ahooks";
import {NoticeCreateDrawerProp, NoticeCreateFormProp} from "@/pages/system/inform/modules.ts";
import {noticeCreate} from "@/api/notice.ts";

const InformCreateDrawer: React.FC<NoticeCreateDrawerProp> = ({visible, close}) => {

    const [form] = Form.useForm<NoticeCreateFormProp>();

    const {loading, run} = useRequest(noticeCreate, {
        manual: true,
        onSuccess: (_)=> {
            message.success('创建公告成功');
            close(true);
        }
    })

    const submitForm = () => form.validateFields().then(value => {
        run(value)
    });

    useEffect(() => {
        if (visible) {
            form.setFieldsValue({status: 1})
        }
    }, [visible])

    return <Drawer
        width={500}
        title="创建公告"
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
                <Form.Item name="noticeTitle" label="公告名称" rules={[{ required: true, message: '请输入公告名称' }]}>
                    <Input placeholder="请输入公告名称" />
                </Form.Item>
                <Form.Item name="noticeType" label="公告类型" rules={[{ required: true, message: '请输入公告类型' }]}>
                    <Radio.Group>
                        <Radio value={1}>通知</Radio>
                        <Radio value={2}>公告</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="noticeContent" label="内容">
                    <Input.TextArea placeholder="请输入内容" />
                </Form.Item>
                <Form.Item name="status" label="状态">
                    <Radio.Group>
                        <Radio value={1}>正常</Radio>
                        <Radio value={0}>停用</Radio>
                    </Radio.Group>
                </Form.Item>
            </Form>
        </Spin>
    </Drawer>
}

export default InformCreateDrawer;