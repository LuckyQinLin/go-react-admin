import {IRouteObject, RouterMap} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";

const LoggerRouter: IRouteObject[] = [
    {
        path: '/logger',
        element: RouterMap.get('Layout'),
        meta: {
            icon: "lucky-nav_icon_rzgl_spe",
            sort: 3,
        },
        children: [
            {
                path: '/logger/operate',
                element: lazyLoad(() => import("@/pages/logger/operate")),
                meta: {
                    sort: 1,
                    permission: ['logger:operate'],
                    icon: "lucky-caozuorizhi",
                    title: '操作日志',
                },
            },
            {
                path: '/logger/login',
                element: lazyLoad(() => import("@/pages/logger/login")),
                meta: {
                    sort: 2,
                    permission: ['logger:login'],
                    icon: "lucky-denglurizhi",
                    title: '登录日志',
                },
            }
        ]
    }
];

export default LoggerRouter;