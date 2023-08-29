import {IRouteObject, RouterMap} from "@/router/modules.ts";

// 首页
const HomeRouter: IRouteObject[] = [
    {
        path: '/home',
        redirect: '/home/index',
        element: RouterMap.get('Layout'),
        meta: {
            key: 'home',
            icon: "lucky-shouye1",
            isRoot: true,
            sort: 1,
        },
        children: [
            {
                path: '/home/index',
                element: RouterMap.get('HomePage'),
                meta: {
                    key: 'home:index',
                    title: '首页',
                    icon: "lucky-shouye1",
                    sort: 1,
                },
            }
        ]
    }
]

export default HomeRouter;