import React, {useEffect, useState} from "react";
import {Layout, Menu, MenuProps} from "antd";
import {useSelector} from "@/redux/hooks";
import {BreadcrumbProp} from "@/pages/layout/components/header";
type MenuItem = Required<MenuProps>['items'][number];
import './index.less';

interface LayoutHeaderProp {
	collapsed: boolean;
	breadcrumb: (data: BreadcrumbProp[]) => void;
}

const LayoutSider: React.FC<LayoutHeaderProp> = ({collapsed}) => {

	const [defaultMenu, _] = useState<string>('');

	const [menus, setMenus] = useState<MenuItem[]>([]);

	const userStore = useSelector((state) => state.user);


	useEffect(() => {
		if (userStore.menus) {
			setMenus(userStore.menus)
		}
		return () => {}
	}, [])

	useEffect(() => {
		console.log('发生变化', userStore.menus)
		if (userStore.menus) {
			setMenus(userStore.menus)
		}
	}, [userStore.menus]);


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
			items={menus}
			// onClick={changePage}
		/>
	</Layout.Sider>
}

export default LayoutSider;