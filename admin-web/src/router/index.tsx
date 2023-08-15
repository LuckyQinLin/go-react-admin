import {Navigate, useRoutes} from "react-router-dom";
import {useSelector} from "@/redux/hooks";
import NotAuthPage from "@/pages/exception/403.tsx";
import NotFoundPage from "@/pages/exception/404.tsx";
import ServerErrorPage from "@/pages/exception/500.tsx";
// import HomePage from "@/pages/home";
// import SystemUserPage from "@/pages/system/user";
// import SystemRolePage from "@/pages/system/role";
// import SystemMenuPage from "@/pages/system/menu";
// import SystemDeptPage from "@/pages/system/dept";
// import SystemPostPage from "@/pages/system/post";
// import SystemDictPage from "@/pages/system/dict";
// import SystemParamPage from "@/pages/system/param";
// import SystemInformPage from "@/pages/system/inform";
// import LoggerOperatePage from "@/pages/logger/operate";
// import LoggerLoginPage from "@/pages/logger/login";
// import MonitorOnlinePage from "@/pages/monitor/online";
// import MonitorTaskPage from "@/pages/monitor/task";
// import MonitorServerPage from "@/pages/monitor/server";
// import MonitorCachePage from "@/pages/monitor/cache";
// import MonitorCacheListPage from "@/pages/monitor/cacheList";
import {lazy, Suspense} from "react";
import {Spin} from "antd";
import {RouteObject} from "react-router/dist/lib/context";

// https://blog.xav1er.com/p/react-router-dom-v6-dynamic-router/
// const lazyLoad_1 = (factory: () => Promise<any>) => {
//     const Module = lazy(factory)
//     return (
//         <Suspense fallback={<Spin tip="加载中......" />}>
//             <Module />
//         </Suspense>
//     )
// }

// https://juejin.cn/post/7132393527501127687
const lazyLoad = (moduleName: string) => {
    const Module = lazy(() => import(`@/pages/${moduleName}`));
    return <Suspense fallback={<Spin tip="加载中......" />}><Module /></Suspense>;
};

// 路由鉴权组件
const AuthRoute = ({ children }: any) => {
    const user = useSelector((state) => state.user)
    return user.status && user.token ? children : <Navigate to="/login" />;
};

// 公共路由
const defaultRoutes: RouteObject[] = [
    {
        path: 'login',
        element: lazyLoad('login'),
    },
    {
        path: '/',
        element: <AuthRoute>{lazyLoad('layout')}</AuthRoute>,
        children: [
            {
                path: '/index',
                element: lazyLoad('home')
            },
            {
                path: '/exception/403',
                element: <NotAuthPage />
            },
            {
                path: '/exception/404',
                element: <NotFoundPage />
            },
            {
                path: '/exception/500',
                element: <ServerErrorPage />
            }
        ]
    }
]

// 私有路由
const privateRoutes: RouteObject[] = [
    {
        path: '/',
        element: <AuthRoute>{lazyLoad('layout')}</AuthRoute>,
        children: [
            {
                path: '/system/user',
                element: lazyLoad('system/user')
            },
            {
                path: '/system/role',
                element: lazyLoad('system/role')
            },
            {
                path: '/system/menu',
                element: lazyLoad('system/menu')
            },
            {
                path: '/system/dept',
                element: lazyLoad('system/dept')
            },
            {
                path: '/system/post',
                element: lazyLoad('system/post')
            },
            {
                path: '/system/dict',
                element: lazyLoad('system/dict')
            },
            {
                path: '/system/param',
                element: lazyLoad('system/param')
            },
            {
                path: '/system/inform',
                element: lazyLoad('system/inform')
            },
            {
                path: '/logger/operate',
                element: lazyLoad('logger/operate')
            },
            {
                path: '/logger/login',
                element: lazyLoad('logger/login')
            },
            {
                path: '/monitor/onlineUser',
                element: lazyLoad('monitor/onlineUser')
            },
            {
                path: '/monitor/timeTask',
                element: lazyLoad('monitor/timeTask')
            },
            {
                path: '/monitor/server',
                element: lazyLoad('monitor/server')
            },
            {
                path: '/monitor/cache',
                element: lazyLoad('monitor/cache')
            },
            {
                path: '/monitor/cacheList',
                element: lazyLoad('monitor/cacheList')
            },
        ]
    }
]


export const SystemRouter = (): React.ReactElement | null => {

    return useRoutes([...defaultRoutes, ...privateRoutes]);
}