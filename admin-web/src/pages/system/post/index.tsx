import {useEffect, useState} from "react";
import {useRequest} from "ahooks";
import {postDelete, postPage} from "@/api/post.ts";
import {PostDrawerProp, PostPageProp, PostPageQueryProp} from "@/pages/system/post/modules.ts";
import {Button, message, Modal, Space, Switch, Table} from "antd";
import {ColumnsType} from "antd/es/table";
import {DeleteOutlined, DownloadOutlined, ExclamationCircleFilled, PlusOutlined} from "@ant-design/icons";
import {PostCreateDrawer, PostUpdateDrawer} from "@/pages/system/post/components";

const SystemPostPage = () => {

    const columns: ColumnsType<PostPageProp> = [
        {
            title: '岗位名称',
            key: 'postName',
            dataIndex: 'postName',
            align: 'center'
        },
        {
            title: '岗位编码',
            key: 'postCode',
            dataIndex: 'postCode',
            align: 'center'
        },
        {
            title: '显示顺序',
            key: 'postSort',
            dataIndex: 'postSort',
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
                    <Button type="link" style={{padding: 4}} onClick={() => openDrawer('update', record.postId)}>修改</Button>
                    <Button type="link" danger style={{padding: 4}} onClick={() => deleteRoleHandler(record.postId)}>删除</Button>
                </Space>
            ),
        },
    ]

    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<PostPageProp[]>([]);
    const [pageQuery] = useState<PostPageQueryProp>({page: 1, size: 10});
    const [postDrawer, setPostDrawer] = useState<PostDrawerProp>({createVisible: false, updateVisible: false});

    const openDrawer = (types: 'create' | 'update', postId?: number) => {
        switch (types) {
            case 'create':
                setPostDrawer({createVisible: true, updateVisible: false});
                break;
            case 'update':
                setPostDrawer({createVisible: false, updateVisible: true, postId: postId});
                break;
            default:
                break
        }
    }

    const closeDrawer = (types: 'create' | 'update', isLoad: boolean) => {
        switch (types) {
            case 'create':
            case 'update':
                setPostDrawer({createVisible: false, updateVisible: false});
                break;
            default:
                break
        }
        if (isLoad) {
            run(pageQuery)
        }
    }

    const deleteRoleHandler = async (id?: number) => {
        Modal.confirm({
            title: '警告',
            icon: <ExclamationCircleFilled />,
            content: '确认删除岗位数据？',
            okText: '删除',
            okType: 'danger',
            cancelText: '取消',
            onOk: async () => {
                await postDelete(id ? [id] : selectedRowKeys);
                run(pageQuery)
                message.success('删除成功')
                setSelectedRowKeys([]);
            },
            onCancel: () => {}
        });
    }

    const {loading, run} = useRequest(postPage, {
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
            <Button type="primary" disabled={selectedRowKeys.length <= 0} danger icon={<DeleteOutlined />} onClick={() => deleteRoleHandler()}>删除</Button>
        </Space>
        <Table
            bordered
            size={'small'}
            columns={columns}
            loading={loading}
            dataSource={datasource}
            style={{ marginTop: 10 }}
            rowKey={(record) => record.postId}
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
        <PostCreateDrawer visible={postDrawer.createVisible} close={(isLoad) => closeDrawer('create', isLoad)} />
        <PostUpdateDrawer visible={postDrawer.updateVisible} close={(isLoad) => closeDrawer('update', isLoad)} postId={postDrawer.postId} />
    </>
}

export default SystemPostPage;