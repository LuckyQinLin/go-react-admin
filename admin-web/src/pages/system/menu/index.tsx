import {Button, Col, Row, Space, Table, Tag, Tree} from "antd";
import React, {useEffect, useState} from "react";
import {
    DrawerProp,
    MenuTableTreeProp,
    PermissionListProp,
    PermissionTreeProp,
    RoleProp,
    UserListProp
} from "@/pages/system/menu/modules";
import {ColumnsType} from "antd/es/table";
import Search from "antd/es/input/Search";
import {permissionList, permissionTree} from "@/api/menu.ts";
import {TreeProps} from "antd/es/tree";
import {Icon} from "@/components";
import {MenuCreateDrawer, MenuUpdateDrawer} from "@/pages/system/menu/components";

const loadTree = (list: PermissionTreeProp[], key: React.Key, children: PermissionTreeProp[]): PermissionTreeProp[] =>
    list.map((node) => {
        if (node.key === key) {
            if (node.children) {
                const keys = node.children.map(item => item.key);
                const temp = children.filter(item => !keys.includes(item.key));
                const sorts = [...children, ...temp].sort((a, b) => a.sortNo! - b.sortNo!);
                return { ...node, sorts };
            } else {
                return { ...node, children };
            }
        }
        if (node.children) {
            return {
                ...node,
                children: loadTree(node.children, key, children),
            };
        }
        return node;
    });

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
            render: (_, record) => record.icon ? <Icon icon={record.icon} /> : null
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
            render: (text) => <Tag color="red">{text}</Tag>
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
            render: (text) => text === 0 ? <Tag color="blue">正常</Tag>: <Tag color="red">禁用</Tag>
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
                    <Button size="small" type="primary" onClick={() => openDrawer('edit', record.id)}>编辑</Button>
                    <Button size="small" type="primary" onClick={() => openDrawer('info', record.id)}>增加</Button>
                    <Button size="small" type="primary" danger>删除</Button>
                </Space>
            ),
        }
    ];

    const [drawerProp, setDrawerProp] = useState<DrawerProp>({types: 1, parentId: 0, createVisible: false, updateVisible: false});
    const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([]);
    const [datasource, setDatasource] = useState<MenuTableTreeProp[]>([]);


    const loadList = async (pId: number) => {
        const list = await permissionList(pId);
        setDatasource(list);
    }

    const openDrawer = (types: 'info' | 'edit', id: number) => {
        if (types === 'info') {
            setDrawerProp({...drawerProp, currId: id, createVisible: false, updateVisible: false});
        } else if (types === 'edit') {
            setDrawerProp({...drawerProp, updateVisible: true, currId: id, createVisible: false});
        }

    }


    const closeDrawer = (isLoad: boolean, data?: PermissionListProp) => {
        setDrawerProp({...drawerProp, updateVisible: false, createVisible: false})
        if (isLoad) {
            loadList(drawerProp.parentId!);
        }
        if (data) {
            // TODO 修改左侧树
        }
    }

    useEffect(() => {
        loadList(0);
    }, [])


    return <Row gutter={[16, 16]}>
        <Col flex="250px">
            <Search style={{ marginBottom: 8 }} placeholder="Search" />
        </Col>
        <Col flex="auto">
            <Space>
                <Button type='primary' onClick={() => setDrawerProp({...drawerProp, createVisible: true})}>创建</Button>
                <Button type='primary'>刷新</Button>
                <Button type='primary' danger>删除</Button>
            </Space>
            <Table
                bordered
                size={'small'}
                loading={loading}
                columns={columns}
                dataSource={datasource}
                style={{ marginTop: 10 }}
                rowKey={(record) => record.id}
                pagination={false}
                rowSelection={{
                    type: 'checkbox',
                    selectedRowKeys: selectedRowKeys,
                    onChange: (selectedRowKeys: React.Key[]) => {
                        setSelectedRowKeys([...selectedRowKeys.map(item => item as number)])
                    }
                }}
            />
        </Col>
        <MenuCreateDrawer
            types={drawerProp.types!}
            parentId={drawerProp.parentId!}
            visible={drawerProp.createVisible}
            close={closeDrawer}
        />
        <MenuUpdateDrawer
            id={drawerProp.currId!}
            visible={drawerProp.updateVisible}
            close={closeDrawer} />
    </Row>
}

export default AuthorityPermissionPage;