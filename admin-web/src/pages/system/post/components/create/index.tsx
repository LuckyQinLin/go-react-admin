import {Button, Drawer, Form, Input, InputNumber, message, Radio, Space, Spin} from "antd";
import React, {useEffect} from "react";
import {PostCreateDrawerProp, PostCreateFormProp} from "@/pages/system/post/modules.ts";
import {useRequest} from "ahooks";
import {postCreate} from "@/api/post.ts";

const PostCreateDrawer: React.FC<PostCreateDrawerProp> = ({visible, close}) => {

    const [form] = Form.useForm<PostCreateFormProp>();

    const {loading, run} = useRequest(postCreate, {
        manual: true,
        onSuccess: (_)=> {
            message.success('创建岗位成功');
            close(true);
        }
    })

    const submitForm = () => form.validateFields().then(value => {
        run(value)
    });

    useEffect(() => {
        if (visible) {
            form.setFieldsValue({postSort: 1, status: 1})
        }
    }, [visible])

    return <Drawer
        width={500}
        title="创建岗位"
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
                <Form.Item name="postName" label="岗位名称" rules={[{ required: true, message: '请输入岗位名称' }]}>
                    <Input placeholder="请输入岗位名称" />
                </Form.Item>
                <Form.Item name="postCode" label="岗位编码" rules={[{ required: true, message: '请输入岗位编码' }]}>
                    <Input placeholder="请输入岗位编码" />
                </Form.Item>
                <Form.Item name="postSort" label="岗位顺序" rules={[{ required: true, message: '请输入岗位顺序' }]}>
                    <InputNumber min={1} placeholder="请输入岗位顺序" style={{width: '100%'}} />
                </Form.Item>
                <Form.Item name="status" label="岗位状态">
                    <Radio.Group>
                        <Radio value={1}>正常</Radio>
                        <Radio value={0}>停用</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="remark" label="备注">
                    <Input.TextArea placeholder="请输入备注" />
                </Form.Item>
            </Form>
        </Spin>
    </Drawer>
}

export default PostCreateDrawer;