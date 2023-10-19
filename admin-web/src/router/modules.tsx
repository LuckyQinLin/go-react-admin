import React from "react";
import lazyLoad from "@/router/lazyLoad.tsx";
import type {MenuProps} from "antd/es/menu";
import IconFont from "@/components/IconFont";
import {Menus} from "@/types";
type MenuItem = Required<MenuProps>['items'][number];


function getItem(
    label: React.ReactNode,
    key: React.Key,
    icon?: React.ReactNode,
    children?: MenuItem[]
): MenuItem {
    return {
        key,
        icon,
        children,
        label
    } as MenuItem;
}


namespace RouterVariate {

    import MenuTitleProp = Menus.MenuTitleProp;

    export enum RouterEnum {
        Layout,
        HomePage,
        LoginPage,


        SystemUserPage,
        SystemRolePage,
        SystemMenuPage,
        SystemDeptPage,
        SystemPostPage,
        SystemDictPage,
        SystemParamPage,
        SystemInformPage,

        MonitorOnlinePage,
        MonitorTaskPage,
        MonitorServerPage,
        MonitorCachePage,
        MonitorCacheListPage,

        LoggerOperatePage,
        LoggerLoginPage,

        NotFoundPage,
        NotAuthPage,
        ServerErrPage,

        PersonUserPage
    }

    export const staticPath = ['/home/index', '/person/index', '/exception/404', '/exception/403', '/exception/500'];

    export const NotFoundPath = '/exception/404';
    export const NotAuthPath = '/exception/403';

    export const LayoutId: string = "Layout"

    export const getRouter = (key: RouterEnum): React.ReactElement | undefined => {
        switch (key) {
            case RouterEnum.Layout:
                return lazyLoad(() => import("@/pages/layout"));
            case RouterEnum.HomePage:
                return lazyLoad(() => import("@/pages/home"));
            case RouterEnum.LoginPage:
                return lazyLoad(() => import("@/pages/login"));
            case RouterEnum.SystemUserPage:
                return lazyLoad(() => import("@/pages/system/user"));
            case RouterEnum.SystemRolePage:
                return lazyLoad(() => import("@/pages/system/role"));
            case RouterEnum.SystemDeptPage:
                return lazyLoad(() => import("@/pages/system/dept"));
            case RouterEnum.SystemMenuPage:
                return lazyLoad(() => import("@/pages/system/menu"));
            case RouterEnum.SystemDictPage:
                return lazyLoad(() => import("@/pages/system/dict"));
            case RouterEnum.SystemInformPage:
                return lazyLoad(() => import("@/pages/system/inform"));
            case RouterEnum.SystemParamPage:
                return lazyLoad(() => import("@/pages/system/param"));
            case RouterEnum.SystemPostPage:
                return lazyLoad(() => import("@/pages/system/post"));
            case RouterEnum.MonitorServerPage:
                return lazyLoad(() => import("@/pages/monitor/server"));
            case RouterEnum.MonitorCacheListPage:
                return lazyLoad(() => import("@/pages/monitor/cacheList"));
            case RouterEnum.MonitorCachePage:
                return lazyLoad(() => import("@/pages/monitor/cache"));
            case RouterEnum.MonitorOnlinePage:
                return lazyLoad(() => import("@/pages/monitor/online"));
            case RouterEnum.MonitorTaskPage:
                return lazyLoad(() => import("@/pages/monitor/task"));
            case RouterEnum.LoggerLoginPage:
                return lazyLoad(() => import("@/pages/logger/login"));
            case RouterEnum.LoggerOperatePage:
                return lazyLoad(() => import("@/pages/logger/operate"));
            case RouterEnum.PersonUserPage:
                return lazyLoad(() => import("@/pages/user"));
            case RouterEnum.NotFoundPage:
                return lazyLoad(() => import("@/pages/exception/404.tsx"));
            case RouterEnum.NotAuthPage:
                return lazyLoad(() => import("@/pages/exception/403.tsx"));
            case RouterEnum.ServerErrPage:
                return lazyLoad(() => import("@/pages/exception/500.tsx"));
        }
    }

    export interface IRouteObject {
        children?: Array<IRouteObject>,
        element?: React.ReactNode,
        redirect?: string,
        isRoot?: boolean,
        path: string
        meta?: {
            key?: string;
            isRoot?: boolean,
            title?: string,
            sort: number,
            icon?: string,
            permission?: Array<string>
        }
    }

    export const HomeItems = [getItem('首页', '/home/index', <IconFont type="lucky-shouye1" />)];
    export const PersonItems = [getItem('个人中心', '/person/index', <IconFont type="lucky-jiankong" />)];

    export const menuItems: MenuItem[] = [
        ...HomeItems,
        getItem('系统管理', '/system', <IconFont type="lucky-shezhi1" />, [
            getItem('用户管理', '/system/user', <IconFont type="lucky-yonghu" />),
            getItem('角色管理', '/system/role', <IconFont type="lucky-jiaose" />),
            getItem('菜单管理', '/system/menu', <IconFont type="lucky-caidan" />),
            getItem('部门管理', '/system/dept', <IconFont type="lucky-bumenguanli" />),
            getItem('岗位管理', '/system/post', <IconFont type="lucky-gangwei" />),
            getItem('字典管理', '/system/dict', <IconFont type="lucky-zidianmuluguanli" />),
            getItem('参数管理', '/system/param', <IconFont type="lucky-shujucanshu" />),
            getItem('通知公告', '/system/inform', <IconFont type="lucky-tongzhi" />),
        ]),

        getItem('系统监控', '/monitor', <IconFont type="lucky-jiankong" />, [
            getItem('在线用户', '/monitor/onlineUser', <IconFont type="lucky-zaixianyonghuguanli1" />),
            getItem('定时任务', '/monitor/timeTask', <IconFont type="lucky-dingshirenwuguanli" />),
            getItem('服务器监控', '/monitor/server', <IconFont type="lucky-fuwuqijiankong" />),
            getItem('缓冲监控', '/monitor/cache', <IconFont type="lucky-huanchongfenxi" />),
            getItem('缓冲列表', '/monitor/cacheList', <IconFont type="lucky-cityworksjichugongnengtubiao-" />),
        ]),
        getItem('日志管理', '/logger', <IconFont type="lucky-nav_icon_rzgl_spe" />, [
            getItem('操作日志', '/logger/operate', <IconFont type="lucky-caozuorizhi" />),
            getItem('登录日志', '/logger/login', <IconFont type="lucky-denglurizhi" />),
        ]),
        ...PersonItems
    ];

    export const menuTitleItems: MenuTitleProp[] = [
        {title: '首页', path: '/home/index'},

        {title: '用户管理', path: '/system/user'},
        {title: '角色管理', path: '/system/role'},
        {title: '菜单管理', path: '/system/menu'},
        {title: '部门管理', path: '/system/dept'},
        {title: '岗位管理', path: '/system/post'},
        {title: '字典管理', path: '/system/dict'},
        {title: '参数管理', path: '/system/param'},
        {title: '通知公告', path: '/system/inform'},

        {title: '在线用户', path: '/monitor/onlineUser'},
        {title: '定时任务', path: '/monitor/timeTask'},
        {title: '服务器监控', path: '/monitor/server'},
        {title: '缓冲监控', path: '/monitor/cache'},
        {title: '缓冲列表', path: '/monitor/cacheList'},

        {title: '操作日志', path: '/logger/operate'},
        {title: '登录日志', path: '/logger/login'},

        {title: '404', path: '/exception/404'},
        {title: '403', path: '/exception/403'},
        {title: '500', path: '/exception/500'},

        {title: '个人中心', path: '/person/index'},
    ]

}

export default RouterVariate;
