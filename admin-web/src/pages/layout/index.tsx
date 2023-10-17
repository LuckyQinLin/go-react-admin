import React, {useEffect, useState} from "react";
import {Outlet} from "react-router-dom";
import {Layout, theme} from "antd";
import {LayoutHeader, LayoutSider, LayoutTabview} from "@/pages/layout/components";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import {PermInfo} from "@/redux/user/reducer";
import "./index.less";


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
    const [collapsed, setCollapsed] = useState(false);
    const [breadcrumb, setBreadcrumb] = useState<BreadcrumbProp[]>([]);

    useEffect(() => {
        // getUserInfo()
    }, []);

    const contentCss: React.CSSProperties = {
        margin: '0px 10px 10px 10px',
        padding: 16,
        minHeight: 'calc(100vh - 55px - 40px - 32px)',
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
