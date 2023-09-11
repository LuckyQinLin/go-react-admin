
// PostPageQueryProp 岗位表查询
interface PostPageQueryProp {
    page: number;
    size: number;
    postName?: string;
    postCode?: string;
    status?: number;
}

// PostInfoProp 岗位信息
interface PostInfoProp {
    value: number;
    label: string;
}

// PostPageProp 岗位表属性
interface PostPageProp {
    postId: number;      // 岗位ID
    postName: string;    // 岗位名称
    postCode: string;    // 岗位编码
    postSort: number;    // 岗位排序
    createTime: string;  // 创建时间
    status: number;      // 状态
}

// PostDrawerProp 岗位抽屉属性
interface PostDrawerProp {
    createVisible: boolean;
    updateVisible: boolean;
    postId?: number;
}

// PostCreateDrawerProp 岗位创建属性
interface PostCreateDrawerProp {
    visible: boolean;
    close: (isLoad: boolean) => void;
}

// PostCreateFormProp 岗位创建表单属性
interface PostCreateFormProp {
    postName: string;    // 岗位名称
    postCode: string;    // 岗位编码
    postSort: number;    // 岗位排序
    status: number;      // 状态
    remark?: string; // 备注
}

// PostUpdateDrawerProp 岗位更新属性
interface PostUpdateDrawerProp {
    visible: boolean;
    postId?: number;
    close: (isLoad: boolean) => void;
}

// PostUpdateFormProp 岗位修改
interface PostUpdateFormProp extends PostCreateFormProp {
    postId: number; // 岗位ID
}


export type {
    PostInfoProp,
    PostPageProp,
    PostPageQueryProp,
    PostDrawerProp,
    PostCreateDrawerProp,
    PostCreateFormProp,
    PostUpdateDrawerProp,
    PostUpdateFormProp
}