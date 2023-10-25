// import React from "react";
// import {Outlet, useLocation, useNavigate} from "react-router-dom";
// import {useEffect} from "react";
// import {useSelector} from "@/redux/hooks";
//
// const RouteGuard = () => {
//
//     const user = useSelector((state) => state.user)
//     const navigate = useNavigate();
//     const location = useLocation();
//
//     useEffect(() => {
//         // 在这里可以进行你的路由守卫逻辑判断
//         const isAuthenticated = user.status; // 根据你的需求判断用户是否已认证
//
//         // 如果用户未认证，则重定向到登录页或其他页面
//         if (!isAuthenticated) {
//             navigate('/login'); // 重定向到登录页
//         }
//     }, [navigate, location]);
//
//     return <Outlet />
// }
//
//
// export default RouteGuard;
