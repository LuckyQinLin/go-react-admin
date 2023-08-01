import {Button, message, Modal, Space, Table, Tag} from "antd";
import {useEffect, useState} from "react";
import {ColumnsType} from "antd/es/table";
import {DeptDrawerProp, DeptTableQueryProp, DeptTableTreeProp} from "@/pages/system/dept/modules.ts";
import {useRequest} from "ahooks";
import {deptDelete, deptTable} from "@/api/dept.ts";
import {ExclamationCircleFilled} from "@ant-design/icons";
import {DeptCreateDrawer, DeptUpdateDrawer} from "@/pages/system/dept/components";

const SystemDeptPage = () => {

    const columns: ColumnsType<DeptTableTreeProp> = [
        {
            title: '部门名称',
            key: 'title',
            dataIndex: 'title',
            align: 'center',
            width: 150
        },
        {
            title: '排序',
            key: 'order',
            dataIndex: 'order',
            align: 'center',
            width: 80
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

    const [tableQuery] = useState<DeptTableQueryProp>({});
    const [drawerProp, setDrawerProp] = useState<DeptDrawerProp>({createVisible: false, updateVisible: false});
    const [datasource, setDatasource] = useState<DeptTableTreeProp[]>([]);

    const {run, loading} = useRequest(deptTable, {
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

    const deleteMenu = (deptId: number) => {
        Modal.confirm({
            title: '警告',
            icon: <ExclamationCircleFilled />,
            content: '确认删除当前部门？',
            okText: '删除',
            okType: 'danger',
            cancelText: '取消',
            onOk: async () => {
                const msg = await deptDelete(deptId);
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
        />
        <DeptCreateDrawer
            parentId={drawerProp.parentId!}
            visible={drawerProp.createVisible}
            close={closeDrawer}
        />
        <DeptUpdateDrawer
            deptId={drawerProp.currId!}
            visible={drawerProp.updateVisible}
            close={closeDrawer} />
    </>
}

export default SystemDeptPage;