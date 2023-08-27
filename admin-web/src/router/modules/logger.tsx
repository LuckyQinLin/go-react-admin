import {IRouteObject} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";
import IconFont from "@/components/IconFont";

const LoggerRouter: IRouteObject[] = [
    {
        path: 'logger',
        element: lazyLoad(() => import("@/pages/layout")),
        meta: {
            icon: <IconFont type="lucky-nav_icon_rzgl_spe" />,
            sort: 3,
        },
        children: [
            {
                path: 'operate',
                element: lazyLoad(() => import("@/pages/logger/operate")),
                meta: {
                    sort: 1,
                    permission: ['logger:operate'],
                    icon: <IconFont type="lucky-caozuorizhi" />,
                    title: '操作日志',
                },
            },
            {
                path: 'login',
                element: lazyLoad(() => import("@/pages/logger/login")),
                meta: {
                    sort: 2,
                    permission: ['logger:login'],
                    icon: <IconFont type="lucky-denglurizhi" />,
                    title: '登录日志',
                },
            }
        ]
    }
];

export default LoggerRouter;