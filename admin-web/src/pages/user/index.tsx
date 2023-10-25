import {Avatar, Card, Col, Row, Tabs, TabsProps} from "antd";
import {
    AntDesignOutlined, FieldTimeOutlined,
    MailOutlined, PartitionOutlined,
    PhoneOutlined,
    TeamOutlined,
    UserOutlined
} from "@ant-design/icons";
import styled from "@emotion/styled";
import {BaseInfoForm, PasswordForm} from "@/pages/user/components";

const UserCenterPage = () => {

    const items: TabsProps['items'] = [
        {
            key: '1',
            label: `基础信息`,
            children: <BaseInfoForm />,
        },
        {
            key: '2',
            label: `修改密码`,
            children: <PasswordForm />,
        }
    ];

    return <>
        <Container gutter={8}>
            <Col span={6} style={{marginRight: 0}}>
                <Card title="个人信息" bordered={false}>
                    <Avatar
                        size={{ xs: 24, sm: 32, md: 40, lg: 64, xl: 80, xxl: 100 }}
                        icon={<AntDesignOutlined />}
                    />
                    <div className="info_user">
                        <div className="info_user_item">
                            <div><UserOutlined style={{marginRight: 5}} />用户名称</div>
                            <div>张三</div>
                        </div>
                        <div className="info_user_item">
                            <div><PhoneOutlined style={{marginRight: 5}}/>手机号码</div>
                            <div>13323065745</div>
                        </div>
                        <div className="info_user_item">
                            <div><MailOutlined style={{marginRight: 5}}/>邮箱</div>
                            <div>2354713722@qq.com</div>
                        </div>
                        <div className="info_user_item">
                            <div><PartitionOutlined style={{marginRight: 5}}/>所属部门</div>
                            <div>开发部/项目经理</div>
                        </div>
                        <div className="info_user_item">
                            <div><TeamOutlined style={{marginRight: 5}}/>所属角色</div>
                            <div>超级管理员</div>
                        </div>
                        <div className="info_user_item">
                            <div><FieldTimeOutlined style={{marginRight: 5}} />创建时间</div>
                            <div>2023-08-31 10:47:33</div>
                        </div>
                    </div>
                </Card>
            </Col>
            <Col span={18}>
                <Card title="基础资料" bordered={false}>
                    <Tabs defaultActiveKey="1" items={items} />
                </Card>
            </Col>
        </Container>
    </>
}

const Container = styled(Row)`
    .info_user {
        margin-top: 15px;
        display: flex;
        flex-direction: column;
        flex-wrap: nowrap;
        .info_user_item {
            display: flex;
            flex-direction: row;
            flex-wrap: nowrap;
            justify-content: space-between;
            align-items: center;
            height: 50px;
            border-top: 1px solid #ccc;
            .info_user_item_left {
            
            }
            .info_user_item_right {
            
            }
            &:last-child {
                border-bottom: 1px solid #ccc; /* 添加底部边框 */
            }
        }
    }

`

export default UserCenterPage;