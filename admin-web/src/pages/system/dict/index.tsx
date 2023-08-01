import {useEffect, useState} from "react";
import {useRequest} from "ahooks";
import {Button, message, Modal, Space, Switch, Table} from "antd";
import {ColumnsType} from "antd/es/table";
import {DeleteOutlined, DownloadOutlined, ExclamationCircleFilled, PlusOutlined} from "@ant-design/icons";
import {DictDrawerProp, DictPageProp, DictPageQueryProp} from "@/pages/system/dict/modules.ts";
import {dictDelete, dictPage} from "@/api/dict.ts";
import {DictCreateDrawer, DictUpdateDrawer} from "@/pages/system/dict/components";

const SystemDictPage = () => {

    const columns: ColumnsType<DictPageProp> = [
        {
            title: '字典名称',
            key: 'dictName',
            dataIndex: 'dictName',
            align: 'center'
        },
        {
            title: '字典类型',
            key: 'dictType',
            dataIndex: 'dictType',
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
            title: '备注',
            key: 'remark',
            dataIndex: 'remark',
            align: 'center'
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
                    <Button type="link" style={{padding: 4}} onClick={() => openDrawer('update', record.dictId)}>修改</Button>
                    <Button type="link" danger style={{padding: 4}} onClick={() => deleteDictHandler(record.dictId)}>删除</Button>
                </Space>
            ),
        },
    ]

    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<DictPageProp[]>([]);
    const [pageQuery] = useState<DictPageQueryProp>({page: 1, size: 10});
    const [postDrawer, setDictDrawer] = useState<DictDrawerProp>({createVisible: false, updateVisible: false});

    const openDrawer = (types: 'create' | 'update', dictId?: number) => {
        switch (types) {
            case 'create':
                setDictDrawer({createVisible: true, updateVisible: false});
                break;
            case 'update':
                setDictDrawer({createVisible: false, updateVisible: true, dictId: dictId});
                break;
            default:
                break
        }
    }

    const closeDrawer = (types: 'create' | 'update', isLoad: boolean) => {
        switch (types) {
            case 'create':
            case 'update':
                setDictDrawer({createVisible: false, updateVisible: false});
                break;
            default:
                break
        }
        if (isLoad) {
            run(pageQuery)
        }
    }

    const deleteDictHandler = async (id?: number) => {
        Modal.confirm({
            title: '警告',
            icon: <ExclamationCircleFilled />,
            content: '确认删除岗位数据？',
            okText: '删除',
            okType: 'danger',
            cancelText: '取消',
            onOk: async () => {
                await dictDelete(id ? [id] : selectedRowKeys);
                run(pageQuery)
                message.success('删除成功')
                setSelectedRowKeys([]);
            },
            onCancel: () => {}
        });
    }

    const {loading, run} = useRequest(dictPage, {
        manual: true,
        onSuccess: (data)=> {
            setDatasource(data.records);
            setTotal(data.total)
        }
    })

    useEffect(() => run(pageQuery), [])

    return <>
        <Space>
            <Button type="primary" icon={<PlusOutlined />} disabled={selectedRowKeys.length > 0} onClick={() => openDrawer('create')}>新增</Button>
            <Button type="primary" icon={<DownloadOutlined />} disabled={selectedRowKeys.length <= 0}>导出</Button>
            <Button type="primary" disabled={selectedRowKeys.length <= 0} danger icon={<DeleteOutlined />} onClick={() => deleteDictHandler()}>删除</Button>
        </Space>
        <Table
            bordered
            size={'small'}
            columns={columns}
            loading={loading}
            dataSource={datasource}
            style={{ marginTop: 10 }}
            rowKey={(record) => record.dictId}
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
        <DictCreateDrawer visible={postDrawer.createVisible} close={(isLoad) => closeDrawer('create', isLoad)} />
        <DictUpdateDrawer visible={postDrawer.updateVisible} close={(isLoad) => closeDrawer('update', isLoad)} dictId={postDrawer.dictId} />
    </>
}

export default SystemDictPage;