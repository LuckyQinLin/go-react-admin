import {Navigate, useRoutes} from "react-router-dom";
import {useSelector} from "@/redux/hooks";
import LoginPage from "@/pages/login";
import LayoutPage from "@/pages/layout";
import HomePage from "@/pages/home";

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
            ]
        }]);
}