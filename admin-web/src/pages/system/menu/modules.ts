// MenuTreeProp 菜单树属性
interface MenuTreeProp {
    key: number;
    title: string;
    children: MenuTreeProp[];
}

export type {
    MenuTreeProp
}