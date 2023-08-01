
// DictPageQueryProp 岗位表查询
interface DictPageQueryProp {
    page: number;
    size: number;
    dictName?: string;
    dictType?: string;
    status?: number;
}


// DictPageProp 岗位表属性
interface DictPageProp {
    dictId: number;      // 岗位ID
    dictName: string;    // 岗位名称
    dictType: string;    // 岗位编码
    createTime: string;  // 创建时间
    remark?: string; // 备注
    status: number;      // 状态
}

// DictDrawerProp 岗位抽屉属性
interface DictDrawerProp {
    createVisible: boolean;
    updateVisible: boolean;
    dictId?: number;
}

// DictCreateDrawerProp 岗位创建属性
interface DictCreateDrawerProp {
    visible: boolean;
    close: (isLoad: boolean) => void;
}

// DictCreateFormProp 岗位创建表单属性
interface DictCreateFormProp {
    dictName: string;    // 岗位名称
    dictType: string;    // 岗位编码
    status: number;      // 状态
    remark?: string; // 备注
}

// DictUpdateDrawerProp 岗位更新属性
interface DictUpdateDrawerProp {
    visible: boolean;
    dictId?: number;
    close: (isLoad: boolean) => void;
}

// DictUpdateFormProp 岗位修改
interface DictUpdateFormProp extends DictCreateFormProp {
    dictId: number; // 岗位ID
}


export type {
    DictPageProp,
    DictPageQueryProp,
    DictDrawerProp,
    DictCreateDrawerProp,
    DictCreateFormProp,
    DictUpdateDrawerProp,
    DictUpdateFormProp
}