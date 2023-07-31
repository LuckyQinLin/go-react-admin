import {Button, Drawer, Form, Input, InputNumber, message, Radio, Space, Spin} from "antd";
import React, {useEffect} from "react";
import {PostUpdateDrawerProp, PostUpdateFormProp} from "@/pages/system/post/modules.ts";
import {useRequest} from "ahooks";
import {postUpdate, postInfo} from "@/api/post.ts";

const PostUpdateDrawer: React.FC<PostUpdateDrawerProp> = ({visible, postId, close}) => {

    const [form] = Form.useForm<PostUpdateFormProp>();
    const loadInfo = useRequest(postInfo, {
        manual: true,
        onSuccess: (data) => form.setFieldsValue({...data})
    });

    const updatePost = useRequest(postUpdate, {
        manual: true,
        onSuccess: (_) => {
            message.success('修改岗位成功')
            close(true)
        }
    });

    const submitForm = () => {
        form.validateFields().then(value => {
            // console.log("update", value)
            updatePost.run(value)
        })
    };

    useEffect(() => {
        if (visible && postId) {
            loadInfo.run(postId)
        }
        return () => {
            form.resetFields();
        }
    }, [visible])

    return <Drawer
        width={500}
        title="修改岗位"
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
        <Spin tip="加载中......" spinning={updatePost.loading}>
            <Form labelCol={{ span: 5 }}
                  wrapperCol={{ span: 19 }}
                  form={form}
                  layout="horizontal"
                  name="form_in_modal"
            >
                <Form.Item name="postId" label="主键" style={{display: 'none'}}>
                    <Input />
                </Form.Item>
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

export default PostUpdateDrawer;