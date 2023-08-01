import {DeptCreateDrawerProp, DeptCreateFormProp, DeptTreeProp} from "@/pages/system/dept/modules.ts";
import {Button, Drawer, Form, Input, InputNumber, message, Radio, Space, Spin, TreeSelect} from "antd";
import React, {useEffect, useState} from "react";
import {useRequest} from "ahooks";
import {deptCreate, deptTree} from "@/api/dept.ts";

const DeptCreateDrawer: React.FC<DeptCreateDrawerProp> = ({visible, parentId, close}) => {

    const [form] = Form.useForm<DeptCreateFormProp>();
    const [tree, setTree] = useState<DeptTreeProp[]>([]);
    const loadTree = useRequest(deptTree, {manual: true, onSuccess: (data) => {
            setTree([{key: 0, title: '主类目', children: data}]);
        }
    });

    const createDept = useRequest(deptCreate, {manual: true, onSuccess: (_) => {
            message.success('创建部门成功')
            close(true)
        }
    });

    const submitForm = () => {
        form.validateFields().then(value => {
            createDept.run(value)
        })
    };

    useEffect(() => {
        if (visible) {
            loadTree.run();
            form.setFieldsValue({parentId: parentId, orderNum: 1, status: 1})
        }
        return () => {
            form.resetFields();
            setTree([]);
        }
    }, [visible])

    return <Drawer
        width={500}
        title="添加部门"
        placement="right"
        onClose={() => close(false)}
        open={visible}
        extra={
            <Space>
                <Button type="primary" danger onClick={()=> close(false)}>取消</Button>
                <Button type="primary" loading={createDept.loading} onClick={submitForm}>保存</Button>
            </Space>
        }
    >
        <Spin tip="加载中......" spinning={loadTree.loading}>
            <Form labelCol={{ span: 5 }}
                  wrapperCol={{ span: 19 }}
                  form={form}
                  layout="horizontal"
                  name="form_in_modal">
                <Form.Item name="parentId" label="上级部门" rules={[{ required: true, message: '请选择上级部门' }]}>
                    <TreeSelect
                        allowClear
                        treeData={tree}
                        placeholder="请选择上级部门"
                        fieldNames={{label: 'title', value: 'key', children: 'children'}}
                    />
                </Form.Item>
                <Form.Item name="deptName" label="部门名称" rules={[{ required: true, message: '请输入部门名称' }]}>
                    <Input placeholder="请输入部门名称" />
                </Form.Item>
                <Form.Item name="orderNum" label="显示顺序" rules={[{ required: true, message: '请输入实现顺序' }]}>
                    <InputNumber min={1} placeholder="请输入实现顺序" style={{width: '100%'}} />
                </Form.Item>
                <Form.Item name="leader" label="负责人">
                    <Input placeholder="请输入部门负责人" />
                </Form.Item>
                <Form.Item name="phone" label="联系电话">
                    <Input placeholder="请输入联系电话" />
                </Form.Item>
                <Form.Item name="email" label="邮箱">
                    <Input placeholder="请输入邮箱" />
                </Form.Item>
                <Form.Item name="status" label="部门状态">
                    <Radio.Group>
                        <Radio value={1}>正常</Radio>
                        <Radio value={0}>停用</Radio>
                    </Radio.Group>
                </Form.Item>
            </Form>
        </Spin>
    </Drawer>
}

export default DeptCreateDrawer;