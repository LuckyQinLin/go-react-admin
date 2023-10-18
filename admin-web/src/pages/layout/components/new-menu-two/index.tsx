import React, {useEffect, useMemo, useState} from "react";
import {Layout, Menu, MenuProps} from "antd";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import './index.less';
import {Link, useRouteLoaderData} from "react-router-dom";
import {User, Menu as Menus} from "@/types";
import IconFont from "@/components/IconFont";
import Router from "@/new-router/modules.tsx";
import {asyncRoutes} from "@/router";
import {routerBuild, routerBuildMenu} from "@/router/routerFilter.tsx";
import HomeRouter from "@/router/modules/home.tsx";
import PersonRouter from "@/router/modules/person.tsx";
import {AppstoreOutlined, ContainerOutlined, DesktopOutlined, MailOutlined, PieChartOutlined} from "@ant-design/icons";
type MenuItem = Required<MenuProps>['items'][number]

interface LayoutHeaderProp {
    collapsed: boolean;
    breadcrumb: (data: BreadcrumbProp[]) => void;
}



function getItem(
    label: React.ReactNode,
    key: React.Key,
    icon?: React.ReactNode,
    children?: MenuItem[],
    type?: 'group',
): MenuItem {
    return {
        key,
        icon,
        children,
        label,
        type,
    } as MenuItem;
}


const LayoutNewTwoSider: React.FC<LayoutHeaderProp> = ({collapsed}) => {

    // const data = useRouteLoaderData(Router.LayoutId) as User.UserPermissionProp;
    // const [menuList, setMenuList] = useState<MenuItem[]>([]);
    // const [defaultMenu, setDefaultMenu] = useState<string>('home');
    // const [openKeys] = useState<string[]>([]);

    const items: MenuItem[] = [
        getItem('Option 1', '1', <PieChartOutlined />),
        getItem('Option 2', '2', <DesktopOutlined />),
        getItem('Option 3', '3', <ContainerOutlined />),

        getItem('Navigation One', 'sub1', <MailOutlined />, [
            getItem('Option 5', '5'),
            getItem('Option 6', '6'),
            getItem('Option 7', '7'),
            getItem('Option 8', '8'),
        ]),

        getItem('Navigation Two', 'sub2', <AppstoreOutlined />, [
            getItem('Option 9', '9'),
            getItem('Option 10', '10'),

            getItem('Submenu', 'sub3', null, [getItem('Option 11', '11'), getItem('Option 12', '12')]),
        ]),
    ];

    // const clickMenu = (key: string) => {
    //     setDefaultMenu(key)
    // }


    // useEffect(() => {
    //     // const treeMenu = routerBuildMenu([home, ...data.menus, person]);
    //     const treeMenu = routerBuildMenu(data.menus);
    //     const items = Router.menuItems;
    //     console.log("new-tree Menu => ", treeMenu);
    //     console.log("new-tree all Menu => ", items);
    //     setMenuList(items!);
    // }, []);


    return <Layout.Sider width={230} trigger={null} collapsible collapsed={collapsed}>
        <div className="ant-pro-sider-logo">
            <a href="">
                <img src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg" alt="logo"/>
                {!collapsed && <h1>通用管理系统</h1>}
            </a>
        </div>
        <Menu
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            mode="inline"
            theme="dark"
            inlineCollapsed={collapsed}
            items={Router.menuItems}
        />
    </Layout.Sider>
}

export default LayoutNewTwoSider;
