import React from "react";

namespace Menus {
    // 资源数据
    export interface MenuItemProp {
        id:        number;        // 菜单主键
        pId:       number;        // 上级菜单
        icon?:     string;        // 图标
        types:     string;        // 类型 M目录 C菜单 F按钮
        title:     string;        // 菜单名称
        perms:     string;        // 权限字符
        sort:      number;        // 显示顺序
        path?:     string;        // 路由地址
        component: string;        // 组件路由
        children?: MenuItemProp[] // 下级菜单
    }

    // MenuTitleProp 菜单对应的标题和路由
    export interface MenuTitleProp {
        title: string;
        path: string;
    }

    // TabViewProp 路由菜单栏属性
    export interface TabViewProp {
        key: string | number;
        title: string | React.ReactNode;
        closeIcon?: boolean | React.ReactNode
    }
}

export default Menus
