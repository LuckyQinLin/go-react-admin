import {IRouteObject} from "@/router/modules.ts";
import lazyLoad from "@/router/lazyLoad.tsx";

const ExceptionRouter: IRouteObject[] = [
    {
        path: 'exception',
        element: lazyLoad(() => import("@/pages/layout")),
        children: [
            {
                path: '403',
                element: lazyLoad(() => import("@/pages/exception/403.tsx")),
            },
            {
                path: '404',
                element: lazyLoad(() => import("@/pages/exception/404.tsx")),
            },
            {
                path: '500',
                element: lazyLoad(() => import("@/pages/exception/500.tsx")),
            },
        ]
    }
];

export default ExceptionRouter;