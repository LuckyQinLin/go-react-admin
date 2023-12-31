import React, {useEffect, useState} from "react";
import {Layout, Menu, MenuProps} from "antd";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import {useNavigate, useRouteLoaderData} from "react-router-dom";
import {User, Menus} from "@/types";
import IconFont from "@/components/IconFont";
import RouterVariate from "@/router/modules.tsx";
import useStore from "@/store/store.ts";
type MenuItem = Required<MenuProps>['items'][number]
import './index.less';

interface LayoutHeaderProp {
    collapsed: boolean;
    breadcrumb: (data: BreadcrumbProp[]) => void;
}


const LayoutSider: React.FC<LayoutHeaderProp> = ({collapsed}) => {

    let navigate = useNavigate();
    const data = useRouteLoaderData(RouterVariate.LayoutId) as User.UserPermissionProp;
    const userProp = useStore(state => state.userProp);
    const addTabView = useStore(state => state.addTabView);
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


    const clickMenu = (data: string) => {
        const target = RouterVariate.menuTitleItems.find(item => data === item.path);
        console.log("clickMenu", target)
        if (target) {
            addTabView({title: target.title, key: target.path as string})
            setSelectedKeys([data])
            navigate(data)
        }
    }

    const openSubKey = (keys: string[]) => {
        const items = openKeys.length <= 0 ? keys : keys.filter(item => !openKeys.includes(item))
        console.log("keys", keys, items);
        setOpenKeys(items)
    }


    useEffect(() => {
        // if (userProp === undefined) {
        //     navigate(HOME_PAGE)
        //     return
        // }
        // debugger;
        const items = userProp?.isSuper ?
            RouterVariate.menuItems :
            [
                ...RouterVariate.HomeItems,
                ...routerBuildMenu(data.menus),
                ...RouterVariate.PersonItems
            ]
        console.log("menu", userProp, data, items)
        setMenuList(items);
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
            items={menuList}
            openKeys={openKeys}
            selectedKeys={selectedKeys}
            onOpenChange={openSubKey}
            onClick={({key}) => clickMenu(key)}
        />
    </Layout.Sider>
}

export default LayoutSider;
