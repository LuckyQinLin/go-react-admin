import React, {useEffect, useState} from "react";
import {Layout, Menu, MenuProps} from "antd";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import {useNavigate, useRouteLoaderData} from "react-router-dom";
import {User, Menu as Menus} from "@/types";
import IconFont from "@/components/IconFont";
import Router from "@/new-router/modules.tsx";
import useStore from "@/store/store.ts";
import './index.less';
import {HOME_PAGE} from "@/constant/setting.ts";
import HomeItems = Router.HomeItems;
import PersonItems = Router.PersonItems;
type MenuItem = Required<MenuProps>['items'][number]

interface LayoutHeaderProp {
    collapsed: boolean;
    breadcrumb: (data: BreadcrumbProp[]) => void;
}


const LayoutNewSider: React.FC<LayoutHeaderProp> = ({collapsed}) => {

    let navigate = useNavigate();
    const data = useRouteLoaderData(Router.LayoutId) as User.UserPermissionProp;
    const userProp = useStore(state => state.userProp)
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
        if (userProp === undefined) {
            navigate(HOME_PAGE)
            return
        }
        setMenuList(userProp.isSuper ?
            Router.menuItems :
            [
                ...HomeItems,
                ...routerBuildMenu(data.menus),
                ...PersonItems
            ]
        );
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
            items={menuList}
            openKeys={openKeys}
            selectedKeys={selectedKeys}
            onOpenChange={openSubKey}
            onClick={({ key}) => clickMenu(key)}
        />
    </Layout.Sider>
}

export default LayoutNewSider;
