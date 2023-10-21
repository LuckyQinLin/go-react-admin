import {ColumnsType} from "antd/es/table";
import {Button, message, Modal, Space, Table, Tag} from "antd";
import {useEffect, useState} from "react";
import {DeleteOutlined, DownloadOutlined, ExclamationCircleFilled, PlusOutlined} from "@ant-design/icons";
import {useRequest} from "ahooks";
import {noticeDelete} from "@/api/notice.ts";
import {ParamCreateDrawer, ParamUpdateDrawer} from "@/pages/system/param/components";
import {ConfigDrawerProp, ConfigPageProp, ConfigPageQueryProp} from "@/pages/system/param/modules.ts";
import {configPage} from "@/api/config.ts";
import styled from "@emotion/styled";

const SystemParamPage = () => {


    const columns: ColumnsType<ConfigPageProp> = [
        {
            title: '参数名称',
            key: 'configName',
            dataIndex: 'configName',
            align: 'center',
            width: 200,
            ellipsis: true
        },
        {
            title: '参数键名',
            key: 'configKey',
            dataIndex: 'configKey',
            align: 'center',
            ellipsis: true
        },
        {
            title: '参数键值',
            key: 'configValue',
            dataIndex: 'configValue',
            align: 'center',
            ellipsis: true
        },
        {
            title: '状态',
            key: 'status',
            dataIndex: 'status',
            align: 'center',
            width: 80,
            render: (_, record) => <Tag color={record.configType === 1 ? 'green' : 'blue'}>{record.configType === 1 ? '是' : '否'}</Tag>
        },
        {
            title: '备注',
            key: 'remark',
            dataIndex: 'remark',
            align: 'center',
            width: 250,
            ellipsis: true
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
                    <Button type="link" size='small' style={{padding: 4}} onClick={() => openDrawer('update', record.configId)}>修改</Button>
                    <Button type="link" size='small' danger style={{padding: 4}} onClick={() => deleteDictHandler(record.configId)}>删除</Button>
                </Space>
            ),
        },
    ]

    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<ConfigPageProp[]>([]);
    const [pageQuery] = useState<ConfigPageQueryProp>({page: 1, size: 10, configType: -1});
    const [configDrawer, setConfigDrawer] = useState<ConfigDrawerProp>({createVisible: false, updateVisible: false});

    const openDrawer = (types: 'create' | 'update', configId?: number) => {
        switch (types) {
            case 'create':
                setConfigDrawer({createVisible: true, updateVisible: false});
                break;
            case 'update':
                setConfigDrawer({createVisible: false, updateVisible: true, configId: configId});
                break;
            default:
                break
        }
    }

    const closeDrawer = (types: 'create' | 'update', isLoad: boolean) => {
        switch (types) {
            case 'create':
            case 'update':
                setConfigDrawer({createVisible: false, updateVisible: false});
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
                await noticeDelete(id ? [id] : selectedRowKeys);
                run(pageQuery)
                message.success('删除成功')
                setSelectedRowKeys([]);
            },
            onCancel: () => {}
        });
    }

    const {loading, run} = useRequest(configPage, {
        manual: true,
        onSuccess: (data)=> {
            setDatasource(data.records);
            setTotal(data.total)
        }
    })

    useEffect(() => run(pageQuery), [])

    return <Container>
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
            rowKey={(record) => record.configId}
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
        <ParamCreateDrawer visible={configDrawer.createVisible} close={(isLoad) => closeDrawer('create', isLoad)} />
        <ParamUpdateDrawer visible={configDrawer.updateVisible} close={(isLoad) => closeDrawer('update', isLoad)} configId={configDrawer.configId} />
    </Container>
}

const Container = styled.div`
    background-color: #ffffff;
    padding: 16px;
    border-radius: 5px;
`

export default SystemParamPage;