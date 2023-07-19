import React, {useEffect, useState} from "react";
import {Layout, Menu, MenuProps} from "antd";
import {useNavigate} from "react-router-dom";
import {useSelector} from "@/redux/hooks";
import {useDispatch} from "react-redux";
import {changeMenuActionCreator} from "@/redux/system/action";
import {BreadcrumbProp} from "@/pages/layout/components/header";
import {MenuItemType, SubMenuType} from "antd/es/menu/hooks/useItems";
import {PermInfo} from "@/redux/user/reducer";
type MenuItem = Required<MenuProps>['items'][number];
import {Icon} from "@/components";
import './index.less';
import {paths, permsKeys} from "@/pages/layout";

interface LayoutHeaderProp {
	collapsed: boolean;
	breadcrumb: (data: BreadcrumbProp[]) => void;
}

const buildMenus = (perms: PermInfo[]): MenuItem[] => {
	let result = perms.map(item => {
		return {
			key: item.code,
			label: item.name,
			path: item.path,
			icon: <Icon icon={item.icon!} />,
			children: item.children && item.children.length > 0 ?
				item.children.map<MenuItem>(inner => ({key: inner.code, path: inner.path!, label: inner.name}))
				: undefined
		} as MenuItem
	})
	console.log("menus", result);
	return result;
}

export const menusRouter = (perms: PermInfo[]): Map<string, string> => {
	let map: Map<string, string> = new Map()
	perms.forEach(item => {
		if (item.path) {
			map.set(item.code, item.path)
		}
		if (item.children) {
			item.children.forEach(inner => {
				if (inner.path) {
					map.set(inner.code, inner.path)
				}
			})
		}
	})
	return map;
}

const LayoutSider: React.FC<LayoutHeaderProp> = ({collapsed, breadcrumb}) => {

	const [defaultMenu, setDefaultMenu] = useState<string>('');
	const [routerMap, setRouterMap] = useState<Map<string, string>>();
	const [menus, setMenus] = useState<MenuItem[]>([]);

	const system = useSelector((state) => state.system);
	const userStore = useSelector((state) => state.user);
	const dispatch = useDispatch();
	const navigate = useNavigate();


	const changePage: MenuProps['onClick'] = (e) => {
		console.log("changePage - 0", e)
		const path = routerMap?.get(e.key);
		const keys = [...e.keyPath].reverse();
		if (path) {
			dispatch(changeMenuActionCreator({menuKey: e.key}))
			setDefaultMenu(e.key);
			navigate(path)
		}
		if (keys.length === 1) {
			const temp = menus.find(item => item?.key === keys[0]);
			console.log("changePage - 1", menus, temp);
			if (temp) {
				const temp1 = temp as MenuItemType;
				breadcrumb([{
					label: temp1.label as string,
					path: temp1.key as string,
					icon: temp1.icon,
				}])
			}
		} else if (keys.length === 2) {
			menus.forEach(item => {
				if (item && item.key === keys[0]) {
					const temp = item as SubMenuType;
					if (temp.children) {
						const children = temp.children.find(inner => inner && inner.key === keys[1])
						if (children) {
							const temp1 = children as MenuItemType;
							breadcrumb([{
								label: temp.label as string,
								path: temp.key,
								icon: temp.icon,
							}, {
								label: temp1.label as string,
								path: temp1.key as string
							}])
						}

					}
				}
			})
		}
	};

	useEffect(() => {
		if (userStore.perms) {
			let menusList = buildMenus(userStore.perms);
			setRouterMap(menusRouter(userStore.perms));
			setMenus(menusList);
			setDefaultMenu(permsKeys(userStore.perms)[0]);
			const temp = menusList.find(item => item?.key === defaultMenu)! as MenuItemType;
			if (temp) {
				breadcrumb([{
					label: temp.label as string,
					path: temp.key as string,
					icon: temp.icon,
				}])
			}
		}
		return () => {
			setMenus([])
			setRouterMap(new Map<string, string>());
		}
	}, [])


	return <Layout.Sider width={230} trigger={null} collapsible collapsed={collapsed}>
		<div className="ant-pro-sider-logo">
			<a href="">
				<img src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg" alt="logo"/>
				{!collapsed && <h1>服务器助手</h1>}
			</a>
		</div>
		<Menu
			theme="dark"
			mode="inline"
			defaultSelectedKeys={[defaultMenu]}
			items={menus}
			onClick={changePage}
		/>
	</Layout.Sider>
}

export default LayoutSider;