import {useRoutes} from "react-router-dom";
import HomeRouter from "@/router/modules/home.tsx";
import ExceptionRouter from "@/router/modules/exception.tsx";
import PersonRouter from "@/router/modules/person.tsx";
import SettingRouter from "@/router/modules/setting.tsx";
import LoggerRouter from "@/router/modules/logger.tsx";
import MonitorRouter from "@/router/modules/monitor.tsx";
import LoginRouter from "@/router/modules/login.tsx";
import React from "react";
import useLoadRoutes from "@/router/useLoadRoutes.tsx";

// 无需验证的普通路由
export const constantRouter = [
    ...LoginRouter,
    ...HomeRouter,
    ...ExceptionRouter,
    ...PersonRouter,
]

// 需要验证的路由
export const asyncRoutes = [...SettingRouter, ...LoggerRouter, ...MonitorRouter]

const Router: React.FC = () => {

    const [routes] = useLoadRoutes();
    return useRoutes(routes);
}

export default Router;
