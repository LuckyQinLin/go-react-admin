import {Button, Divider, Input, Select, Space, DatePicker, Table, Switch} from "antd";
import {DeleteOutlined, EditFilled, DownloadOutlined, PlusOutlined} from "@ant-design/icons";
import {ColumnsType} from "antd/es/table";
import {RoleDrawerProp, RolePageProp, RolePageQueryProp} from "@/pages/system/role/modules.ts";
import {useRequest} from "ahooks";
import {rolePage} from "@/api/role.ts";
import {useEffect, useState} from "react";
import {RoleCreateDrawer} from "@/pages/system/role/components";

const SystemRolePage = () => {

    const columns: ColumnsType<RolePageProp> = [
        {
            title: '角色名称',
            key: 'roleName',
            dataIndex: 'roleName',
            align: 'center'
        },
        {
            title: '权限字符',
            key: 'roleKey',
            dataIndex: 'roleKey',
            align: 'center'
        },
        {
            title: '显示顺序',
            key: 'roleSort',
            dataIndex: 'roleSort',
            align: 'center',
            width: 160,
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
                    <Button type="link" style={{padding: 4}}>数据权限</Button>
                    <Button type="link" style={{padding: 4}}>分配用户</Button>
                    <Button type="link" style={{padding: 4}}>修改</Button>
                    <Button type="link" danger style={{padding: 4}}>删除</Button>
                </Space>
            ),
        },
    ]

    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<RolePageProp[]>([]);
    const [pageQuery, setPageQuery] = useState<RolePageQueryProp>({page: 1, size: 10});
    const [roleDrawer, setRoleDrawer] = useState<RoleDrawerProp>({createVisible: false});

    const {loading, run} = useRequest(rolePage, {
        manual: true,
        onSuccess: (data)=> {
            setDatasource(data.records);
            setTotal(data.total)
        }
    })

    const openDrawer = (types: 'create' | 'update') => {
        switch (types) {
            case 'create':
                setRoleDrawer({createVisible: true});
                break;
            default:
                break
        }
    }

    const closeDrawer = (types: 'create' | 'update', isLoad: boolean) => {
        switch (types) {
            case 'create':
                setRoleDrawer({createVisible: false});
                break;
            default:
                break
        }
        if (isLoad) {
            run(pageQuery)
        }
    }

    useEffect(() => run(pageQuery), [])


    return <>
        <Space>
            <Button type="primary" icon={<PlusOutlined />} onClick={() => openDrawer('create')}>新增</Button>
            <Button type="primary" icon={<EditFilled />}>修改</Button>
            <Button type="primary" icon={<DownloadOutlined />}>导出</Button>
            <Button type="primary" danger icon={<DeleteOutlined />}>删除</Button>
            <Divider type="vertical" />
            <Input style={{ width: 220 }} placeholder="请输入角色名称或者权限字符" />
            <Select
                defaultValue={2}
                placeholder="请选择角色状态"
                style={{ width: 160 }}
                allowClear
                options={[{value: 2, label: '全部'},{ value: 1, label: '正常' }, {value: 0, label: '停用'}]}
            />
            <DatePicker.RangePicker showTime />
        </Space>
        <Table
            bordered
            size={'small'}
            columns={columns}
            loading={loading}
            dataSource={datasource}
            style={{ marginTop: 10 }}
            rowKey={(record) => record.roleId}
            pagination={{
                onShowSizeChange: (current, size) => run({...pageQuery, page: current, size: size}),
                onChange: (page, pageSize) => run({...pageQuery, page: page, size: pageSize}),
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
        <RoleCreateDrawer visible={roleDrawer.createVisible} close={(isLoad) => closeDrawer('create', isLoad)} />
    </>
}

export default SystemRolePage;