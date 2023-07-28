import {Button, message, Modal, Space, Table, Tag} from "antd";
import React, {useEffect, useState} from "react";
import {
    DrawerProp,
    MenuTableTreeProp,
    MenuTableTreeQueryProp,
} from "@/pages/system/menu/modules";
import {ColumnsType} from "antd/es/table";
import {menuDelete, menuTable} from "@/api/menu.ts";
import {useRequest} from "ahooks";
import {MenuCreateDrawer, MenuUpdateDrawer} from "@/pages/system/menu/components";
import IconFont from "@/components/IconFont";
import {ExclamationCircleFilled} from "@ant-design/icons";
import {roleDelete} from "@/api/role.ts";


const AuthorityPermissionPage = () => {

    const columns: ColumnsType<MenuTableTreeProp> = [
        {
            title: '菜单名称',
            key: 'title',
            dataIndex: 'title',
            align: 'center',
            width: 150
        },
        {
            title: '图标',
            key: 'icon',
            dataIndex: 'icon',
            align: 'center',
            width: 80,
            render: (_, record) => record.icon ? <IconFont type={record.icon} /> : null
        },
        {
            title: '排序',
            key: 'order',
            dataIndex: 'order',
            align: 'center',
            width: 80
        },
        {
            title: '权限标识',
            key: 'code',
            dataIndex: 'code',
            align: 'center',
            width: 150,
            render: (text) => text ? <Tag color="red">{text}</Tag> : null
        },
        {
            title: '组件路径',
            key: 'path',
            dataIndex: 'path',
            align: 'center',
            width: 150
        },
        {
            title: '状态',
            key: 'status',
            dataIndex: 'status',
            align: 'center',
            width: 150,
            render: (text) => text === 1 ? <Tag color="blue">正常</Tag>: <Tag color="red">禁用</Tag>
        },
        {
            title: '创建时间',
            key: 'createTime',
            dataIndex: 'createTime',
            align: 'center',
            width: 150
        },
        {
            title: '操作',
            key: 'active',
            align: 'center',
            width: 200,
            render: (_, record) => (
                <Space size={'small'}>
                    <Button size="small" type="primary" onClick={() => openDrawer('edit', record.key)}>编辑</Button>
                    <Button size="small" type="primary" onClick={() => openDrawer('add', record.key)}>增加</Button>
                    <Button size="small" type="primary" danger onClick={() => deleteMenu(record.key)}>删除</Button>
                </Space>
            ),
        }
    ];

    const [tableQuery] = useState<MenuTableTreeQueryProp>({});
    const [drawerProp, setDrawerProp] = useState<DrawerProp>({createVisible: false, updateVisible: false});
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<MenuTableTreeProp[]>([]);

    const {run, loading} = useRequest(menuTable, {
        manual: true,
        onSuccess: (data) => {
            setDatasource(data);
        }
    })

    const openDrawer = (types: 'add' | 'edit', id: number) => {
        if (types === 'add') {
            setDrawerProp({...drawerProp, parentId: id, createVisible: true});
        } else if (types === 'edit') {
            setDrawerProp({...drawerProp, updateVisible: true, currId: id});
        }
    }

    const deleteMenu = (menuId: number) => {
        Modal.confirm({
            title: '警告',
            icon: <ExclamationCircleFilled />,
            content: '确认删除当前菜单？',
            okText: '删除',
            okType: 'danger',
            cancelText: '取消',
            onOk: async () => {
                const msg = await menuDelete(menuId);
                message.success(msg)
            },
            onCancel: () => {}
        });
    }


    const closeDrawer = (isLoad: boolean) => {
        setDrawerProp({...drawerProp,
                                updateVisible: false,
                                createVisible: false,
                                currId: undefined,
                                parentId: undefined
                              })
        if (isLoad) {
            run(tableQuery);
        }
    }

    useEffect(() => {
        run(tableQuery)
    }, [])


    return <>
        <Space>
            <Button type='primary' onClick={() => setDrawerProp({...drawerProp, createVisible: true, parentId: 0})}>创建</Button>
            <Button type='primary' onClick={() => run(tableQuery)}>刷新</Button>
        </Space>
        <Table
            bordered
            size={'small'}
            loading={loading}
            columns={columns}
            dataSource={datasource}
            style={{ marginTop: 10 }}
            rowKey={(record) => record.key}
            pagination={false}
            rowSelection={{
                type: 'checkbox',
                selectedRowKeys: selectedRowKeys,
                onChange: (selectedRowKeys: React.Key[]) => {
                    setSelectedRowKeys([...selectedRowKeys.map(item => item as number)])
                }
            }}
        />
        <MenuCreateDrawer
            parentId={drawerProp.parentId!}
            visible={drawerProp.createVisible}
            close={closeDrawer}
        />
        <MenuUpdateDrawer
            menuId={drawerProp.currId!}
            visible={drawerProp.updateVisible}
            close={closeDrawer} />
    </>
}

export default AuthorityPermissionPage;