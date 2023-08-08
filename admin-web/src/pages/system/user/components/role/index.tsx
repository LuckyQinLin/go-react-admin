import React from "react";
import {Button, Drawer, Space} from "antd";

const UserRoleDrawer: React.FC<UserRoleDrawerProp> = ({userId, visible, close}) => {
    return <Drawer
        width={500}
        title="分配角色"
        placement="right"
        onClose={() => close(false)}
        open={visible}
        extra={
            <Space>
                <Button type="primary" danger onClick={()=> close(false)}>取消</Button>
                <Button type="primary" onClick={submitForm}>保存</Button>
            </Space>
        }
    ></Drawer>
}

export default UserRoleDrawer;