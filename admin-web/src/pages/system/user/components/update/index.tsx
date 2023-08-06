import React, {useEffect, useState} from "react";
import {UserUpdateDrawerProp, UserUpdateFormProp, validateEmail, validateMobile} from "@/pages/system/user/modules.ts";
import {useRequest} from "ahooks";
import {deptTree} from "@/api/dept.ts";
import {userInfo, userUpdate} from "@/api/user.ts";
import {Button, Drawer, Form, Input, message, Radio, Select, Space, Spin, TreeSelect} from "antd";
import {DeptTreeProp} from "@/pages/system/dept/modules.ts";
import {PostInfoProp} from "@/pages/system/post/modules.ts";
import {RoleInfoProp} from "@/pages/system/role/modules.ts";
import {postList} from "@/api/post.ts";
import {roleList} from "@/api/role.ts";


const UserUpdateDrawer: React.FC<UserUpdateDrawerProp> = ({visible, userId, close}) => {


    const [form] = Form.useForm<UserUpdateFormProp>();
    const [depts, setDepts] = useState<DeptTreeProp[]>([]);
    const [posts, setPosts] = useState<PostInfoProp[]>([]);
    const [roles, setRoles] = useState<RoleInfoProp[]>([]);
    const loadDepts = useRequest(deptTree, {manual: true, onSuccess: (data) => setDepts(data)})
    const loadPosts = useRequest(postList, {manual: true, onSuccess: (data) => setPosts(data)});
    const loadRoles = useRequest(roleList, {manual: true, onSuccess: (data) => setRoles(data)});

    const loadUserInfo = useRequest(
        userInfo,
        {
            manual: true,
            onSuccess: (data) => {
                form.setFieldsValue({...data})
            }
        });

    const updateHandler = useRequest(userUpdate, {
        manual: true,
        onSuccess: (_)=> {
            message.success('修改用户信息成功');
            close(true);
        }
    })

    // 提交表单
    const submitForm = () => form.validateFields().then(value => updateHandler.run({...value}));

    useEffect(() => {
        if (visible && userId) {
            loadDepts.run();
            loadPosts.run();
            loadRoles.run();
            loadUserInfo.run(userId)
        }
        return () => {
            form.resetFields()
        }
    }, [visible])


    return <Drawer
        width={500}
        title="修改角色"
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
        <Spin tip="加载中......" spinning={updateHandler.loading}>
            <Form labelCol={{ span: 5 }}
                  wrapperCol={{ span: 19 }}
                  form={form}
                  layout="horizontal"
                  name="form_in_modal"
            >
                <Form.Item name="userId" label="主键" style={{display: 'none'}}>
                    <Input />
                </Form.Item>
                <Form.Item name="userName" label="用户名称" rules={[{ required: true, message: '请输入用户名称' }]}>
                    <Input placeholder="请输入用户名称" />
                </Form.Item>
                <Form.Item name="nickName" label="用户昵称" rules={[{ required: true, message: '请输入用户昵称' }]}>
                    <Input placeholder="请输入用户昵称" />
                </Form.Item>
                <Form.Item name="deptId" label="所属部门" rules={[{ required: true, message: '请选择所属部门' }]}>
                    <TreeSelect placeholder="请选择所属部门" treeData={depts} fieldNames={{label: 'title', value: 'key', children: 'children'}}/>
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
                <Form.Item name="phone" label="手机号码" rules={[{validator: validateMobile, required: false, message: '请输入正确的手机号码'}]}>
                    <Input placeholder="请输入手机号码" />
                </Form.Item>
                <Form.Item name="email" label="邮箱" rules={[{validator: validateEmail, required: false, message: '请输入正确的邮箱格式'}]}>
                    <Input placeholder="请输入邮箱" />
                </Form.Item>
                <Form.Item name="sex" label="性别">
                    <Select
                        options={[
                            { value: 0, label: '女' },
                            { value: 1, label: '男' },
                            { value: 2, label: '未知' },
                        ]}
                    />
                </Form.Item>
                <Form.Item name="status" label="状态">
                    <Radio.Group>
                        <Radio value={1}>正常</Radio>
                        <Radio value={0}>停用</Radio>
                    </Radio.Group>
                </Form.Item>
                <Form.Item name="roleId" label="选择角色">
                    <Select
                        mode="multiple"
                        allowClear
                        style={{ width: '100%' }}
                        placeholder="请选择对应的角色"
                        options={roles}
                    />
                </Form.Item>
                <Form.Item name="remark" label="备注">
                    <Input.TextArea rows={3} placeholder="请输入备注" />
                </Form.Item>
            </Form>
        </Spin>

    </Drawer>
}

export default UserUpdateDrawer;