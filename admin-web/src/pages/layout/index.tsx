import React, {useEffect, useState} from "react";
import {Layout, theme} from "antd";
import {LayoutHeader, LayoutSider, LayoutTabview} from "@/pages/layout/components";
import {Navigate, Outlet, useLocation, useRouteLoaderData} from "react-router-dom";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import {User} from "@/types";
import RouterSpace from "@/router";
import RouterVariate from "@/router/modules.tsx";
import useStore from "@/store/store.ts";
import searchRoute = RouterSpace.searchRoute;
import "./index.less";


const LayoutPage: React.FC = () => {

    const { pathname } = useLocation();
    const {token: { colorBgContainer }} = theme.useToken();
    const [collapsed, setCollapsed] = useState(false);
    const [breadcrumb, setBreadcrumb] = useState<BreadcrumbProp[]>([]);
    const useInfoFetch = useStore((state) => state.useInfoFetch)

    useEffect(() => { useInfoFetch()}, []);

    const dataLoader = useRouteLoaderData(RouterVariate.LayoutId) as User.UserPermissionProp;
    if (!searchRoute(pathname, RouterSpace.routers)) {
        return <Navigate to={RouterVariate.NotFoundPath} />
    }
    if (!RouterVariate.staticPath.includes(pathname) && !dataLoader.paths.includes(pathname)) {
        return <Navigate to={RouterVariate.NotAuthPath} />
    }

    const contentCss: React.CSSProperties = {
        margin: '0px 10px 10px 10px',
        // padding: 16,
        minHeight: 'calc(100vh - 55px - 41px - 11px)',
        // background: colorBgContainer
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
