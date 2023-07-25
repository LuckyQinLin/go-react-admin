import {Button, Divider, Input, Select, Space, DatePicker, Table, Switch, Modal, message} from "antd";
import {DeleteOutlined, DownloadOutlined, PlusOutlined, ExclamationCircleFilled} from "@ant-design/icons";
import {ColumnsType} from "antd/es/table";
import {RoleDrawerProp, RolePageProp, RolePageQueryProp} from "@/pages/system/role/modules.ts";
import {useRequest} from "ahooks";
import {roleDelete, rolePage} from "@/api/role.ts";
import {useEffect, useState} from "react";
import {RoleCreateDrawer, RoleUpdateDrawer} from "@/pages/system/role/components";

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
                    <Button type="link" style={{padding: 4}} onClick={() => openDrawer('update', record.roleId)}>修改</Button>
                    <Button type="link" danger style={{padding: 4}} onClick={() => deleteRoleHandler(record.roleId)}>删除</Button>
                </Space>
            ),
        },
    ]

    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<RolePageProp[]>([]);
    const [pageQuery] = useState<RolePageQueryProp>({page: 1, size: 10});
    const [roleDrawer, setRoleDrawer] = useState<RoleDrawerProp>({createVisible: false, updateVisible: false});

    const {loading, run} = useRequest(rolePage, {
        manual: true,
        onSuccess: (data)=> {
            setDatasource(data.records);
            setTotal(data.total)
        }
    })

    const deleteRoleHandler = async (id?: number) => {
        Modal.confirm({
            title: '警告',
            icon: <ExclamationCircleFilled />,
            content: '确认删除当前角色？',
            okText: '删除',
            okType: 'danger',
            cancelText: '取消',
            onOk: async () => {
                await roleDelete(id ? [id] : selectedRowKeys);
                run(pageQuery)
                message.success('删除成功')
                setSelectedRowKeys([]);
            },
            onCancel: () => {}
        });
    }

    const openDrawer = (types: 'create' | 'update', roleId?: number) => {
        switch (types) {
            case 'create':
                setRoleDrawer({createVisible: true, updateVisible: false});
                break;
            case 'update':
                setRoleDrawer({createVisible: false, updateVisible: true, roleId: roleId});
                break;
            default:
                break
        }
    }

    const closeDrawer = (types: 'create' | 'update', isLoad: boolean) => {
        switch (types) {
            case 'create':
            case 'update':
                setRoleDrawer({createVisible: false, updateVisible: false});
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
            <Button type="primary" icon={<PlusOutlined />} disabled={selectedRowKeys.length > 0} onClick={() => openDrawer('create')}>新增</Button>
            <Button type="primary" icon={<DownloadOutlined />} disabled={selectedRowKeys.length <= 0}>导出</Button>
            <Button type="primary" disabled={selectedRowKeys.length <= 0} danger icon={<DeleteOutlined />} onClick={() => deleteRoleHandler()}>删除</Button>
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
        <RoleUpdateDrawer visible={roleDrawer.updateVisible} close={isLoad => closeDrawer('update', isLoad)} roleId={roleDrawer.roleId} />
    </>
}

export default SystemRolePage;