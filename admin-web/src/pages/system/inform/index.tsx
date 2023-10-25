import {ColumnsType} from "antd/es/table";
import {NoticeDrawerProp, NoticePageProp, NoticePageQueryProp} from "@/pages/system/inform/modules.ts";
import {Button, message, Modal, Space, Switch, Table, Tag} from "antd";
import {useEffect, useState} from "react";
import {DeleteOutlined, DownloadOutlined, ExclamationCircleFilled, PlusOutlined} from "@ant-design/icons";
import {noticeDelete, noticePage} from "@/api/notice.ts";
import {useRequest} from "ahooks";
import {NoticeCreateDrawer, NoticeUpdateDrawer} from "@/pages/system/inform/components";
import styled from "@emotion/styled";

const SystemInformPage = () => {
    const columns: ColumnsType<NoticePageProp> = [
        {
            title: '公告标题',
            key: 'noticeTitle',
            dataIndex: 'noticeTitle',
            align: 'center'
        },
        {
            title: '公告类型',
            key: 'noticeType',
            dataIndex: 'noticeType',
            align: 'center',
            render: (_, record) => <Tag color={record.noticeType === 1 ? 'green' : 'blue'}>{record.noticeType === 1 ? '通知' : '公告'}</Tag>
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
            title: '创建者',
            key: 'createBy',
            dataIndex: 'createBy',
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
                    <Button type="link" size='small' style={{padding: 4}} onClick={() => openDrawer('update', record.noticeId)}>修改</Button>
                    <Button type="link" size='small' danger style={{padding: 4}} onClick={() => deleteDictHandler(record.noticeId)}>删除</Button>
                </Space>
            ),
        },
    ]

    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<NoticePageProp[]>([]);
    const [pageQuery] = useState<NoticePageQueryProp>({page: 1, size: 10});
    const [noticeDrawer, setNoticeDrawer] = useState<NoticeDrawerProp>({createVisible: false, updateVisible: false});

    const openDrawer = (types: 'create' | 'update', noticeId?: number) => {
        switch (types) {
            case 'create':
                setNoticeDrawer({createVisible: true, updateVisible: false});
                break;
            case 'update':
                setNoticeDrawer({createVisible: false, updateVisible: true, noticeId: noticeId});
                break;
            default:
                break
        }
    }

    const closeDrawer = (types: 'create' | 'update', isLoad: boolean) => {
        switch (types) {
            case 'create':
            case 'update':
                setNoticeDrawer({createVisible: false, updateVisible: false});
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

    const {loading, run} = useRequest(noticePage, {
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
            rowKey={(record) => record.noticeId}
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
        <NoticeCreateDrawer visible={noticeDrawer.createVisible} close={(isLoad) => closeDrawer('create', isLoad)} />
        <NoticeUpdateDrawer visible={noticeDrawer.updateVisible} close={(isLoad) => closeDrawer('update', isLoad)} noticeId={noticeDrawer.noticeId} />
    </Container>
}

const Container = styled.div`
    background-color: #ffffff;
    padding: 16px;
    border-radius: 5px;
`

export default SystemInformPage;