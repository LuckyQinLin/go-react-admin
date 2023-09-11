import {IRouteObject, RouterMap} from "@/router/modules.ts";
import {UserRouterProp} from "@/pages/system/menu/modules.ts";
import IconFont from "@/components/IconFont";
import {Link} from "react-router-dom";
import type { MenuProps } from 'antd';
import React from "react";

type MenuItem = Required<MenuProps>['items'][number];

export const routerBuild = (data: UserRouterProp[]): IRouteObject[] => {
    return data.map((item) => {
        const currentRouter: IRouteObject = {
            path: `${item.path}`,
            element: RouterMap.get(item.component),
            meta: {
                isRoot: item.isRoot ? item.isRoot : false,
                key: item.perms,
                sort: item.sort,
                title: item.title,
                icon: item.icon,
                permission: [item.perms],
            },
        };
        // 是否有子菜单，并递归处理
        if (item.children && item.children.length > 0) {
            currentRouter.children = routerBuild(item.children);
        }
        return currentRouter;
    }).sort((a, b) => (a.meta?.sort ? a.meta.sort : 0) - (b.meta?.sort ? b.meta.sort : 0));
}

/**
 * 判断根路由 Router
 * */
const isRootRouter = (item: IRouteObject): boolean => item.meta?.isRoot || (item.children?.length === 1);

const existChildren = (item: IRouteObject): boolean => item.children ? item.children.length > 0 : false;

const buildMenuItem = (label: React.ReactNode, key: React.Key, icon?: React.ReactNode, children?: MenuItem[]): MenuItem => {
    return { key, icon, children, label } as MenuItem;
}

export const routerBuildMenu = (data: IRouteObject[]): MenuItem[] => {
    return data.map((item) => {
        const isRoot = isRootRouter(item);
        if (isRoot) {
            const info = isRoot && item.children ? item.children[0] : item;
            const key = info.meta?.key ? info.meta?.key : info.path;
            const title = <Link to={`${info.path}`}>{info.meta?.title}</Link>;
            const icon = info.meta?.icon ? <IconFont type={info.meta.icon} /> : undefined;
            return buildMenuItem(
                title,
                key,
                icon,
            )
        } else {
            const label = existChildren(item) ? item.meta?.title! : <Link to={`${item.path}`}>{item.meta?.title}</Link>
            const key = item.meta?.key ? item.meta?.key : item.path;
            const icon = item.meta?.icon ? <IconFont type={item.meta.icon} /> : undefined;
            const children = existChildren(item) ? routerBuildMenu(item.children!) : undefined;
            return buildMenuItem(
                label,
                key,
                icon,
                children
            )
        }
    });
}

export const existRouter = (data: IRouteObject[], path: string): boolean => {

    const isExist = (item: IRouteObject, path: string): boolean => {
        if (item.path === path) {
            return true;
        }
        if (item.children) {
            for (const node of item.children) {
                if (isExist(node, path)) {
                    return true;
                }
            }
        }
        return false;
    }

    return data.some(item => isExist(item, path))
}