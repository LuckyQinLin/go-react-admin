import {ColumnsType} from "antd/es/table";
import {
    OperateLoggerQueryProp,
    OperateLoggerTableProp
} from "@/pages/logger/modules.ts";
import {Space, Switch, Table} from "antd";
import {useState} from "react";
import {useRequest} from "ahooks";
import {operateLoggerPage} from "@/api/logger.ts";

const LoggerOperatePage = () => {
    const columns: ColumnsType<OperateLoggerTableProp> = [
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
        }
    ]

    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<OperateLoggerTableProp[]>([]);
    const [pageQuery] = useState<OperateLoggerQueryProp>({page: 1, size: 10});

    const {loading, run} = useRequest(operateLoggerPage, {
        manual: true,
        onSuccess: (data)=> {
            setDatasource(data.records);
            setTotal(data.total)
        }
    })

    return <>
        <Space></Space>
        <Table
            bordered
            size={'small'}
            columns={columns}
            loading={loading}
            dataSource={datasource}
            style={{ marginTop: 10 }}
            rowKey={(record) => record.id}
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
    </>
}

export default LoggerOperatePage;