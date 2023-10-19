import {Navigate} from "react-router-dom";
import Router from "@/router/modules.tsx";
import getRouter = Router.getRouter;
import RouterEnum = Router.RouterEnum;
import {User, Menus} from "@/types";
import {getPermissions} from "@/api/menu.ts";

namespace RouterSpace {
    const authLoader = async (): Promise<User.UserPermissionProp> => {
        const data = await getPermissions();
        const [paths, buttons] = parseMenus(data)
        return {
            menus: data,
            buttons: buttons,
            paths: paths
        }
    }

    /**
     * 解析资源信息返回路由和按钮
     * @param data
     */
    const parseMenus = (data: Menus.MenuItemProp[]): [string[], string[]] => {
        // 获取菜单路由
        const getMenuPath = (list: Menus.MenuItemProp[]): string[] => {
            return list.reduce((result: string[], item: Menus.MenuItemProp) => {
                return result.concat(Array.isArray(item.children) && item.types !== 'F' ? getMenuPath(item.children) : item.path + '')
            }, [])
        }
        // 获取菜单按钮
        const getMenuButton = (list: Menus.MenuItemProp[]): string[] => {
            return list.reduce((result: string[], item: Menus.MenuItemProp) => {
                return result.concat(Array.isArray(item.children) && item.types === 'F' ? getMenuPath(item.children) : item.perms)
            }, [])
        }
        // 返回
        return [getMenuPath(data), getMenuButton(data)]
    }

    export const searchRoute: any = (path: string, routes: any = []) => {
        console.log("searchRoute")
        for (const item of routes) {
            if (item.path === path) return true
            if (item.children) {
                return searchRoute(path, item.children)
            }
        }
        return false
    }

    export const routers = [
        {
            path: '/login',
            element: getRouter(RouterEnum.LoginPage)
        },
        {
            id: Router.LayoutId,
            element: getRouter(RouterEnum.Layout),
            // element: LayoutPage,
            loader: authLoader, // 鉴权方法
            children: [
                {
                    path: '/home/index', // 首页
                    element: getRouter(RouterEnum.HomePage),
                },
                {
                    path: '/system/user', // 用户管理
                    element: getRouter(RouterEnum.SystemUserPage),
                },
                {
                    path: '/system/role', // 角色管理
                    element: getRouter(RouterEnum.SystemRolePage),
                },
                {
                    path: '/system/menu', // 菜单管理
                    element: getRouter(RouterEnum.SystemMenuPage),
                },
                {
                    path: '/system/dept', // 部门管理
                    element: getRouter(RouterEnum.SystemDeptPage),
                },
                {
                    path: '/system/post', // 岗位管理
                    element: getRouter(RouterEnum.SystemPostPage),
                },
                {
                    path: '/system/dict', // 字典管理
                    element: getRouter(RouterEnum.SystemDictPage),
                },
                {
                    path: '/system/param', // 参数管理
                    element: getRouter(RouterEnum.SystemParamPage),
                },
                {
                    path: '/system/inform', // 通知管理
                    element: getRouter(RouterEnum.SystemInformPage),
                },
                {
                    path: '/logger/operate', // 操作日志
                    element: getRouter(RouterEnum.LoggerOperatePage),
                },
                {
                    path: '/logger/login', // 登录日志
                    element: getRouter(RouterEnum.LoggerLoginPage),
                },
                {
                    path: '/monitor/onlineUser', // 在线用户
                    element: getRouter(RouterEnum.MonitorOnlinePage),
                },
                {
                    path: '/monitor/timeTask', // 定时任务
                    element: getRouter(RouterEnum.MonitorTaskPage),
                },
                {
                    path: '/monitor/server', // 服务器监控
                    element: getRouter(RouterEnum.MonitorServerPage),
                },
                {
                    path: '/monitor/cache', // 缓冲监控
                    element: getRouter(RouterEnum.MonitorCachePage),
                },
                {
                    path: '/monitor/cacheList', // 缓冲列表
                    element: getRouter(RouterEnum.MonitorCacheListPage),
                },
                {
                    path: '/person/index', // 缓冲列表
                    element: getRouter(RouterEnum.PersonUserPage),
                },
                {
                    path: '/exception/403', // 缓冲列表
                    element: getRouter(RouterEnum.NotAuthPage),
                },
                {
                    path: '/exception/404', // 缓冲列表
                    element: getRouter(RouterEnum.NotFoundPage),
                },
                {
                    path: '/exception/500', // 缓冲列表
                    element: getRouter(RouterEnum.ServerErrPage),
                },
                {
                    path: "*",
                    element: <Navigate to={'/exception/404'} />
                }
            ]

        },
        {
            path: "*",
            element: getRouter(RouterEnum.NotFoundPage),
        }
    ]
}




export default RouterSpace;
