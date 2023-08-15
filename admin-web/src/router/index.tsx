import {Navigate, RouteObject, useRoutes} from "react-router-dom";
import LoginPage from "@/pages/login";
import LayoutPage from "@/pages/layout";
import HomePage from "@/pages/home";
import NotAuthPage from "@/pages/exception/403.tsx";
import NotFoundPage from "@/pages/exception/404.tsx";
import ServerErrorPage from "@/pages/exception/500.tsx";
import SystemUserPage from "@/pages/system/user";
import SystemRolePage from "@/pages/system/role";
import SystemMenuPage from "@/pages/system/menu";
import SystemDeptPage from "@/pages/system/dept";
import SystemPostPage from "@/pages/system/post";
import SystemDictPage from "@/pages/system/dict";
import SystemParamPage from "@/pages/system/param";
import SystemInformPage from "@/pages/system/inform";
import LoggerOperatePage from "@/pages/logger/operate";
import LoggerLoginPage from "@/pages/logger/login";
import MonitorOnlinePage from "@/pages/monitor/online";
import MonitorTaskPage from "@/pages/monitor/task";
import MonitorServerPage from "@/pages/monitor/server";
import MonitorCachePage from "@/pages/monitor/cache";
import MonitorCacheListPage from "@/pages/monitor/cacheList";

export interface IRouteObject {
    children?: IRouteObject[]
    element?: React.ReactNode
    index?: boolean
    path?: string
    meta?: {
        perm?: string
    }
}

export const Routes: IRouteObject[] = [
    {
        path: '*',
        element: <Navigate to={'/index'} />
    },
    {
        path: 'login',
        element: <LoginPage />
    },
    {
        path: '/',
        element: <LayoutPage />,
        children: [
            {
                path: 'index',
                element: <HomePage />
            },
            {
                path: '/system/user',
                element: <SystemUserPage />,
                meta: {perm: "system:user:list"}
            },
            {
                path: '/system/role',
                element: <SystemRolePage />
            },
            {
                path: '/system/menu',
                element: <SystemMenuPage />
            },
            {
                path: '/system/dept',
                element: <SystemDeptPage />
            },
            {
                path: '/system/post',
                element: <SystemPostPage />
            },
            {
                path: '/system/dict',
                element: <SystemDictPage />
            },
            {
                path: '/system/param',
                element: <SystemParamPage />
            },
            {
                path: '/system/inform',
                element: <SystemInformPage />
            },
            {
                path: '/logger/operate',
                element: <LoggerOperatePage />
            },
            {
                path: '/logger/login',
                element: <LoggerLoginPage />
            },
            {
                path: '/monitor/onlineUser',
                element: <MonitorOnlinePage />
            },
            {
                path: '/monitor/timeTask',
                element: <MonitorTaskPage />
            },
            {
                path: '/monitor/server',
                element: <MonitorServerPage />
            },
            {
                path: '/monitor/cache',
                element: <MonitorCachePage />
            },
            {
                path: '/monitor/cacheList',
                element: <MonitorCacheListPage />
            },
            {
                path: '/403',
                element: <NotAuthPage />
            },
            {
                path: '/404',
                element: <NotFoundPage />
            },
            {
                path: '/500',
                element: <ServerErrorPage />
            },
            {
                path: '/*',
                element: <NotFoundPage />
            },
        ]
    }
]

export const SystemRouter = (): React.ReactElement | null => {
    // const user = useSelector((state) => state.user)
    return useRoutes(Routes as RouteObject[]);
}