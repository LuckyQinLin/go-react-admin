import React, {useEffect, useState} from "react";
import {matchRoutes, Outlet, RouteObject, useLocation, useNavigate} from "react-router-dom";
import {Layout, theme} from "antd";
import {LayoutHeader, LayoutSider} from "@/pages/layout/components";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import {PermInfo} from "@/redux/user/reducer";
import "./index.less";
import {useSelector} from "@/redux/hooks.ts";
import {useRequest} from "ahooks";
import {userLoginInfo} from "@/api/user.ts";
import {changeLoginStatusActionCreator} from "@/redux/user/action.ts";
import {useDispatch} from "react-redux";
import {IRouteObject, Routes} from "@/router";

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

    const {token: { colorBgContainer }} = theme.useToken();
    const dispatch = useDispatch();
    const [collapsed, setCollapsed] = useState(false);
    const [breadcrumb, setBreadcrumb] = useState<BreadcrumbProp[]>([]);
    const {token, permissions } = useSelector((state) => state.user);
    const navigate = useNavigate();
    const location = useLocation();


    const loadInfo = useRequest(userLoginInfo, {
        manual: true,
        onSuccess: (data) => {
            dispatch(changeLoginStatusActionCreator({...data}));
        }
    });

    useEffect(() => {
        console.log("加载用户信息......");
        loadInfo.run()
    }, []);

    useEffect(() => {
        const routes = matchRoutes(Routes as RouteObject[], location);
        const map = routes?.filter(item => item.pathname === location.pathname).map(item => item.route as IRouteObject);
        if (['/login', '/index', '/500', '/404', '/405'].includes(location.pathname)) {
            return;
        }
        if (!token) {
            navigate('/login');
            return
        }
        if (permissions) {
            if (permissions.length === 1 && permissions[0] === "*:*:*") {
                return;
            }
            if (map && map[0] && map[0].meta) {
                if (!permissions.some(item => item === map[0].meta?.perm)) {
                    navigate('/404');
                    return;
                }
            }
        }
    }, [navigate, location]);


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
