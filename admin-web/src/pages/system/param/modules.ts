// ConfigPageQueryProp 参数表查询
interface ConfigPageQueryProp {
    page: number;
    size: number;
    configName?: string;
    configKey?: string;
    configType?: number;
}


// ConfigPageProp 参数表属性
interface ConfigPageProp {
    configId: number;      // 参数ID
    configName: string;    // 参数名称
    configKey: string;    // 参数名称
    configValue: string;    // 参数名称
    configType: number;    // 系统内置 1是 0否
    createTime: string;  // 创建时间
    remark: string; // 备注
}

// ConfigDrawerProp 参数抽屉属性
interface ConfigDrawerProp {
    createVisible: boolean;
    updateVisible: boolean;
    configId?: number;
}

// ConfigCreateDrawerProp 参数创建属性
interface ConfigCreateDrawerProp {
    visible: boolean;
    close: (isLoad: boolean) => void;
}

// ConfigCreateFormProp 参数创建表单属性
interface ConfigCreateFormProp {
    configName: string;    // 参数名称
    configKey: string;    // 参数键名
    configValue: string;    // 参数键值
    configType: number;    // 系统内置 1是 0否
    remark?: string; // 备注
}

// ConfigUpdateDrawerProp 参数更新属性
interface ConfigUpdateDrawerProp {
    visible: boolean;
    configId?: number;
    close: (isLoad: boolean) => void;
}

// ConfigUpdateFormProp 参数修改
interface ConfigUpdateFormProp extends ConfigCreateFormProp {
    configId: number; // 参数ID
}


export type {
    ConfigPageProp,
    ConfigPageQueryProp,
    ConfigDrawerProp,
    ConfigCreateDrawerProp,
    ConfigCreateFormProp,
    ConfigUpdateDrawerProp,
    ConfigUpdateFormProp
}