import {Button, Drawer, Form, Input, message, Radio, Space, Spin} from "antd";
import React, {useEffect} from "react";
import {DictUpdateDrawerProp, DictUpdateFormProp} from "@/pages/system/dict/modules.ts";
import {useRequest} from "ahooks";
import {dictUpdate, dictInfo} from "@/api/dict.ts";

const DictUpdateDrawer: React.FC<DictUpdateDrawerProp> = ({visible, dictId, close}) => {

    const [form] = Form.useForm<DictUpdateFormProp>();
    const loadInfo = useRequest(dictInfo, {
        manual: true,
        onSuccess: (data) => form.setFieldsValue({...data})
    });

    const updateDict = useRequest(dictUpdate, {
        manual: true,
        onSuccess: (_) => {
            message.success('修改字典成功')
            close(true)
        }
    });

    const submitForm = () => {
        form.validateFields().then(value => {
            // console.log("update", value)
            updateDict.run(value)
        })
    };

    useEffect(() => {
        if (visible && dictId) {
            loadInfo.run(dictId)
        }
        return () => {
            form.resetFields();
        }
    }, [visible])

    return <Drawer
        width={500}
        title="修改字典"
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
        <Spin tip="加载中......" spinning={updateDict.loading}>
            <Form labelCol={{ span: 5 }}
                  wrapperCol={{ span: 19 }}
                  form={form}
                  layout="horizontal"
                  name="form_in_modal"
            >
                <Form.Item name="dictId" label="主键" style={{display: 'none'}}>
                    <Input />
                </Form.Item>
                <Form.Item name="dictName" label="字典名称" rules={[{ required: true, message: '请输入字典名称' }]}>
                    <Input placeholder="请输入字典名称" />
                </Form.Item>
                <Form.Item name="dictType" label="字典类型" rules={[{ required: true, message: '请输入字典类型' }]}>
                    <Input placeholder="请输入字典类型" />
                </Form.Item>
                <Form.Item name="status" label="字典状态">
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

export default DictUpdateDrawer;