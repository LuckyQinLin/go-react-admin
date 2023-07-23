import React, {useEffect, useState} from "react";
import {Outlet} from "react-router-dom";
import {Layout, theme} from "antd";
import {LayoutHeader, LayoutSider} from "@/pages/layout/components";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import {PermInfo} from "@/redux/user/reducer";
import "./index.less";
import {useSelector} from "@/redux/hooks.ts";

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
    const [collapsed, setCollapsed] = useState(false);
    const [breadcrumb, setBreadcrumb] = useState<BreadcrumbProp[]>([]);
    const user = useSelector((state) => state.user)

    useEffect(() => {
        console.log("user", user);
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
