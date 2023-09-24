// LoginLoggerQueryProp 登录日志查询属性
import {Dayjs} from "dayjs";

interface LoginLoggerQueryProp {
    page:       number;  //
    size:       number;  //
    userName?:  string;  // 用户名称
    status?:    number;  // 访问结果
    startTime?: string;  // 开始时间
    endTime?:   string;  // 结束时间
    address?:   string;  // IP地址
    times?: (Dayjs | null)[];
}

// OperateLoggerQueryProp 操作日志查询属性
interface OperateLoggerQueryProp {
    page:          number;
    size:          number;
    userName?:     string; // 用户名称
    status?:       string; // 访问结果
    startTime?:    string; // 开始时间
    endTime?:      string; // 结束时间
    operateType?:  number; // 操作类型
    businessType?: number; // 业务类型
}

// LoginLoggerTableProp 登录日志表格属性
interface LoginLoggerTableProp {
    id:        number; // 主键
    userName:  string; // 用户名称
    ip:        string; // 登录地址
    address:   string; // 登录地点
    browser:   string; // 浏览器
    os:        string; // 操作系统
    status:    number; // 登录结果
    msg:       string; // 登录信息
    loginTime: string; // 登录时间
}

// OperateLoggerTableProp 操作日志表格属性
interface OperateLoggerTableProp {
    id:           number; // 主键
    title:        string; // 模块名称
    businessType: number; // 操作类型名称
    operatorType: number; // 操作类型名称
    ip:           string; // IP地址
    address:      string; // 操作地址
    status:       number; // 操作结果
    operTime:     string; // 操作时间
    costTime:     number; // 耗时
}

export type {
    LoginLoggerQueryProp,
    OperateLoggerQueryProp,
    LoginLoggerTableProp,
    OperateLoggerTableProp
}