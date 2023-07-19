import {Navigate, useRoutes} from "react-router-dom";
import {useSelector} from "@/redux/hooks";
import LoginPage from "@/pages/login";
import LayoutPage from "@/pages/layout";
import HomePage from "@/pages/home";
import NotAuthPage from "@/pages/exception/403.tsx";
import NotFoundPage from "@/pages/exception/404.tsx";
import ServerErrorPage from "@/pages/exception/500.tsx";

export const SystemRouter = (): React.ReactElement | null => {
    const user = useSelector((state) => state.user)
    return useRoutes(!user.status ? [
        {
            path: '*',
            element: <Navigate to={'/login'} />
        },
        {
            path: 'login',
            element: <LoginPage />
        },
    ]: [{
        path: '*',
        element: <Navigate to={'/index'} />
    },
        {
            path: 'login',
            element: <LoginPage />
        },
        {
            path: '',
            element: <LayoutPage />,
            children: [
                {
                    path: 'index',
                    element: <HomePage />
                },
                {
                    path: '403',
                    element: <NotAuthPage />
                },
                {
                    path: '404',
                    element: <NotFoundPage />
                },
                {
                    path: '500',
                    element: <ServerErrorPage />
                }
            ]
        }
    ]);
}