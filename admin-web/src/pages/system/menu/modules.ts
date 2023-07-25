// MenuTreeProp 菜单树属性
interface MenuTreeProp {
    key: number;
    title: string;
    children: MenuTreeProp[];
}

// MenuTableTreeProp 菜单表格树属性
interface MenuTableTreeProp {
    id: number;       // 主键
    title: string;    // 菜单名称
    code: string;     // 权限字符
    icon: string;     // 图标
    path?: string;    // 路由
    parentId: number; // 上级ID
    status: number;   // 状态
    order: number;    // 排序
    createTime: string; // 创建时间
}

// DrawerProp 抽屉参数
interface DrawerProp {
    types?: number;
    currId?: number;
    parentId?: number;
    createVisible: boolean;
    updateVisible: boolean;
}

export type {
    MenuTreeProp,
    DrawerProp,
    MenuTableTreeProp
}
