import {Button, Col, Input, Row, Space, Switch, Table, Tree} from "antd";
import {useRequest} from "ahooks";
import {deptTree} from "@/api/dept.ts";
import {useEffect, useState} from "react";
import {DeptTreeProp} from "@/pages/system/dept/modules.ts";
import {UserPageQueryProp, UserTableProp} from "@/pages/system/user/modules.ts";
import {userPage} from "@/api/user.ts";
import {ColumnsType} from "antd/es/table";

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
            render: (_, record) => <Switch checkedChildren="正常" unCheckedChildren="停用" checked={record.status === 1} />
        },
        {
            title: '创建时间',
            key: 'createTime',
            dataIndex: 'createTime',
            align: 'center',
            width: 160,
        },
        {
            title: '操作',
            key: 'active',
            align: 'center',
            width: 160,
            render: (_, record) => (
                <Space size={'small'}>
                    <Button type="link" style={{padding: 4}}>修改</Button>
                    <Button type="link" style={{padding: 4}}>删除</Button>
                    <Button type="link" style={{padding: 4}}>重置密码</Button>
                    <Button type="link" style={{padding: 4}}>分配角色</Button>
                </Space>
            ),
        },
    ]


    const [searchValue, setSearchValue] = useState('');
    const [tree, setTree] = useState<DeptTreeProp[]>([]);
    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<UserTableProp[]>([]);
    const [pageQuery] = useState<UserPageQueryProp>({page: 1, size: 10});

    const loadTree = useRequest(deptTree, {
        manual: true,
        onSuccess: (data) => setTree(data)
    });
    const loadUser = useRequest(userPage, {
        manual: true,
        onSuccess: ({records, total}) => {
            setTotal(total);
            setDatasource(records);
        }
    });

    useEffect(() => {
        loadTree.run();
        loadUser.run(pageQuery);
    }, [])

    return <Row gutter={[16, 16]}>
        <Col flex="250px">
            <Input placeholder="输入部门名称搜索" value={searchValue} style={{marginBottom: 10}} />
            {
                tree.length > 0 && <Tree
                    defaultExpandAll={true}
                    defaultExpandParent
                    onExpand={onExpand}
                    treeData={tree}
                />
            }

        </Col>
        <Col flex="auto">
            <Space>
                <Button type="primary">增加</Button>
                <Button type="primary" danger>删除</Button>
                <Button type="primary">导入</Button>
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
}

export default SystemUserPage;