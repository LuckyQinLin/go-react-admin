import {IRouteObject, RouterMap} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";

const SettingRouter: IRouteObject[] = [
    {
        path: '/setting',
        element: RouterMap.get('Layout'),
        meta: {
            sort: 2,
            title: '系统设置',
            permission: ['system'],
            icon: "lucky-shezhi1"
        },
        children: [
            {
                path: '/setting/user',
                element: lazyLoad(() => import("@/pages/system/user")),
                meta: {
                    sort: 1,
                    title: '用户管理',
                    permission: ['system:user'],
                    icon: "lucky-yonghu"
                },
            },
            {
                path: '/setting/role',
                element: lazyLoad(() => import("@/pages/system/role")),
                meta: {
                    sort: 2,
                    icon: "lucky-jiaose",
                    title: '角色管理',
                    permission: ['system:role']
                },
            },
            {
                path: '/setting/menu',
                element: lazyLoad(() => import("@/pages/system/menu")),
                meta: {
                    sort: 3,
                    icon: "lucky-caidan",
                    title: '菜单管理',
                    permission: ['system:menu']
                },
            },
            {
                path: '/setting/dept',
                element: lazyLoad(() => import("@/pages/system/dept")),
                meta: {
                    sort: 4,
                    icon: "lucky-bumenguanli",
                    title: '部门管理',
                    permission: ['system:dept']
                },
            },
            {
                path: '/setting/post',
                element: lazyLoad(() => import("@/pages/system/post")),
                meta: {
                    sort: 5,
                    icon: "lucky-gangwei",
                    title: '岗位管理',
                    permission: ['system:post']
                },
            },
            {
                path: '/setting/dict',
                element: lazyLoad(() => import("@/pages/system/dict")),
                meta: {
                    sort: 6,
                    icon: "lucky-zidianmuluguanli",
                    title: '字典管理',
                    permission: ['system:dict']
                },
            },
            {
                path: '/setting/param',
                element: lazyLoad(() => import("@/pages/system/param")),
                meta: {
                    sort: 7,
                    icon: "lucky-shujucanshu",
                    title: '参数管理',
                    permission: ['system:param']
                },
            },
            {
                path: '/setting/inform',
                element: lazyLoad(() => import("@/pages/system/inform")),
                meta: {
                    sort: 8,
                    icon: "lucky-tongzhi",
                    title: '通知公告',
                    permission: ['system:inform']
                },
            }
        ]
    }
];

export default SettingRouter;