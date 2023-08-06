import React, {useEffect, useState} from "react";
import {Button, Drawer, Form, Input, InputNumber, message, Space, Spin, Radio, Checkbox, Tree} from "antd";
import {RoleCreateDrawerProp, RoleCreateFormProp} from "@/pages/system/role/modules.ts";
import {useRequest} from "ahooks";
import {roleCreate} from "@/api/role.ts";
import {menuTree} from "@/api/menu.ts";
import {MenuTreeProp} from "@/pages/system/menu/modules.ts";
import {CheckboxValueType} from "antd/es/checkbox/Group";

// GetTreeKeys 递归获取Tree的key
function GetTreeKeys(data: MenuTreeProp[]): number[] {
    const list = []
    for (const item of data) {
        list.push(item.key)
        if (item.children && item.children.length > 0) {
            list.push(...GetTreeKeys(item.children))
        }
    }
    return list;
}

// RoleCreateDrawer 角色创建
const RoleCreateDrawer: React.FC<RoleCreateDrawerProp> = ({visible, close}) => {

    const [selectKeys, setSelectKeys] = useState<number[]>([]);
    const [expendKeys, setExpendKeys] = useState<number[]>([]);
    const [linkedFlag, setLinkedFlag] = useState<boolean>(false);
    const [form] = Form.useForm<RoleCreateFormProp>();
    const [tree, setTree] = useState<MenuTreeProp[]>([])
    const [treeSetting, setTreeSetting] = useState<number[]>([]);
    const {loading, run} = useRequest(roleCreate, {
        manual: true,
        onSuccess: (_)=> {
            message.success('创建角色成功');
            close(true);
        }
    })

    const loadTree = useRequest(menuTree, {manual: true, onSuccess: (data) => setTree(data)})

    // 提交表单
    const submitForm = () => form.validateFields().then(value => run({...value, menuIds: selectKeys}));

    const changeTree = (values: CheckboxValueType[]) => {
        const ids = GetTreeKeys(tree);
        const data = values as number[];
        setExpendKeys(data.includes(0) ? ids : [])
        setLinkedFlag(!data.includes(1))
        setSelectKeys(data.includes(2) ? ids : [])
        setTreeSetting(values as number[]);
    }

    useEffect(() => {
        if (visible) {
            loadTree.run()
            form.setFieldsValue({...form.getFieldsValue(), roleSort: 1, status: 1})
        }
    }, [visible])

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
                    <InputNumber min={1} placeholder="请输入角色顺序" style={{width: '100%'}} />
                </Form.Item>
                <Form.Item name="status" label="状态">
                    <Radio.Group>
                        <Radio value={0}>停用</Radio>
                        <Radio value={1}>正常</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item label="菜单权限">
                    <Checkbox.Group style={{marginTop:6, marginBottom: 10}} onChange={changeTree} value={treeSetting}>
                        <Checkbox value={0}>展开/折叠</Checkbox>
                        <Checkbox value={2}>全选/全不选</Checkbox>
                        <Checkbox value={1}>父子联动</Checkbox>
                    </Checkbox.Group>
                    <Tree
                        checkable
                        height={400}
                        treeData={tree}
                        checkedKeys={selectKeys}
                        expandedKeys={expendKeys}
                        checkStrictly={linkedFlag}
                        style={{border: '1px solid #d5d5d5', borderRadius: 5, padding: '8px 5px'}}
                        onCheck={(e) => setSelectKeys(e as React.Key[] as number[])}
                        onExpand={(e) => setExpendKeys(e as number[])}
                    />
                </Form.Item>
                <Form.Item name="remark" label="备注">
                    <Input.TextArea placeholder="请输入备注" />
                </Form.Item>
            </Form>
        </Spin>

    </Drawer>
}

export default RoleCreateDrawer;