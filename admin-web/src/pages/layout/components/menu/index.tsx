import React, {useState} from "react";
import {Layout, Menu, MenuProps} from "antd";
import {useSelector} from "@/redux/hooks";
import {BreadcrumbProp} from "@/pages/layout/components/header";
type MenuItem = Required<MenuProps>['items'][number];
import './index.less';
import {asyncRoutes} from "@/router";
import {routerBuild, routerBuildMenu} from "@/router/routerFilter.tsx";
import HomeRouter from "@/router/modules/home.tsx";
import PersonRouter from "@/router/modules/person.tsx";
import {UserState} from "@/redux/user/reducer.ts";

interface LayoutHeaderProp {
	collapsed: boolean;
	breadcrumb: (data: BreadcrumbProp[]) => void;
}

const LayoutSider: React.FC<LayoutHeaderProp> = ({collapsed}) => {

	const [defaultMenu, _] = useState<string>('');

	const userStore = useSelector((state) => state.user);


	const buildMenu = (userStore: UserState): MenuItem[] => {
		let router = userStore.permissions?.length === 1 && userStore.permissions[0] === '*:*:*' ?
			asyncRoutes : (userStore.userRouter ?
				routerBuild(userStore.userRouter!) :
				asyncRoutes
			);
		const menus = routerBuildMenu([...HomeRouter, ...router, ...PersonRouter]);
		console.log("router-1", router, menus);
		return menus;
	}


	return <Layout.Sider width={230} trigger={null} collapsible collapsed={collapsed}>
		<div className="ant-pro-sider-logo">
			<a href="">
				<img src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg" alt="logo"/>
				{!collapsed && <h1>通用管理系统</h1>}
			</a>
		</div>
		<Menu
			theme="dark"
			mode="inline"
			defaultSelectedKeys={[defaultMenu]}
			items={buildMenu(userStore)}
		/>
	</Layout.Sider>
}

export default LayoutSider;