import React, {useEffect, useState} from "react";
import {AllocateUserDrawerProp} from "@/pages/system/role/modules.ts";
import {Button, Drawer, message, Space, Spin, Switch, Table} from "antd";
import {ColumnsType} from "antd/es/table";
import {UserPageQueryProp, UserTableProp} from "@/pages/system/user/modules.ts";
import {useRequest} from "ahooks";
import {userPage} from "@/api/user.ts";
import {getRoleUser, roleAllocateUser} from "@/api/role.ts";


const AllocateUserDrawer: React.FC<AllocateUserDrawerProp> = ({visible, roleId, close}) => {

    const columns: ColumnsType<UserTableProp> = [
        {
            title: '用户名称',
            key: 'userName',
            dataIndex: 'userName',
            align: 'center',
            ellipsis: true,
        },
        {
            title: '用户昵称',
            key: 'nickName',
            dataIndex: 'nickName',
            align: 'center',
            ellipsis: true,
        },
        {
            title: '手机',
            key: 'phone',
            dataIndex: 'phone',
            align: 'center',
            width: 120,
        },
        {
            title: '状态',
            key: 'status',
            dataIndex: 'status',
            align: 'center',
            width: 100,
            render: (_, record) => <Switch
                checkedChildren="正常"
                unCheckedChildren="停用"
                checked={record.status === 1}
            />
        },
        {
            title: '创建时间',
            key: 'createTime',
            dataIndex: 'createTime',
            align: 'center',
            width: 170,
        }
    ];

    const [pageQuery, setPageQuery] = useState<UserPageQueryProp>({page: 1, size: 10});
    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<UserTableProp[]>([]);

    const loadUser = useRequest(userPage, {
        manual: true,
        onSuccess: ({records, total, page, size}) => {
            setTotal(total);
            setDatasource(records);
            setPageQuery({...pageQuery, page: page, size: size})
        }
    });

    const loadRoleUser = useRequest(getRoleUser, {
        manual: true,
        onSuccess: (data) => {
            setSelectedRowKeys(data)
        }
    });

    const allocateUserLoad = useRequest(roleAllocateUser, {
        manual: true,
        onSuccess: () => {
            message.success('分配完成');
            close()
        }
    })

    const submitForm = () => allocateUserLoad.run(roleId!, selectedRowKeys);

    useEffect(() => {
        if (visible && roleId) {
           loadUser.run(pageQuery)
            loadRoleUser.run(roleId)
        }
    }, [visible]);

    return <Drawer
        width={800}
        title="分配用户"
        placement="right"
        onClose={close}
        open={visible}
        extra={
            <Space>
                <Button type="primary" danger onClick={close}>取消</Button>
                <Button type="primary" onClick={submitForm}>保存</Button>
            </Space>
        }
    >
        <Spin tip="加载中......" spinning={loadUser.loading}>
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
        </Spin>
    </Drawer>
}

export default AllocateUserDrawer;