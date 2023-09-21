import React, {useMemo, useState} from "react";
import {Layout, Menu, MenuProps} from "antd";
import {useSelector} from "@/redux/hooks";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import {asyncRoutes} from "@/router";
import {routerBuild, routerBuildMenu} from "@/router/routerFilter.tsx";
import HomeRouter from "@/router/modules/home.tsx";
import PersonRouter from "@/router/modules/person.tsx";
import './index.less';

interface LayoutHeaderProp {
	collapsed: boolean;
	breadcrumb: (data: BreadcrumbProp[]) => void;
}

const LayoutSider: React.FC<LayoutHeaderProp> = ({collapsed}) => {

	const [defaultMenu, setDefaultMenu] = useState<string>('home');
	const [openKeys, setOpenKeys] = useState<string[]>([]);

	const userStore = useSelector((state) => state.user);


	const menuItems = useMemo(() => {
		let router = userStore.permissions?.length === 1 && userStore.permissions[0] === '*:*:*' ?
			asyncRoutes : (userStore.userRouter ?
					routerBuild(userStore.userRouter!) :
					asyncRoutes
			);
		return routerBuildMenu([...HomeRouter, ...router, ...PersonRouter]);
	}, [userStore.userRouter])

	const clickMenu = (key: string) => {
		setDefaultMenu(key)
	}

	const onOpenChange: MenuProps['onOpenChange'] = (keys) => {
		console.log("menu", menuItems);
		const latestOpenKey = keys.find((key) => openKeys.indexOf(key) === -1);
		const rootSubmenuKeys = asyncRoutes.map(item => item.meta?.key).filter(item => item != undefined);
		if (latestOpenKey && rootSubmenuKeys.indexOf(latestOpenKey!) === -1) {
			setOpenKeys(keys);
		} else {
			setOpenKeys(latestOpenKey ? [latestOpenKey] : []);
		}
	};


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
			openKeys={openKeys}
			onOpenChange={onOpenChange}
			defaultSelectedKeys={[defaultMenu]}
			items={menuItems}
			onClick={({ key}) => clickMenu(key)}
		/>
	</Layout.Sider>
}

export default LayoutSider;