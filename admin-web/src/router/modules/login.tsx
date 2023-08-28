import {IRouteObject, RouterMap} from "@/router/modules.ts";

// 首页
const LoginRouter: IRouteObject[] = [
    {
        path: '/login',
        element: RouterMap.get('Layout'),
    }
]

export default LoginRouter;