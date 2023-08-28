import React, {useEffect, useState} from "react";
import {Layout, Menu, MenuProps} from "antd";
import {Link} from "react-router-dom";
import {useSelector} from "@/redux/hooks";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import IconFont from "@/components/IconFont";
type MenuItem = Required<MenuProps>['items'][number];
import './index.less';

interface LayoutHeaderProp {
	collapsed: boolean;
	breadcrumb: (data: BreadcrumbProp[]) => void;
}

const LayoutSider: React.FC<LayoutHeaderProp> = ({collapsed, breadcrumb}) => {

	const [defaultMenu, setDefaultMenu] = useState<string>('');
	const [menus, setMenus] = useState<MenuItem[]>([
		{
			label: <Link to={'/home/index'}>'系统首页'</Link>,
			key: 'home',
			icon: <IconFont type="lucky-shouye1" />
		},
		{
			label: <Link to={'/person/index'}>个人中心</Link>,
			key: '/person/index',
			icon: <IconFont type="lucky-jiankong" />
		}
	]);

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