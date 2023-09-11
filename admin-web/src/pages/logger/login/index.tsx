import {useEffect, useState} from "react";
import {useRequest} from "ahooks";
import {ColumnsType} from "antd/es/table";
import {loginLoggerPage} from "@/api/logger.ts";
import {Button, DatePicker, Form, Input, Select, Space, Table, Tag} from "antd";
import {LoginLoggerQueryProp, LoginLoggerTableProp} from "@/pages/logger/modules.ts";

const LoggerLoginPage = () => {

    const columns: ColumnsType<LoginLoggerTableProp> = [
        {
            title: '用户名称',
            key: 'userName',
            dataIndex: 'userName',
            align: 'center'
        },
        {
            title: '登录地址',
            key: 'ip',
            dataIndex: 'ip',
            align: 'center'
        },
        {
            title: '登录地点',
            key: 'address',
            dataIndex: 'address',
            align: 'center',
        },
        {
            title: '浏览器',
            key: 'browser',
            dataIndex: 'browser',
            align: 'center'
        },
        {
            title: '操作系统',
            key: 'os',
            dataIndex: 'os',
            align: 'center'
        },
        {
            title: '状态',
            key: 'status',
            dataIndex: 'status',
            align: 'center',
            width: 160,
            render: (_, record) => record.status === 1 ? <Tag color="success">成功</Tag> : <Tag color="error">失败</Tag>
        },
        {
            title: '登录信息',
            key: 'msg',
            dataIndex: 'msg',
            align: 'center'
        },
        {
            title: '登录时间',
            key: 'loginTime',
            dataIndex: 'loginTime',
            align: 'center',
            width: 160,
        }
    ]

    const [form] = Form.useForm<LoginLoggerQueryProp>();
    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<LoginLoggerTableProp[]>([]);

    const {loading, run} = useRequest(loginLoggerPage, {
        manual: true,
        onSuccess: (data)=> {
            form.setFieldsValue({page: data.page, size: data.size})
            setDatasource(data.records);
            setTotal(data.total)
        }
    })

    const resetForm = () => {
        form.setFieldsValue({status: -1, address: undefined, endTime: undefined, startTime: undefined})
    }

    useEffect(() => {
        run({...form.getFieldsValue(), page: 1, size: 10})
    }, []);

    return <>
        <Space>
            <Form
                form={form}
                layout={'inline'}
                style={{ maxWidth: 'none' }}
            >
                <Form.Item name="address" label="登录地址">
                    <Input placeholder="请输入登录地址" />
                </Form.Item>
                <Form.Item name="userName" label="用户名称">
                    <Input placeholder="请输入用户名称" />
                </Form.Item>
                <Form.Item name="status" label="登录状态" initialValue={-1}>
                    <Select style={{width: 100}}>
                        <Select.Option value={-1}>全部</Select.Option>
                        <Select.Option value={1}>成功</Select.Option>
                        <Select.Option value={0}>失败</Select.Option>
                    </Select>
                </Form.Item>
                <Form.Item label="登录时间">
                    <DatePicker.RangePicker />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" onClick={() => run(form.getFieldsValue())}>搜索</Button>
                    <Button style={{marginLeft: 10}} onClick={resetForm}>重置</Button>
                </Form.Item>
            </Form>
        </Space>
        <Table
            bordered
            size={'small'}
            columns={columns}
            loading={loading}
            dataSource={datasource}
            style={{ marginTop: 10 }}
            rowKey={(record) => record.id}
            pagination={{
                onShowSizeChange: (current, size) => run({...form.getFieldsValue(), page: current, size: size}),
                onChange: (page, pageSize) => run({...form.getFieldsValue(), page: page, size: pageSize}),
                showTotal: () => `共 ${total} 个`,
                showQuickJumper: true,
                showSizeChanger: true,
                pageSize: form.getFieldValue("size") as number,
                current: form.getFieldValue("page") as number,
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

export default LoggerLoginPage;