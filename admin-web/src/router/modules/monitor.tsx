import {IRouteObject} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";
import IconFont from "@/components/IconFont";

const MonitorRouter: IRouteObject[] = [
    {
        path: 'monitor',
        element: lazyLoad(() => import("@/pages/layout")),
        meta: {
            sort: 4,
            title: '系统监控',
            permission: ['monitor'],
            icon: <IconFont type="lucky-jiankong" />,
        },
        children: [
            {
                path: 'onlineUser',
                element: lazyLoad(() => import("@/pages/monitor/online")),
                meta: {
                    sort: 1,
                    title: '在线用户',
                    permission: ['monitor:onlineUser'],
                    icon: <IconFont type="lucky-zaixianyonghuguanli1" />,
                },
            },
            {
                path: 'timeTask',
                element: lazyLoad(() => import("@/pages/monitor/task")),
                meta: {
                    sort: 2,
                    title: '定时任务',
                    permission: ['monitor:timeTask'],
                    icon: <IconFont type="lucky-dingshirenwuguanli" />,
                },
            },
            {
                path: 'server',
                element: lazyLoad(() => import("@/pages/monitor/server")),
                meta: {
                    sort: 3,
                    title: '服务器监控',
                    permission: ['monitor:server'],
                    icon: <IconFont type="lucky-fuwuqijiankong" />,
                },
            },
            {
                path: 'cache',
                element: lazyLoad(() => import("@/pages/monitor/cache")),
                meta: {
                    sort: 4,
                    title: '缓冲监控',
                    permission: ['monitor:cache'],
                    icon: <IconFont type="lucky-huanchongfenxi" />,
                },
            },
            {
                path: 'cacheList',
                element: lazyLoad(() => import("@/pages/monitor/cacheList")),
                meta: {
                    sort: 5,
                    title: '缓冲列表',
                    permission: ['monitor:cacheList'],
                    icon: <IconFont type="lucky-cityworksjichugongnengtubiao-" />,
                },
            }
        ]
    }
];

export default MonitorRouter;