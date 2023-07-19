import React, {useEffect, useState} from "react";
import {Outlet, useLocation, useNavigate} from "react-router-dom";
import {Layout, theme} from "antd";
import {LayoutHeader, LayoutSider} from "@/pages/layout/components";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import "./index.less";
import {useSelector} from "@/redux/hooks";
import {PermInfo} from "@/redux/user/reducer";

export const paths = (perms: PermInfo[]): Set<string> => {
    let path = new Set<string>()
    perms.forEach(item => {
        if (item.path) {
            path.add(item.path)
        }
        if (item.children) {
            item.children.forEach(inner => {
                if (inner.path) {
                    path.add(inner.path)
                }
            })
        }
    })
    return path
}

export const permsKeys = (perms: PermInfo[]): string[] => {
    let result: string[] = [];
    perms.forEach(item => {
        result.push(item.code)
        if (item.children) {
            item.children.forEach(inner => {
                result.push(inner.code)
                if (inner.children) {
                    result.push(...inner.children.map(end => end.code))
                }
            })
        }
    })
    return result;
}

const LayoutPage: React.FC = () => {

    const routers = ['/login', '/404', '/403', '/500'];
    const {token: { colorBgContainer }} = theme.useToken();
    const [collapsed, setCollapsed] = useState(false);
    const [breadcrumb, setBreadcrumb] = useState<BreadcrumbProp[]>([]);
    const [routerSet, setRouterSet] = useState<Set<string>>(new Set<string>());

    const user = useSelector((state) => state.user)
    const navigate = useNavigate();
    const location = useLocation();

    useEffect(() => {
        // 在这里可以进行你的路由守卫逻辑判断
        // console.log("navigate", navigate);
        console.log("location", location, user);

        if (!routers.includes(location.pathname)) {
            const isAuthenticated = user.status; // 根据你的需求判断用户是否已认证
            // 如果用户未认证，则重定向到登录页或其他页面
            if (!isAuthenticated) {
                navigate('/login'); // 重定向到登录页
                return
            }
            // 防止第一次无法加载路径导致404
            if (routerSet.size <= 0 && paths(user.perms!).has(location.pathname)) {
                return;
            }
            if (!routerSet.has(location.pathname)) {
                navigate('/404');
                return;
            }
        }
    }, [navigate, location]);


    useEffect(() => {
        if (user.status && user.perms) {
            setRouterSet(paths(user.perms));
        }
        return () => setRouterSet(new Set<string>())
    }, [user])


    const contentCss: React.CSSProperties = {
        margin: '16px 16px',
        padding: 16,
        minHeight: 'calc(100vh - 64px - 32px)',
        background: colorBgContainer
    }

    return <Layout className="admin-layout-area">
        <LayoutSider collapsed={collapsed} breadcrumb={setBreadcrumb} />
        <Layout>
            <LayoutHeader
                breadcrumb={breadcrumb}
                bgColor={colorBgContainer}
                collapsed={collapsed}
                changeCollapsed={setCollapsed}
            />
            <Layout.Content style={contentCss}>
                <Outlet />
            </Layout.Content>
        </Layout>
    </Layout>
}

export default LayoutPage;
