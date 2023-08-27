import {IRouteObject} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";
import IconFont from "@/components/IconFont";

const SettingRouter: IRouteObject[] = [
    {
        path: 'setting',
        element: lazyLoad(() => import("@/pages/layout")),
        meta: {
            sort: 2,
            title: '系统设置',
            permission: ['system'],
            icon: <IconFont type="lucky-shezhi1" />
        },
        children: [
            {
                path: 'user',
                element: lazyLoad(() => import("@/pages/system/user")),
                meta: {
                    sort: 1,
                    title: '用户管理',
                    permission: ['system:user'],
                    icon: <IconFont type="lucky-yonghu" />
                },
            },
            {
                path: 'role',
                element: lazyLoad(() => import("@/pages/system/role")),
                meta: {
                    sort: 2,
                    icon: <IconFont type="lucky-jiaose" />,
                    title: '角色管理',
                    permission: ['system:role']
                },
            },
            {
                path: 'menu',
                element: lazyLoad(() => import("@/pages/system/menu")),
                meta: {
                    sort: 3,
                    icon: <IconFont type="lucky-caidan" />,
                    title: '菜单管理',
                    permission: ['system:menu']
                },
            },
            {
                path: 'dept',
                element: lazyLoad(() => import("@/pages/system/dept")),
                meta: {
                    sort: 4,
                    icon: <IconFont type="lucky-bumenguanli" />,
                    title: '部门管理',
                    permission: ['system:dept']
                },
            },
            {
                path: 'post',
                element: lazyLoad(() => import("@/pages/system/post")),
                meta: {
                    sort: 5,
                    icon: <IconFont type="lucky-gangwei" />,
                    title: '岗位管理',
                    permission: ['system:post']
                },
            },
            {
                path: 'dict',
                element: lazyLoad(() => import("@/pages/system/dict")),
                meta: {
                    sort: 6,
                    icon: <IconFont type="lucky-zidianmuluguanli" />,
                    title: '字典管理',
                    permission: ['system:dict']
                },
            },
            {
                path: 'param',
                element: lazyLoad(() => import("@/pages/system/param")),
                meta: {
                    sort: 7,
                    icon: <IconFont type="lucky-shujucanshu" />,
                    title: '参数管理',
                    permission: ['system:param']
                },
            },
            {
                path: 'inform',
                element: lazyLoad(() => import("@/pages/system/inform")),
                meta: {
                    sort: 8,
                    icon: <IconFont type="lucky-tongzhi" />,
                    title: '通知公告',
                    permission: ['system:inform']
                },
            }
        ]
    }
];

export default SettingRouter;