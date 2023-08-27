import {IRouteObject} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";
import IconFont from "@/components/IconFont";

// 首页
const HomeRouter: IRouteObject[] = [
    {
        path: 'home',
        redirect: '/home/index',
        element: lazyLoad(() => import("@/pages/layout")),
        meta: {
            icon: <IconFont type="lucky-shouye1" />,
            isRoot: true,
            sort: 1,
        },
        children: [
            {
                path: 'index',
                element: lazyLoad(() => import("@/pages/home")),
                meta: {
                    title: '首页',
                },
            }
        ]
    }
]

export default HomeRouter;