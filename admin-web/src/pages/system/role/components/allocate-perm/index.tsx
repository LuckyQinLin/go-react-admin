import React from "react";
import {AllocatePermDrawerProp} from "@/pages/system/role/modules.ts";
import {Button, Drawer, Space, Spin} from "antd";


const AllocatePermDrawer: React.FC<AllocatePermDrawerProp> = ({visible, roleId, close}) => {

    const submitForm = () => {

    }

    return <Drawer
        width={500}
        title="分配资源"
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
        <Spin tip="加载中......" spinning={true}>

        </Spin>
    </Drawer>
}

export default AllocatePermDrawer;