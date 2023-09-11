import React, {useEffect, useState} from "react";
import {Button, Drawer, message, Space, Table} from "antd";
import {UserRoleDrawerProp} from "@/pages/system/user/modules.ts";
import {useRequest} from "ahooks";
import {roleByUserId, rolePage} from "@/api/role.ts";
import {ColumnsType} from "antd/es/table";
import {RolePageProp, RolePageQueryProp} from "@/pages/system/role/modules.ts";
import {userRole} from "@/api/user.ts";

const UserRoleDrawer: React.FC<UserRoleDrawerProp> = ({userId, visible, close}) => {

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
            title: '创建时间',
            key: 'createTime',
            dataIndex: 'createTime',
            align: 'center',
            width: 160,
        }
    ]

    const loadUserRole = useRequest(roleByUserId, {
        manual: true,
        onSuccess: (data)=> {
            setSelectedRowKeys(data)
        }
    })

    const loadRole = useRequest(rolePage, {
        manual: true,
        onSuccess: (data)=> {
            setDatasource(data.records);
            setTotal(data.total)
        }
    })

    const saveUserRole = useRequest(userRole, {
        manual: true,
        onSuccess: (_)=> {
            message.success('分配成功')
        }
    })

    const submitForm = () => saveUserRole.run(userId!, selectedRowKeys)

    const [total, setTotal] = useState<number>(0);
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<RolePageProp[]>([]);
    const [pageQuery] = useState<RolePageQueryProp>({page: 1, size: 10});

    useEffect(() => {
        if (visible && userId) {
            loadRole.run({...pageQuery, status: 1});
            loadUserRole.run(userId);
        }
        return () => {

        }
    }, [visible])

    return <Drawer
        width={550}
        title="分配角色"
        placement="right"
        onClose={close}
        open={visible}
        extra={
            <Space>
                <Button type="primary" danger onClick={close}>取消</Button>
                <Button type="primary" loading={saveUserRole.loading} onClick={submitForm}>保存</Button>
            </Space>
        }
    >
        <Table
            bordered
            size={'small'}
            columns={columns}
            loading={loadRole.loading}
            dataSource={datasource}
            style={{ marginTop: 10 }}
            rowKey={(record) => record.roleId}
            pagination={{
                onShowSizeChange: (current, size) => loadRole.run({...pageQuery, page: current, size: size, status: 1}),
                onChange: (page, pageSize) => loadRole.run({...pageQuery, page: page, size: pageSize, status: 1}),
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
    </Drawer>
}

export default UserRoleDrawer;