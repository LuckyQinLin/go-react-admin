import {Button, Drawer, Form, Input, InputNumber, message, Radio, Space, Spin, TreeSelect} from "antd";
import React, {useEffect, useState} from "react";
import {
    MenuTreeProp,
    MenuUpdateDrawerProp,
    MenuUpdateFormProp
} from "@/pages/system/menu/modules.ts";
import {useRequest} from "ahooks";
import {menuInfo, menuTree, menuUpdate} from "@/api/menu.ts";

const MenuUpdateDrawer: React.FC<MenuUpdateDrawerProp> = ({visible, menuId, close}) => {

    const [menuType, setMenuType] = useState<string>('M');
    const [form] = Form.useForm<MenuUpdateFormProp>();
    const [tree, setTree] = useState<MenuTreeProp[]>([]);
    const loadTree = useRequest(menuTree, {
        manual: true,
        onSuccess: (data) => setTree([{key: 0, title: '主类目', children: data}])
    });

    const loadInfo = useRequest(menuInfo, {
        manual: true,
        onSuccess: (data) => form.setFieldsValue({...data})
    });

    const updateTree = useRequest(menuUpdate, {
        manual: true,
        onSuccess: (_) => {
            message.success('修改菜单成功')
            close(true)
        }
    });


    const changeMenuType = (e: string) => {
        setMenuType(e);
        if (e === 'M') {
            form.setFieldsValue({isLink: false, isShow: true, status: 1, menuType: e})
        } else if (e === 'C') {
            form.setFieldsValue({isLink: false, isShow: true, status: 1, isCache: true, menuType: e})
        } else if (e === 'F') {
            form.setFieldsValue({isShow: false, status: 1, menuType: e})
        }
    }

    const submitForm = () => {
        form.validateFields().then(value => {
            // console.log("update", value)
            updateTree.run(value)
        })
    }

    useEffect(() => {
        if (visible && menuId) {
            loadTree.run();
            loadInfo.run(menuId)
        }
        return () => {
            form.resetFields();
            setTree([]);
        }
    }, [visible])

    return <Drawer
        width={500}
        title="修改菜单"
        placement="right"
        onClose={() => close(false)}
        open={visible}
        extra={
            <Space>
                <Button type="primary" danger onClick={()=> close(false)}>取消</Button>
                <Button type="primary" loading={updateTree.loading} onClick={submitForm}>保存</Button>
            </Space>
        }
    >
        <Spin tip="加载中......" spinning={loadInfo.loading || loadTree.loading}>
            <Form labelCol={{ span: 5 }}
                  wrapperCol={{ span: 19 }}
                  form={form}
                  layout="horizontal"
                  name="form_in_modal">
                <Form.Item name="menuId" label="主键" style={{display: 'none'}} rules={[{ required: true, message: '请输入菜单名称' }]}>
                    <Input placeholder="请输入主键" />
                </Form.Item>
                <Form.Item name="parentId" label="上级菜单" rules={[{ required: true, message: '请选择上级菜单' }]}>
                    <TreeSelect
                        allowClear
                        treeData={tree}
                        placeholder="请选择上级菜单"
                        fieldNames={{label: 'title', value: 'key', children: 'children'}}
                    />
                </Form.Item>
                <Form.Item name="menuName" label="菜单名称" rules={[{ required: true, message: '请输入菜单名称' }]}>
                    <Input placeholder="请输入菜单名称" />
                </Form.Item>
                <Form.Item name="menuType" label="菜单类型">
                    <Radio.Group onChange={e => changeMenuType(e.target.value as string)}>
                        <Radio value={'M'}>目录</Radio>
                        <Radio value={'C'}>菜单</Radio>
                        <Radio value={'F'}>按钮</Radio>
                    </Radio.Group>
                </Form.Item>
                {menuType !== 'F' && <Form.Item name="icon" label="菜单图标">
                    <Input placeholder="请选择菜单图标" />
                </Form.Item>}
                {menuType !== 'M' && <Form.Item name="perms" label="权限字符">
                    <Input placeholder="请输入权限字符" />
                </Form.Item>}
                {menuType !== 'F' && <Form.Item name="path" label="路由地址" rules={[{ required: true, message: '请输入路由地址' }]}>
                    <Input placeholder="请输入路由地址" />
                </Form.Item>}
                {menuType === 'C' && <Form.Item name="component" label="组件路径">
                    <Input placeholder="请输入组件路径" />
                </Form.Item>}
                {menuType === 'C' && <Form.Item name="param" label="路由参数">
                    <Input placeholder="请输入路由参数" />
                </Form.Item>}
                <Form.Item name="menuSort" label="显示顺序" rules={[{ required: true, message: '请输入实现顺序' }]}>
                    <InputNumber min={1} placeholder="请输入实现顺序" style={{width: '100%'}} />
                </Form.Item>
                {menuType !== 'F' && <Form.Item name="isLink" label="是否外链">
                    <Radio.Group>
                        <Radio value={true}>是</Radio>
                        <Radio value={false}>否</Radio>
                    </Radio.Group>
                </Form.Item>}
                {menuType !== 'F' && <Form.Item name="isShow" label="显示状态">
                    <Radio.Group>
                        <Radio value={true}>显示</Radio>
                        <Radio value={false}>隐藏</Radio>
                    </Radio.Group>
                </Form.Item>}
                {menuType === 'C' && <Form.Item name="isCache" label="是否缓冲">
                    <Radio.Group>
                        <Radio value={true}>缓冲</Radio>
                        <Radio value={false}>不缓冲</Radio>
                    </Radio.Group>
                </Form.Item>}
                {menuType !== 'F' && <Form.Item name="status" label="菜单状态">
                    <Radio.Group>
                        <Radio value={1}>正常</Radio>
                        <Radio value={0}>停用</Radio>
                    </Radio.Group>
                </Form.Item>}
            </Form>
        </Spin>
    </Drawer>
}

export default MenuUpdateDrawer;