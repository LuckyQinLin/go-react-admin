import React, {useEffect, useState} from "react";
import {Navigate, Outlet, useLocation, useRouteLoaderData} from "react-router-dom";
import {Layout, theme} from "antd";
import {LayoutHeader, LayoutSider, LayoutTabview} from "@/pages/layout/components";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import "./index.less";
import {User} from "@/types";
import NewRouter from "src/router";
import searchRoute = NewRouter.searchRoute;
import routers = NewRouter.routers;
import Router from "@/router/modules.tsx";
import staticPath = Router.staticPath;
import NotFoundPath = Router.NotFoundPath;
import NotAuthPath = Router.NotAuthPath;
import useStore from "@/store/store.ts";


const LayoutPage: React.FC = () => {

    const { pathname } = useLocation();
    const {token: { colorBgContainer }} = theme.useToken();
    const [collapsed, setCollapsed] = useState(false);
    const [breadcrumb, setBreadcrumb] = useState<BreadcrumbProp[]>([]);
    const useInfoFetch = useStore((state) => state.useInfoFetch)

    useEffect(() => { useInfoFetch()}, []);

    const dataLoader = useRouteLoaderData(Router.LayoutId) as User.UserPermissionProp;
    if (!searchRoute(pathname, routers)) {
        return <Navigate to={NotFoundPath} />
    }
    if (!staticPath.includes(pathname) && !dataLoader.paths.includes(pathname)) {
        return <Navigate to={NotAuthPath} />
    }

    const contentCss: React.CSSProperties = {
        margin: '0px 10px 10px 10px',
        padding: 16,
        minHeight: 'calc(100vh - 55px - 48px - 11px)',
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
            <LayoutTabview />
            <Layout.Content style={contentCss}>
                <Outlet />
            </Layout.Content>
        </Layout>
    </Layout>
}

export default LayoutPage;
