import {Button, Checkbox, Drawer, Form, Input, Select, Space, Spin, Tree} from "antd";
import React, {useEffect, useState} from "react";
import {Role} from "@/types";
import {CheckboxValueType} from "antd/es/checkbox/Group";
import {useRequest} from "ahooks";
import {deptTree} from "@/api/dept.ts";
import {DeptTreeProp} from "@/pages/system/dept/modules.ts";

// GetTreeKeys 递归获取Tree的key
function GetTreeKeys(data: DeptTreeProp[]): number[] {
    const list = []
    for (const item of data) {
        list.push(item.key)
        if (item.children && item.children.length > 0) {
            list.push(...GetTreeKeys(item.children))
        }
    }
    return list;
}

const DataPermDrawer: React.FC<Role.RoleDataPermProp> = ({roleProp,visible, close}) => {

    const [form] = Form.useForm<Role.DataPermFormProp>();
    const [selectKeys, setSelectKeys] = useState<number[]>([]);
    const [expendKeys, setExpendKeys] = useState<number[]>([]);
    const [linkedFlag, setLinkedFlag] = useState<boolean>(false);
    const [loading, setLoading] = useState<boolean>(false);
    const [treeSetting, setTreeSetting] = useState<number[]>([0, 1]);
    const [tree, setTree] = useState<DeptTreeProp[]>([]);


    const loadTree = useRequest(deptTree, {manual: true, onSuccess: (data) => setTree(data)})

    const submitForm = () => {
        form.validateFields().then(value => {
            console.log("value", value);
        })
    }

    const changeTree = (values: CheckboxValueType[]) => {
        const ids = GetTreeKeys(tree);
        const data = values as number[];
        setExpendKeys(data.includes(0) ? ids : [])
        setLinkedFlag(!data.includes(1))
        setSelectKeys(data.includes(2) ? ids : [])
        setTreeSetting(values as number[]);
    }

    const onCheck = (value: { checked: React.Key[]; halfChecked: React.Key[]; } | React.Key[]) => {
        let keys = 'checked' in value ? (value.checked as number[]) : (value as React.Key[] as number[]);
        setSelectKeys(keys);
        form.setFieldValue('scopeValue', keys);
    };

    const changeDataScope = (e: string) => {
        form.setFieldValue("scopeType", e);
        if (e === "2") {
            loadTree.run();
            const ids = GetTreeKeys(tree);
            setExpendKeys(ids)
        }
    }

    useEffect(() => {
        if (visible) {
            form.setFieldsValue({roleId: roleProp.roleId, roleName: roleProp.roleName, roleCode: roleProp.roleKey, scopeType: '1'})
        }
        return () => {
            form.resetFields();
        }
    }, [visible]);




    return <Drawer
        width={500}
        title="数据权限"
        placement="right"
        maskClosable={false}
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
                <Form.Item name="roleId" style={{display: 'none'}}><Input /></Form.Item>
                <Form.Item name="roleName" label="角色名称" rules={[{ required: true, message: '请输入角色名称' }]}>
                    <Input placeholder="请输入角色名称" readOnly />
                </Form.Item>
                <Form.Item name="roleCode" label="角色编码" rules={[{ required: true, message: '请输入角色编码' }]}>
                    <Input placeholder="请输入角色编码" readOnly />
                </Form.Item>
                <Form.Item name="scopeType" label="数据范围" rules={[{ required: true, message: '请选择数据范围' }]}>
                    <Select
                        style={{ width: '100%' }}
                        onChange={e => changeDataScope(e)}
                        options={[
                            { value: "1", label: "全部数据权限" },
                            { value: "2", label: "自定数据权限" },
                            { value: "3", label: "本部门数据权限" },
                            { value: "4", label: "本部门及以下数据权限" },
                            { value: "5", label: "仅本人数据权限" }
                        ]}
                    />
                </Form.Item>
                {form.getFieldValue('scopeType') === '2' && <Form.Item label="数据权限" name="scopeValue" rules={[
                    {required: false},
                    () => ({
                        validator(_, value: number[]) {
                            if (form.getFieldValue("scopeType") === 2 && value.length <= 0) {
                                return Promise.reject(new Error('请选择自定义的数据范围！'));
                            }
                            return Promise.resolve();
                        }
                    })]}>
                    <Checkbox.Group style={{marginTop:6, marginBottom: 10}} onChange={changeTree} value={treeSetting}>
                        <Checkbox value={0}>展开/折叠</Checkbox>
                        <Checkbox value={2}>全选/全不选</Checkbox>
                        <Checkbox value={1}>父子联动</Checkbox>
                    </Checkbox.Group>
                    <Spin tip="加载中......" spinning={loadTree.loading}>
                        <Tree
                            checkable
                            height={400}
                            treeData={tree}
                            checkedKeys={selectKeys}
                            expandedKeys={expendKeys}
                            checkStrictly={treeSetting.includes(1)}
                            style={{border: '1px solid #d5d5d5', borderRadius: 5, padding: '8px 5px'}}
                            onCheck={onCheck}
                            onExpand={(e) => setExpendKeys(e as number[])}
                        />
                    </Spin>
                </Form.Item>}
            </Form>
        </Spin>
    </Drawer>
}

export default DataPermDrawer;
