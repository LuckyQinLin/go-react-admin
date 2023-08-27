import {IRouteObject} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";
import IconFont from "@/components/IconFont";

// 首页
const PersonRouter: IRouteObject[] = [
    {
        path: 'person',
        redirect: '/person/index',
        element: lazyLoad(() => import("@/pages/layout")),
        meta: {
            sort: 5,
            isRoot: true,
            icon: <IconFont type="lucky-jiankong" />,
        },
        children: [
            {
                path: 'index',
                element: lazyLoad(() => import("@/pages/user")),
                meta: {
                    title: '个人中心',
                },
            }
        ]
    }
]

export default PersonRouter;