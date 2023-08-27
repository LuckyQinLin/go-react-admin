import { useEffect, useState } from "react"
import {IRouteObject} from "@/router/modules.ts";
import {asyncRoutes} from "@/router/index.tsx";
import {useLocation, useNavigate} from "react-router-dom";
import {useSelector} from "@/redux/hooks.ts";
import {useRequest} from "ahooks";
import {userLoginInfo} from "@/api/user.ts";
import {changeLoginStatusActionCreator} from "@/redux/user/action.ts";
import {useDispatch} from "react-redux";
import {userRouter} from "@/api/menu.ts";

const useLoadRoutes = (): [
    router: IRouteObject[],
    setCurrRoute: (data: IRouteObject | undefined | string) => void,
    getUserInfo: () => void,
] => {

    const location = useLocation();
    const navigate = useNavigate();
    const dispatch = useDispatch();
    const {token } = useSelector((state) => state.user);
    const [routes, setRoutes] = useState<IRouteObject[]>([]);
    const [currRoute, setCurrRoute] = useState<IRouteObject | undefined | string>(undefined);
    // 加载用户信息
    const loadInfo = useRequest(userLoginInfo, {
        manual: true,
        onSuccess: (data) => {
            dispatch(changeLoginStatusActionCreator({...data}));
        }
    });

    // 加载当前用户的路由信息
    const loadRouter = useRequest(userRouter, {
        manual: true,
        onSuccess: (data) => {
            // 生成路由和菜单信息
            console.log("菜单：", data);
        }
    })

    const getUserInfo = () => {
        loadInfo.run();
        loadRouter.run();
    }


    useEffect(() => {
        console.log("location", location.pathname, location.state, location.search, location.key)
        if (!token && location.pathname != '/login') {
            console.log('不存在登录信息')
            navigate('/login');
        }
    }, [location.pathname]);

    useEffect(() => {
        console.log("当前的路由 => ", currRoute)
        if (currRoute === undefined) {
            // navigate(LOGIN_PAGE);
        } else if (currRoute instanceof String) {
            // 返回string,将在页面404时进行处理
            navigate(currRoute as string);
        } else {
            // 设置tab和面包屑

        }

    }, [currRoute]);

    // 用来获取用户权限信息
    useEffect(() => {
        console.log('获取用户路由信息......')
        setRoutes([...asyncRoutes])
    }, []);

    return [routes, setCurrRoute, getUserInfo]

}

export default useLoadRoutes