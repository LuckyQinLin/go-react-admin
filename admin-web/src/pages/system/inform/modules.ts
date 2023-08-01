
// NoticePageQueryProp 通知表查询
interface NoticePageQueryProp {
    page: number;
    size: number;
    noticeTitle?: string; // 通知标题
    noticeType?: string; // 通知类型
    userName?: string; // 操作人
}


// NoticePageProp 通知表属性
interface NoticePageProp {
    noticeId: number;      // 通知ID
    noticeTitle: string;    // 通知名称
    noticeType: number;    // 通知类型
    createTime: string;  // 创建时间
    createBy: string; // 创建人
    status: number;      // 状态
}

// NoticeDrawerProp 通知抽屉属性
interface NoticeDrawerProp {
    createVisible: boolean;
    updateVisible: boolean;
    noticeId?: number;
}

// NoticeCreateDrawerProp 通知创建属性
interface NoticeCreateDrawerProp {
    visible: boolean;
    close: (isLoad: boolean) => void;
}

// NoticeCreateFormProp 通知创建表单属性
interface NoticeCreateFormProp {
    noticeTitle: string;    // 通知名称
    noticeType: number;    // 通知编码  1通知 2:公告
    noticeContent: string; // 通知内容
    status: number;      // 状态
}

// NoticeUpdateDrawerProp 通知更新属性
interface NoticeUpdateDrawerProp {
    visible: boolean;
    noticeId?: number;
    close: (isLoad: boolean) => void;
}

// NoticeUpdateFormProp 通知修改
interface NoticeUpdateFormProp extends NoticeCreateFormProp {
    noticeId: number; // 通知ID
}


export type {
    NoticePageProp,
    NoticePageQueryProp,
    NoticeDrawerProp,
    NoticeCreateDrawerProp,
    NoticeCreateFormProp,
    NoticeUpdateDrawerProp,
    NoticeUpdateFormProp
}