import {IRouteObject, RouterMap} from "@/router/modules.ts";

// 首页
const PersonRouter: IRouteObject[] = [
    {
        path: '/person',
        redirect: '/person/index',
        element: RouterMap.get('Layout'),
        meta: {
            sort: 5,
            isRoot: true,
            icon: "lucky-jiankong",
        },
        children: [
            {
                path: '/person/index',
                element: RouterMap.get('PersonUserPage'),
                meta: {
                    sort: 1,
                    title: '个人中心',
                    icon: "lucky-jiankong",
                },
            }
        ]
    }
]

export default PersonRouter;