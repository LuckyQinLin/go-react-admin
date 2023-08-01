import {Button, Drawer, Form, Input, message, Radio, Space, Spin} from "antd";
import React, {useEffect} from "react";
import {NoticeUpdateDrawerProp, NoticeUpdateFormProp} from "@/pages/system/inform/modules.ts";
import {useRequest} from "ahooks";
import {noticeUpdate, noticeInfo} from "@/api/notice.ts";

const DictUpdateDrawer: React.FC<NoticeUpdateDrawerProp> = ({visible, noticeId, close}) => {

    const [form] = Form.useForm<NoticeUpdateFormProp>();
    const loadInfo = useRequest(noticeInfo, {
        manual: true,
        onSuccess: (data) => form.setFieldsValue({...data})
    });

    const updateNotice = useRequest(noticeUpdate, {
        manual: true,
        onSuccess: (_) => {
            message.success('修改公告成功')
            close(true)
        }
    });

    const submitForm = () => {
        form.validateFields().then(value => {
            // console.log("update", value)
            updateNotice.run(value)
        })
    };

    useEffect(() => {
        if (visible && noticeId) {
            loadInfo.run(noticeId)
        }
        return () => {
            form.resetFields();
        }
    }, [visible])

    return <Drawer
        width={500}
        title="修改公告"
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
        <Spin tip="加载中......" spinning={updateNotice.loading}>
            <Form labelCol={{ span: 5 }}
                  wrapperCol={{ span: 19 }}
                  form={form}
                  layout="horizontal"
                  name="form_in_modal"
            >
                <Form.Item name="noticeId" label="主键" style={{display: 'none'}}>
                    <Input />
                </Form.Item>
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

export default DictUpdateDrawer;