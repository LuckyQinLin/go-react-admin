import {Button, Col, Divider, Input, message, Modal, Row, Select, Space, Switch, Table, Tree} from "antd";
import {useRequest} from "ahooks";
import {deptTree} from "@/api/dept.ts";
import {useEffect, useMemo, useState} from "react";
import {DeptTreeProp} from "@/pages/system/dept/modules.ts";
import {UserDrawerProp, UserPageQueryProp, UserTableProp} from "@/pages/system/user/modules.ts";
import {userDelete, userPage, userStatus} from "@/api/user.ts";
import {ColumnsType} from "antd/es/table";
import {UserCreateDrawer, UserRoleDrawer, UserUpdateDrawer} from "@/pages/system/user/components";
import {ExclamationCircleFilled} from "@ant-design/icons";
import styled from "@emotion/styled";

const SystemUserPage = () => {

    const columns: ColumnsType<UserTableProp> = [
        {
            title: '用户名称',
            key: 'userName',
            dataIndex: 'userName',
            align: 'center'
        },
        {
            title: '用户昵称',
            key: 'nickName',
            dataIndex: 'nickName',
            align: 'center'
        },
        {
            title: '部门',
            key: 'deptName',
            dataIndex: 'deptName',
            align: 'center',
            width: 160,
        },
        {
            title: '手机',
            key: 'phone',
            dataIndex: 'phone',
            align: 'center'
        },
        {
            title: '状态',
            key: 'status',
            dataIndex: 'status',
            align: 'center',
            width: 160,
            render: (_, record) => <Switch
                checkedChildren="正常"
                unCheckedChildren="停用"
                checked={record.status === 1}
                onChange={e => changeUserStatus(record.userId, e)}
            />
        },
        {
            title: '创建时间',
            key: 'createTime',
            dataIndex: 'createTime',
            align: 'center',
            width: 170,
        },
        {
            title: '操作',
            key: 'active',
            align: 'center',
            width: 160,
            render: (_, record) => (
                record.isSuper ? null : <Space size={'small'}>
                    <Button type="link" size='small' style={{padding: 4}} onClick={() => openDrawer('update', record.userId)}>修改</Button>
                    <Button type="link" size='small' style={{padding: 4}}>重置密码</Button>
                    <Button type="link" size='small' style={{padding: 4}} onClick={() => openDrawer('role', record.userId)}>分配角色</Button>
                    <Button type="link" size='small' style={{padding: 4}} danger onClick={() => deleteRoleHandler(record.userId)}>删除</Button>
                </Space>
            ),
        },
    ]


    const [searchValue, setSearchValue] = useState('');
    const [tree, setTree] = useState<DeptTreeProp[]>([]);
    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<UserTableProp[]>([]);
    const [pageQuery, setPageQuery] = useState<UserPageQueryProp>({page: 1, size: 10});
    const [userDrawer, setUserDrawer] = useState<UserDrawerProp>({createVisible: false, updateVisible: false, roleVisible: false});

    const loadTree = useRequest(deptTree, {
        manual: true,
        onSuccess: (data) => setTree(data)
    });

    const loadUserStatus = useRequest(userStatus, {
        manual: true,
        onSuccess: (data) => {
            message.success(data)
            loadUser.run(pageQuery)
        }
    });

    // changeUserStatus 修改用户状态
    const changeUserStatus = (userId: number, status: boolean) => {
        loadUserStatus.run(userId, status)
    }

    const loadUser = useRequest(userPage, {
        manual: true,
        onSuccess: ({records, total}) => {
            setTotal(total);
            setDatasource(records);
        }
    });

    const openDrawer = (types: 'create' | 'update' | 'role', userId?: number) => {
        switch (types) {
            case 'create':
                setUserDrawer({createVisible: true, updateVisible: false, roleVisible: false});
                break;
            case 'update':
                setUserDrawer({createVisible: false, updateVisible: true, roleVisible: false, userId: userId});
                break;
            case 'role':
                setUserDrawer({createVisible: false, updateVisible: false, roleVisible: true, userId: userId});
                break;
            default:
                break
        }
    }

    const closeDrawer = (types: 'create' | 'update' | 'role', isLoad: boolean) => {
        switch (types) {
            case 'create':
            case 'update':
            case 'role':
                setUserDrawer({createVisible: false, updateVisible: false, roleVisible: false, userId: undefined});
                break;
            default:
                break
        }
        if (isLoad) {
            loadUser.run(pageQuery)
        }
    }

    const deleteRoleHandler = async (id?: number) => {
        if (id === undefined && selectedRowKeys.length <= 0 ) {
            message.error('请先选择数据在进行删除')
            return
        }
        Modal.confirm({
            title: '警告',
            icon: <ExclamationCircleFilled />,
            content: '确认删除选中的用户？',
            okText: '删除',
            okType: 'danger',
            cancelText: '取消',
            onOk: async () => {
                await userDelete(id ? [id] : selectedRowKeys);
                loadUser.run(pageQuery)
                message.success('删除成功')
                setSelectedRowKeys([]);
            },
            onCancel: () => {}
        });
    }

    const getParentKey = (key: React.Key, tree: DeptTreeProp[]): React.Key => {
        let parentKey: React.Key;
        for (let i = 0; i < tree.length; i++) {
            const node = tree[i];
            if (node.children) {
                if (node.children.some((item) => item.key === key)) {
                    parentKey = node.key;
                } else if (getParentKey(key, node.children)) {
                    parentKey = getParentKey(key, node.children);
                }
            }
        }
        return parentKey!;
    };

    const treeData = useMemo(() => {
        const loop = (data: DeptTreeProp[]): DeptTreeProp[] =>
            data.map((item) => {
                const strTitle = item.title as string;
                const index = strTitle.indexOf(searchValue);
                const beforeStr = strTitle.substring(0, index);
                const afterStr = strTitle.slice(index + searchValue.length);
                const title = index > -1 ? (
                        <span>
                            {beforeStr}
                            <span style={{color: "#f50"}}>{searchValue}</span>
                            {afterStr}
                        </span>
                    ) : (
                        <span>{strTitle}</span>
                    );
                if (item.children) {
                    return { title, key: item.key, children: loop(item.children) } as DeptTreeProp;
                }
                return {
                    title,
                    key: item.key,
                } as DeptTreeProp;
            });
        return loop(tree);
    }, [searchValue]);


    useEffect(() => {
        loadUser.run(pageQuery);
    }, [pageQuery])

    useEffect(() => {
        loadTree.run();
        loadUser.run(pageQuery);
    }, []);

    return <Container>
            <Row gutter={[16, 16]} style={{display: "flex", flexDirection: "row", flexFlow: "nowrap"}}>
                <Col flex="250px">
                <Input
                    allowClear
                    placeholder="输入部门名称搜索"
                    value={searchValue}
                    style={{marginBottom: 10}}
                    onChange={e => setSearchValue(e.target.value)}
                />
                {
                    tree.length > 0 && <Tree
                        defaultExpandAll={true}
                        defaultExpandParent={true}
                        treeData={searchValue === "" ? tree : treeData}
                        onSelect={selectedRowKeys => setPageQuery({...pageQuery, deptId: selectedRowKeys[0] as number})}
                    />
                }
            </Col>
                <Col flex="auto">
                <Space>
                    <Button type="primary" onClick={() => openDrawer('create')}>增加</Button>
                    <Button type="primary" danger onClick={() => deleteRoleHandler()}>删除</Button>
                    <Button type="primary">导入</Button>
                    <Divider type="vertical" />
                    <Select
                        onChange={e => setPageQuery({...pageQuery, status: e})}
                        placeholder="选择用户状态"
                        defaultValue={-1}
                        style={{ width: 180 }}
                        allowClear
                        options={[
                            { value: -1, label: '全部' },
                            { value: 1, label: '正常' },
                            { value: 0, label: '停用' }
                        ]}
                    />
                    <Input placeholder="输入手机号/邮箱搜索" onChange={e => setPageQuery({...pageQuery, name: e.target.value})} />
                </Space>
                <Table
                    bordered
                    size={'small'}
                    columns={columns}
                    loading={loadUser.loading}
                    dataSource={datasource}
                    style={{ marginTop: 10 }}
                    rowKey={(record) => record.userId}
                    pagination={{
                        onShowSizeChange: (current, size) => loadUser.run({...pageQuery, page: current, size: size}),
                        onChange: (page, pageSize) => loadUser.run({...pageQuery, page: page, size: pageSize}),
                        showTotal: () => `共 ${total} 个`,
                        showQuickJumper: true,
                        showSizeChanger: true,
                        pageSize: pageQuery.size,
                        current: pageQuery.page,
                        size: 'default',
                        total: total,
                    }}
                    rowSelection={{
                        type: 'checkbox',
                        selectedRowKeys: selectedRowKeys,
                        onChange: (selectedRowKeys: React.Key[]) => {
                            setSelectedRowKeys([...selectedRowKeys.map(item => item as number)])
                        }
                    }}
                />
            </Col>
        </Row>
        <UserCreateDrawer visible={userDrawer.createVisible} close={(isLoad) => closeDrawer('create', isLoad)} />
        <UserUpdateDrawer visible={userDrawer.updateVisible} close={(isLoad) => closeDrawer('update', isLoad)} userId={userDrawer.userId} />
        <UserRoleDrawer visible={userDrawer.roleVisible} close={() => closeDrawer('role', false)} userId={userDrawer.userId} />
    </Container>
}

const Container = styled.div`
    background-color: #ffffff;
    padding: 16px;
    border-radius: 5px;
`

export default SystemUserPage;
