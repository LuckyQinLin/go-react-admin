import {Button, Drawer, Form, Input, InputNumber, Radio, Space, Spin, TreeSelect} from "antd";
import React, {useState} from "react";
import {MenuCreateFormProp, MenuTreeProp, MenuUpdateDrawerProp} from "@/pages/system/menu/modules.ts";
import {useRequest} from "ahooks";
import {menuTree} from "@/api/menu.ts";

const MenuUpdateDrawer: React.FC<MenuUpdateDrawerProp> = ({visible, menuId, close}) => {

    const [form] = Form.useForm<MenuCreateFormProp>();
    const [tree, setTree] = useState<MenuTreeProp[]>([]);
    const loadTree = useRequest(menuTree, {manual: true, onSuccess: (data) => setTree(data)});
    const submitForm = () => {}

    return <Drawer
        width={500}
        title="修改菜单"
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
        <Spin tip="加载中......" spinning={loadTree.loading}>
            <Form labelCol={{ span: 5 }}
                  wrapperCol={{ span: 19 }}
                  form={form}
                  layout="horizontal"
                  name="form_in_modal">
                <Form.Item name="menuName" label="菜单名称" rules={[{ required: true, message: '请选择上级菜单名称' }]}>
                    <TreeSelect
                        showSearch
                        style={{ width: '100%' }}
                        dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                        placeholder="请选择上级菜单名称"
                        allowClear
                        treeDefaultExpandAll
                        treeData={tree}
                    />
                </Form.Item>
                <Form.Item name="menuType" label="菜单类型">
                    <Radio.Group>
                        <Radio value={0}>目录</Radio>
                        <Radio value={1}>菜单</Radio>
                        <Radio value={2}>按钮</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="icon" label="菜单图标" rules={[{ required: true, message: '请选择菜单图标' }]}>
                    <Input placeholder="请选择菜单图标" />
                </Form.Item>
                <Form.Item name="menuName" label="菜单名称" rules={[{ required: true, message: '请输入菜单名称' }]}>
                    <Input placeholder="请输入菜单名称" />
                </Form.Item>
                <Form.Item name="path" label="路由地址" rules={[{ required: true, message: '请输入路由地址' }]}>
                    <Input placeholder="请输入路由地址" />
                </Form.Item>
                <Form.Item name="menuSort" label="显示顺序" rules={[{ required: true, message: '请输入实现顺序' }]}>
                    <InputNumber min={1} placeholder="请输入实现顺序" style={{width: '100%'}} />
                </Form.Item>
                <Form.Item name="isLink" label="是否外链">
                    <Radio.Group>
                        <Radio value={true}>是</Radio>
                        <Radio value={false}>否</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="show" label="显示状态">
                    <Radio.Group>
                        <Radio value={true}>显示</Radio>
                        <Radio value={false}>隐藏</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="show" label="菜单状态">
                    <Radio.Group>
                        <Radio value={true}>正常</Radio>
                        <Radio value={false}>停用</Radio>
                    </Radio.Group>
                </Form.Item>
            </Form>
        </Spin>
    </Drawer>
}

export default MenuUpdateDrawer;