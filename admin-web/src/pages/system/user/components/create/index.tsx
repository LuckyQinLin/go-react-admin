import React, {useEffect, useState} from "react";
import {UserCreateDrawerProp, UserCreateFormProp} from "@/pages/system/user/modules.ts";
import {
    Button,
    Drawer,
    Form,
    Input,
    message,
    Radio,
    Select,
    Space,
    Spin,
    TreeSelect
} from "antd";
import {useRequest} from "ahooks";
import {DeptTreeProp} from "@/pages/system/dept/modules.ts";
import {userCreate} from "@/api/user.ts";
import {deptTree} from "@/api/dept.ts";
import {postList} from "@/api/post.ts";
import {PostInfoProp} from "@/pages/system/post/modules.ts";
import {RoleInfoProp} from "@/pages/system/role/modules.ts";
import {roleList} from "@/api/role.ts";

const UserCreateDrawer: React.FC<UserCreateDrawerProp> = ({visible, deptId, close}) => {

    const [form] = Form.useForm<UserCreateFormProp>();
    const [depts, setDepts] = useState<DeptTreeProp[]>([]);
    const [posts, setPosts] = useState<PostInfoProp[]>([]);
    const [roles, setRoles] = useState<RoleInfoProp[]>([]);
    const loadDepts = useRequest(deptTree, {manual: true, onSuccess: (data) => setDepts(data)})
    const loadPosts = useRequest(postList, {manual: true, onSuccess: (data) => setPosts(data)});
    const loadRoles = useRequest(roleList, {manual: true, onSuccess: (data) => setRoles(data)});

    // 提交表单
    const submitForm = () => form.validateFields().then(value => run({...value}));

    const {loading, run} = useRequest(userCreate, {
        manual: true,
        onSuccess: (_)=> {
            message.success('创建用户成功');
            close(true);
        }
    })

    useEffect(() => {
        if (visible) {
            loadDepts.run();
            loadPosts.run();
            loadRoles.run();
            form.setFieldsValue({deptId: deptId, sex: 2, status: 1})
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
                <Form.Item name="userName" label="用户名称" rules={[{ required: true, message: '请输入用户名称' }]}>
                    <Input placeholder="请输入用户名称" />
                </Form.Item>
                <Form.Item name="nickName" label="用户昵称" rules={[{ required: true, message: '请输入用户昵称' }]}>
                    <Input placeholder="请输入用户昵称" />
                </Form.Item>
                <Form.Item name="deptId" label="所属部门">
                    <TreeSelect treeData={depts}/>
                </Form.Item>
                <Form.Item name="postId" label="归属岗位">
                    <Select
                        mode="multiple"
                        allowClear
                        style={{ width: '100%' }}
                        placeholder="请选择对应的岗位"
                        options={posts}
                    />
                </Form.Item>
                <Form.Item name="phone" label="手机号码">
                    <Input placeholder="请输入手机号码" />
                </Form.Item>
                <Form.Item name="email" label="邮箱">
                    <Input placeholder="请输入邮箱" />
                </Form.Item>
                <Form.Item name="sex" label="性别">
                    <Radio.Group>
                        <Radio value={0}>男</Radio>
                        <Radio value={1}>女</Radio>
                        <Radio value={2}>未知</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="status" label="状态">
                    <Radio.Group>
                        <Radio value={0}>停用</Radio>
                        <Radio value={1}>正常</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="roleId" label="角色">
                    <Select
                        mode="multiple"
                        allowClear
                        style={{ width: '100%' }}
                        placeholder="请选择对应的角色"
                        options={roles}
                    />
                </Form.Item>
                <Form.Item name="remark" label="备注">
                    <Input placeholder="请输入备注" />
                </Form.Item>
            </Form>
        </Spin>

    </Drawer>
}

export default UserCreateDrawer;