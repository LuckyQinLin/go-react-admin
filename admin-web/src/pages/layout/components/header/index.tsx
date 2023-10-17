import React from "react";
import {Avatar, Breadcrumb, Button, Dropdown, Layout, MenuProps, message, Space} from "antd";
import {
    HomeOutlined,
    LogoutOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    RedoOutlined,
    SettingOutlined
} from "@ant-design/icons";
import {ItemType as BreadcrumbType} from "antd/es/breadcrumb/Breadcrumb";
import './index.less';
import {AdminIcon} from "@/components";
import {useDispatch} from "react-redux";
import {cleanUserStoreActionCreator} from "@/redux/user/action";
import {useNavigate} from "react-router-dom";
import {useSelector} from "@/redux/hooks";

interface LayoutHeaderProp {
    breadcrumb?: BreadcrumbProp[];
    bgColor: string;
    collapsed: boolean;
    changeCollapsed: (collapsed: boolean) => void;
}

export interface BreadcrumbProp {
    path: string;
    label: string;
    icon?: React.ReactNode;
}

const LayoutHeader: React.FC<LayoutHeaderProp> = ({breadcrumb, bgColor, collapsed, changeCollapsed}) => {

    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [messageApi, contextHolder] = message.useMessage();
    const {user} = useSelector((state) => state.user);

    const dropdownItems: MenuProps['items'] = [
        { key: '1', icon: <SettingOutlined />, label: '个人设置'},
        { type: 'divider' },
        { key: '2', icon: <LogoutOutlined />, label: '退出登录', },
    ];

    const languageItems: MenuProps = {
        items: [
            { key: '1', label: '中文'},
            { key: '2', label: 'English'},
        ], onClick: (e) => {
            console.log(e.key)
        }
    }

    const onClick: MenuProps['onClick'] = (e) => {
        if (e.key === '2') {
            messageApi.success('退出成功');
            dispatch(cleanUserStoreActionCreator());
            navigate('/login');
        }
    };



    const items = (): BreadcrumbType[] => {
				const temp: BreadcrumbType[] = [];
				breadcrumb?.forEach(item => temp.push({
					href: item.path,
					title: item.icon ? <>{item.icon}<span>{item.label}</span></> : <span>{item.label}</span>
				}))
				return [{href: '', title: <HomeOutlined />}, ...temp];
    }

    return <Layout.Header className="admin-layout-header" style={{ padding: 0, background: bgColor, height: 55 }}>
        {contextHolder}
        <Space className="admin-layout-header-left">
            <Button
                type="link"
                icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                onClick={() => changeCollapsed(!collapsed)}
                style={{
                    color: 'rgba(0, 0, 0, 0.88)',
                    fontSize: '16px',
                    width: 10,
                    height: 55,
                    marginLeft: 16
                }}
            />
            <Button
                type="link"
                icon={<RedoOutlined />}
                style={{
                    color: 'rgba(0, 0, 0, 0.88)',
                    fontSize: '16px',
                    width: 55,
                    height: 64,
                }}
            />
            <Breadcrumb items={items()} />
        </Space>
        <Space className="admin-layout-header-right">
            <Dropdown menu={{items: dropdownItems, onClick}} placement="top" arrow className="header-right">
                <span className="alhr-dropdown">
                    <Avatar src={<img src={'https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png'} alt="avatar" />} />
                    <span className="alhrd-title">{user?.userName}</span>
                </span>
            </Dropdown>
            <Dropdown menu={languageItems} placement="bottomRight" arrow className="header-right">
                <span className="alhr-dropdown" style={{padding: '0 5px'}}>
                    <AdminIcon type="icon-yuyanqiehuan2" style={{fontSize: 20, margin: '22px 0', color: '#595959'}} />
                </span>
            </Dropdown>
        </Space>
    </Layout.Header>
}

export default LayoutHeader;
