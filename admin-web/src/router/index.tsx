import {useRoutes} from "react-router-dom";
import HomeRouter from "@/router/modules/home.tsx";
import ExceptionRouter from "@/router/modules/exception.tsx";
import PersonRouter from "@/router/modules/person.tsx";
import SettingRouter from "@/router/modules/setting.tsx";
import LoggerRouter from "@/router/modules/logger.tsx";
import MonitorRouter from "@/router/modules/monitor.tsx";
import LoginRouter from "@/router/modules/login.tsx";
import React, {useEffect, useMemo} from "react";
import {useDispatch} from "react-redux";
import {useSelector} from "@/redux/hooks.ts";
import {IRouteObject} from "@/router/modules.ts";
import {useRequest} from "ahooks";
import {userLoginInfo} from "@/api/user.ts";
import {changeLoginStatusActionCreator, changeMenuStatusActionCreator} from "@/redux/user/action.ts";
import {userPageRouter} from "@/api/menu.ts";
import {routerBuild} from "@/router/routerFilter.tsx";
import { Navigate } from 'react-router-dom';

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

    const dispatch = useDispatch();
    const {token, permissions, userRouter} = useSelector((state) => state.user);

    // 加载用户信息
    const loadInfo = useRequest(userLoginInfo, {
        manual: true,
        onSuccess: (data) => {
            dispatch(changeLoginStatusActionCreator({...data}));
        }
    });

    // 加载当前用户的路由信息
    const loadRouter = useRequest(userPageRouter, {
        manual: true,
        onSuccess: (data) => {
            dispatch(changeMenuStatusActionCreator({userRouter: data}));
        }
    })

    const getUserInfo = () => {
        loadInfo.run();
        loadRouter.run();
    }

    // if (token) {
    //     getUserInfo();
    // }


    const routes = useMemo(() => {
        let router: IRouteObject[] = permissions?.length === 1 && permissions[0] === '*:*:*' ? asyncRoutes : (userRouter ? routerBuild(userRouter) : constantRouter);
        return [...constantRouter, ...router, {path: '*', element: <Navigate to="/exception/404" />}]
    }, [userRouter])

    useEffect(() => {
        if (token) {
            getUserInfo();
        }
    }, []);



    return useRoutes(routes);
}

export default Router;
