// DeptTreeProp 部门树
interface DeptTreeProp {
    key: number;
    title: string | React.ReactElement;
    children: DeptTreeProp[];
}

// MenuTableTreeProp 部门表格树属性
interface DeptTableTreeProp {
    key: number;                     // 主键
    title: string;                   // 部门名称
    order: number;                   // 排序
    status: number;                  // 状态
    createTime: string;              // 创建时间
    children: DeptTableTreeProp[];   // 下级部门
}

// DeptTableQueryProp 部门表格查询
interface DeptTableQueryProp {
    name?: string; // 部门名称
    status?: number; // 部门状态
}

// DrawerProp 抽屉参数
interface DeptDrawerProp {
    currId?: number;
    parentId?: number;
    createVisible: boolean;
    updateVisible: boolean;
}

// DeptCreateDrawerProp 部门创建抽屉属性
interface DeptCreateDrawerProp {
    visible: boolean;
    parentId: number;
    close: (isLoad: boolean) => void;
}

// DeptUpdateDrawerProp 部门更新抽屉属性
interface DeptUpdateDrawerProp {
    visible: boolean;
    deptId?: number;
    close: (isLoad: boolean) => void;
}

// DeptCreateFormProp 部门创建表单
interface DeptCreateFormProp {
    parentId: number; // 上级部门ID
    deptName: string; // 部门名称
    orderNum: number; // 显示顺序
    leader: string;   // 负责人
    phone: string;    // 手机
    email: string;    // 邮箱
    status: number;   // 状态
}

// DeptUpdateFormProp 部门更新表单
interface DeptUpdateFormProp extends DeptCreateFormProp {
    deptId: number; // 修改ID
}

export type {
    DeptTreeProp,
    DeptTableTreeProp,
    DeptTableQueryProp,
    DeptDrawerProp,
    DeptCreateDrawerProp,
    DeptUpdateDrawerProp,
    DeptCreateFormProp,
    DeptUpdateFormProp
}