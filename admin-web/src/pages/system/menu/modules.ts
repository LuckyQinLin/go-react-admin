// MenuTreeProp 菜单树属性
interface MenuTreeProp {
    key: number;
    title: string;
    children: MenuTreeProp[];
}

// MenuTableTreeQueryProp 菜单表格查询
interface MenuTableTreeQueryProp {
    name?: string;
    status?: number;
}

// MenuTableTreeProp 菜单表格树属性
interface MenuTableTreeProp {
    key: number;                     // 主键
    title: string;                   // 菜单名称
    code: string;                    // 权限字符
    icon: string;                    // 图标
    path?: string;                   // 路由
    parentId: number;                // 上级ID
    status: number;                  // 状态
    order: number;                   // 排序
    createTime: string;              // 创建时间
    children: MenuTableTreeProp[];
}

// DrawerProp 抽屉参数
interface DrawerProp {
    types?: number;
    currId?: number;
    parentId?: number;
    createVisible: boolean;
    updateVisible: boolean;
}

// MenuCreateDrawerProp 菜单创建抽屉属性
interface MenuCreateDrawerProp {
    visible: boolean;
    parentId: number;
    close: (isLoad: boolean) => void;
}

// MenuUpdateDrawerProp 菜单更新抽屉属性
interface MenuUpdateDrawerProp {
    visible: boolean;
    menuId?: number;
    close: (isLoad: boolean) => void;
}

// MenuCreateFormProp 菜单创建表单属性
interface MenuCreateFormProp {
    parentId:  number;  // 上级菜单
    menuType:  string;  // 菜单类型 (M目录 C菜单 F按钮)
    perms: string; // 权限字符
    icon:      string;  // 图标
    menuName:  string;  // 菜单名称
    menuSort:  number;  // 菜单排序
    path:      string;  // 路由地址
    component: string;  // 组件路径
    param:     string;  // 路由参数
    isLink:    boolean; // 是否外链 true:是 false:不是
    isShow:    boolean; // 显示状态 true:是 false:不是
    isCache:   boolean; // 缓冲 true:是 false:不是
    status:    number;  // 菜单状态 1:正常 0:停用
}

// MenuUpdateFormProp 菜单创建表单属性
interface MenuUpdateFormProp extends MenuCreateFormProp{
    menuId: number; // 菜单ID
}

export type {
    MenuTreeProp,
    DrawerProp,
    MenuTableTreeProp,
    MenuTableTreeQueryProp,
    MenuCreateDrawerProp,
    MenuUpdateDrawerProp,
    MenuCreateFormProp,
    MenuUpdateFormProp
}
