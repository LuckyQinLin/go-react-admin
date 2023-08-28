import {IRouteObject, RouterMap} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";

const MonitorRouter: IRouteObject[] = [
    {
        path: '/monitor',
        element: RouterMap.get('Layout'),
        meta: {
            sort: 4,
            title: '系统监控',
            permission: ['monitor'],
            icon: "lucky-jiankong",
        },
        children: [
            {
                path: 'onlineUser',
                element: lazyLoad(() => import("@/pages/monitor/online")),
                meta: {
                    sort: 1,
                    title: '在线用户',
                    permission: ['monitor:onlineUser'],
                    icon: "lucky-zaixianyonghuguanli1",
                },
            },
            {
                path: 'timeTask',
                element: lazyLoad(() => import("@/pages/monitor/task")),
                meta: {
                    sort: 2,
                    title: '定时任务',
                    permission: ['monitor:timeTask'],
                    icon: "lucky-dingshirenwuguanli",
                },
            },
            {
                path: 'server',
                element: lazyLoad(() => import("@/pages/monitor/server")),
                meta: {
                    sort: 3,
                    title: '服务器监控',
                    permission: ['monitor:server'],
                    icon: "lucky-fuwuqijiankong",
                },
            },
            {
                path: 'cache',
                element: lazyLoad(() => import("@/pages/monitor/cache")),
                meta: {
                    sort: 4,
                    title: '缓冲监控',
                    permission: ['monitor:cache'],
                    icon: "lucky-huanchongfenxi",
                },
            },
            {
                path: 'cacheList',
                element: lazyLoad(() => import("@/pages/monitor/cacheList")),
                meta: {
                    sort: 5,
                    title: '缓冲列表',
                    permission: ['monitor:cacheList'],
                    icon: "lucky-cityworksjichugongnengtubiao-",
                },
            }
        ]
    }
];

export default MonitorRouter;