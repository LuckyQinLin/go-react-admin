import React, {useState} from "react";
import {Button, Drawer, Form, Input, InputNumber, message, Space, Spin, Radio, Checkbox} from "antd";
import {RoleCreateDrawerProp, RoleCreateFormProp} from "@/pages/system/role/modules.ts";
import {useRequest} from "ahooks";
import {roleCreate} from "@/api/role.ts";

// RoleCreateDrawer 角色创建
const RoleCreateDrawer: React.FC<RoleCreateDrawerProp> = ({visible, close}) => {

    const [expendFlag, setExpendFlag] = useState<boolean>(false);
    const [allFlag, setAllFlag] = useState<boolean>(false);
    const [linkedFlag, setLinkedFlag] = useState<boolean>(false);
    const [form] = Form.useForm<RoleCreateFormProp>();
    const {loading, run} = useRequest(roleCreate, {
        manual: true,
        onSuccess: (_)=> {
            message.success('创建角色成功');
            close(true);
        }
    })

    const submitForm = () => form.validateFields().then(value => run(value));

    return <Drawer
        width={500}
        title="创建角色"
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
                <Form.Item name="roleName" label="角色名称" rules={[{ required: true, message: '请输入角色名称' }]}>
                    <Input placeholder="请输入角色名称" />
                </Form.Item>
                <Form.Item name="roleKey" label="权限字符" rules={[{ required: true, message: '请输入权限字符' }]}>
                    <Input placeholder="请输入权限字符" />
                </Form.Item>
                <Form.Item name="roleSort" label="角色顺序" rules={[{ required: true, message: '请输入角色顺序' }]}>
                    <InputNumber defaultValue={1} min={1} placeholder="请输入角色顺序" style={{width: '100%'}} />
                </Form.Item>
                <Form.Item name="status" label="状态">
                    <Radio.Group defaultValue={1}>
                        <Radio value={0}>停用</Radio>
                        <Radio value={1}>正常</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="menuIds" label="菜单权限">
                    <Checkbox.Group>
                        <Checkbox value={0}>展开/折叠</Checkbox>
                        <Checkbox value={1}>全选/全不选</Checkbox>
                        <Checkbox value={2}>父子联动</Checkbox>
                    </Checkbox.Group>
                </Form.Item>
                <Form.Item name="remark" label="备注">
                    <Input placeholder="请输入备注" />
                </Form.Item>
            </Form>
        </Spin>

    </Drawer>
}

export default RoleCreateDrawer;