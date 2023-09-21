import React from "react";
import lazyLoad from "@/router/lazyLoad.tsx";

const Layout = lazyLoad(() => import("@/pages/layout"));
const HomePage = lazyLoad(() => import("@/pages/home"));

const LoggerOperatePage = lazyLoad(() => import("@/pages/logger/operate"));
const LoggerLoginPage = lazyLoad(() => import("@/pages/logger/login"));

const SystemUserPage = lazyLoad(() => import("@/pages/system/user"));
const SystemRolePage = lazyLoad(() => import("@/pages/system/role"));
const SystemMenuPage = lazyLoad(() => import("@/pages/system/menu"));
const SystemDeptPage = lazyLoad(() => import("@/pages/system/dept"));
const SystemPostPage = lazyLoad(() => import("@/pages/system/post"));
const SystemDictPage = lazyLoad(() => import("@/pages/system/dict"));
const SystemParamPage = lazyLoad(() => import("@/pages/system/param"));
const SystemInformPage = lazyLoad(() => import("@/pages/system/inform"));

const MonitorOnlinePage = lazyLoad(() => import("@/pages/monitor/online"));
const MonitorTaskPage = lazyLoad(() => import("@/pages/monitor/task"));
const MonitorServerPage = lazyLoad(() => import("@/pages/monitor/server"));
const MonitorCachePage = lazyLoad(() => import("@/pages/monitor/cache"));
const MonitorCacheListPage = lazyLoad(() => import("@/pages/monitor/cacheList"));

const LoginPage = lazyLoad(() => import("@/pages/login"));

const PersonUserPage = lazyLoad(() => import("@/pages/user"));

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

export const RouterMap = new Map<String, React.ReactNode>([
    ["Layout", Layout],
    ["HomePage", HomePage],
    ["LoggerOperatePage", LoggerOperatePage],
    ["LoggerLoginPage", LoggerLoginPage],
    ["SystemUserPage", SystemUserPage],
    ["SystemRolePage", SystemRolePage],
    ["SystemMenuPage", SystemMenuPage],
    ["SystemPostPage", SystemPostPage],
    ["SystemDeptPage", SystemDeptPage],
    ["SystemDictPage", SystemDictPage],
    ["SystemParamPage", SystemParamPage],
    ["SystemInformPage", SystemInformPage],
    ["MonitorOnlinePage", MonitorOnlinePage],
    ["MonitorTaskPage", MonitorTaskPage],
    ["MonitorServerPage", MonitorServerPage],
    ["MonitorCachePage", MonitorCachePage],
    ["MonitorCacheListPage", MonitorCacheListPage],
    ["PersonUserPage", PersonUserPage],
    ["LoginPage", LoginPage],
])