import {useRoutes} from "react-router-dom";
// import LoginPage from "@/pages/login";
// import LayoutPage from "@/pages/layout";
// import HomePage from "@/pages/home";
// import NotAuthPage from "@/pages/exception/403.tsx";
// import NotFoundPage from "@/pages/exception/404.tsx";
// import ServerErrorPage from "@/pages/exception/500.tsx";
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
// import UserCenterPage from "@/pages/user";
import HomeRouter from "@/router/modules/home.tsx";
import ExceptionRouter from "@/router/modules/exception.tsx";
import PersonRouter from "@/router/modules/person.tsx";
// import {IRouteObject} from "@/router/modules.ts";
import SettingRouter from "@/router/modules/setting.tsx";
import LoggerRouter from "@/router/modules/logger.tsx";
import MonitorRouter from "@/router/modules/monitor.tsx";
import LoginRouter from "@/router/modules/login.tsx";
import React, {useEffect} from "react";
import useLoadRoutes from "@/router/useLoadRoutes.tsx";
import {IRouteObject} from "@/router/modules.ts";
import {NOT_FOUND_PAGE} from "@/constant/setting.ts";

// 无需验证的普通路由
export const constantRouter = [
    ...LoginRouter,
    ...HomeRouter,
    ...ExceptionRouter,
    ...PersonRouter,
]

// 需要验证的路由
export const asyncRoutes = [...SettingRouter, ...LoggerRouter, ...MonitorRouter]

// export const Routes: IRouteObject[] = [
//     {
//         path: '*',
//         element: <Navigate to={'/index'} />
//     },
//     {
//         path: 'login',
//         element: <LoginPage />
//     },
//     {
//         path: '/',
//         element: <LayoutPage />,
//         children: [
//             {
//                 path: 'index',
//                 element: <HomePage />
//             },
//             {
//                 path: '/system/user',
//                 element: <SystemUserPage />,
//                 meta: {perm: "system:user:list"}
//             },
//             {
//                 path: '/system/role',
//                 element: <SystemRolePage />
//             },
//             {
//                 path: '/system/menu',
//                 element: <SystemMenuPage />
//             },
//             {
//                 path: '/system/dept',
//                 element: <SystemDeptPage />
//             },
//             {
//                 path: '/system/post',
//                 element: <SystemPostPage />
//             },
//             {
//                 path: '/system/dict',
//                 element: <SystemDictPage />
//             },
//             {
//                 path: '/system/param',
//                 element: <SystemParamPage />
//             },
//             {
//                 path: '/system/inform',
//                 element: <SystemInformPage />
//             },
//             {
//                 path: '/logger/operate',
//                 element: <LoggerOperatePage />
//             },
//             {
//                 path: '/logger/login',
//                 element: <LoggerLoginPage />
//             },
//             {
//                 path: '/monitor/onlineUser',
//                 element: <MonitorOnlinePage />
//             },
//             {
//                 path: '/monitor/timeTask',
//                 element: <MonitorTaskPage />
//             },
//             {
//                 path: '/monitor/server',
//                 element: <MonitorServerPage />
//             },
//             {
//                 path: '/monitor/cache',
//                 element: <MonitorCachePage />
//             },
//             {
//                 path: '/monitor/cacheList',
//                 element: <MonitorCacheListPage />
//             },
//             {
//                 path: '/user/setting',
//                 element: <UserCenterPage />,
//             },
//             {
//                 path: '/403',
//                 element: <NotAuthPage />
//             },
//             {
//                 path: '/404',
//                 element: <NotFoundPage />
//             },
//             {
//                 path: '/500',
//                 element: <ServerErrorPage />
//             },
//             {
//                 path: '/*',
//                 element: <NotFoundPage />
//             },
//         ]
//     }
// ]

const Router: React.FC = () => {

    const [routes, setCurrRoute] = useLoadRoutes();
    const routeResult = useRoutes(constantRouter.concat(routes));
    useEffect(() => {
        setCurrRoute(routeResult && routeResult.props ?
            routeResult.props.match.route as IRouteObject :
            NOT_FOUND_PAGE)
    }, [routeResult]);
    return routeResult;
}

export default Router;

// export const SystemRouter = (): React.ReactElement | null => {
//     // const user = useSelector((state) => state.user)
//     return useRoutes(constantRouter as RouteObject[]);
// }