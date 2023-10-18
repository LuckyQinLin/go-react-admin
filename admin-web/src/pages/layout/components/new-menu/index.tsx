import React, {useEffect, useMemo, useState} from "react";
import {Layout, Menu, MenuProps} from "antd";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import './index.less';
import {Link, useNavigate, useRouteLoaderData} from "react-router-dom";
import {User, Menu as Menus} from "@/types";
import IconFont from "@/components/IconFont";
import Router from "@/new-router/modules.tsx";
import {asyncRoutes} from "@/router";
import {routerBuild, routerBuildMenu} from "@/router/routerFilter.tsx";
import HomeRouter from "@/router/modules/home.tsx";
import PersonRouter from "@/router/modules/person.tsx";
type MenuItem = Required<MenuProps>['items'][number]

interface LayoutHeaderProp {
    collapsed: boolean;
    breadcrumb: (data: BreadcrumbProp[]) => void;
}

const home: Menus.MenuItemProp = {
    pId: 1,
    id: 1,
    sort: 0,
    title: '首页',
    perms: 'home',
    path: '/home/index',
    icon: 'lucky-shouye1',
    types: 'M',
    component: '',
}

const person: Menus.MenuItemProp = {
    pId: 0,
    id: 1,
    sort: 0,
    title: '个人中心',
    perms: 'person',
    path: '/person/index',
    icon: 'lucky-jiankong',
    types: 'M',
    component: '',
}


const LayoutNewSider: React.FC<LayoutHeaderProp> = ({collapsed}) => {

    let navigate = useNavigate();
    const data = useRouteLoaderData(Router.LayoutId) as User.UserPermissionProp;
    const [menuList, setMenuList] = useState<MenuItem[]>([]);
    const [openKeys, setOpenKeys] = useState<string[]>([]);
    const [selectedKeys, setSelectedKeys] = useState<string[]>(['home']);

    const buildMenuItem = (label: React.ReactNode, key: React.Key, icon?: React.ReactNode, children?: MenuItem[]): MenuItem => {
        return { key, icon, children, label } as MenuItem;
    }

    const routerBuildMenu = (data: Menus.MenuItemProp[]): MenuItem[] => {
        return data.map((item) => {
            const icon = item.icon ? <IconFont type={item.icon} /> : undefined;
            const title = item.title;
            const key = item.path ?? item.perms;
            if (item.children && item.children.length > 0) {
                const children = routerBuildMenu(item.children)
                return buildMenuItem(
                    title,
                    key,
                    icon,
                    children
                )
            } else {
                return buildMenuItem(
                    title,
                    key,
                    icon,
                )
            }
        });
    }

    const clickMenu = (key: string) => {
        setSelectedKeys([key])
        navigate(key)
    }

    const openSubKey = (keys: string[]) => {
        const items = openKeys.length <= 0 ? keys : keys.filter(item => !openKeys.includes(item))
        console.log("keys", keys, items);
        setOpenKeys(items)
    }


    useEffect(() => {
        // const treeMenu = routerBuildMenu([home, ...data.menus, person]);
        const treeMenu = routerBuildMenu(data.menus);
        const items = Router.menuItems;
        console.log("new-tree Menu => ", treeMenu);
        console.log("new-tree all Menu => ", items);
        setMenuList(items!);
    }, []);


    return <Layout.Sider width={230} trigger={null} collapsible collapsed={collapsed}>
        <div className="ant-pro-sider-logo">
            <a href="">
                <img src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg" alt="logo"/>
                {!collapsed && <h1>通用管理系统</h1>}
            </a>
        </div>
        <Menu
            mode="inline"
            theme="dark"
            inlineCollapsed={collapsed}
            items={Router.menuItems}
            // openKeys={openKeys}
            openKeys={openKeys}
            selectedKeys={selectedKeys}
            onOpenChange={openSubKey}
            onClick={({ key}) => clickMenu(key)}
        />
    </Layout.Sider>
}

export default LayoutNewSider;
