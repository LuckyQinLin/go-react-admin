import {IRouteObject} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";

// 首页
const LoginRouter: IRouteObject[] = [
    {
        path: 'login',
        element: lazyLoad(() => import("@/pages/login")),
    }
]

export default LoginRouter;