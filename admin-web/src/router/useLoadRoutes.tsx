import { useEffect, useState } from "react"
import {IRouteObject} from "@/router/modules.ts";
import {useLocation, useNavigate} from "react-router-dom";
import {useSelector} from "@/redux/hooks.ts";
import {useRequest} from "ahooks";
import {userLoginInfo} from "@/api/user.ts";
import {changeLoginStatusActionCreator, changeMenuStatusActionCreator} from "@/redux/user/action.ts";
import {useDispatch} from "react-redux";
import {userRouter} from "@/api/menu.ts";
import {existRouter, routerBuild, routerBuildMenu} from "@/router/routerFilter.tsx";
import PersonRouter from "@/router/modules/person.tsx";
import HomeRouter from "@/router/modules/home.tsx";
import {constantRouter} from "@/router/index.tsx";
import {LOGIN_PAGE, NOT_FOUND_PAGE} from "@/constant/setting.ts";

const useLoadRoutes = (): [
    router: IRouteObject[],
    getUserInfo: () => void,
] => {

    const location = useLocation();
    const navigate = useNavigate();
    const dispatch = useDispatch();
    const {token } = useSelector((state) => state.user);
    const [routes, setRoutes] = useState<IRouteObject[]>(constantRouter);

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
            console.log("菜单-data：", data);
            const router = routerBuild(data);
            console.log("菜单-router-1：", router);
            console.log("菜单-router-2：", [...HomeRouter, ...router, ...PersonRouter]);
            const menus = routerBuildMenu([...HomeRouter, ...router, ...PersonRouter]);
            console.log("菜单-menu：", menus);
            setRoutes([...routes, ...router]); // 生成路由
            // setRoutes([...routes]); // 生成路由
            dispatch(changeMenuStatusActionCreator({menus: menus}));
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
            navigate(LOGIN_PAGE);
            return;
        }
        debugger;
        if (!existRouter(routes, location.pathname)) {
            navigate(NOT_FOUND_PAGE);
            return;
        }

    }, [location.pathname]);


    // 用来获取用户权限信息
    useEffect(() => {
        console.log('获取用户路由信息......')
        if (token) {
            getUserInfo();
        }
    }, []);

    return [routes, getUserInfo]

}

export default useLoadRoutes