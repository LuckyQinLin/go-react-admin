import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {ConfigProvider} from "antd";
import RouterSpace from "@/router";
import zhCN from "antd/lib/locale/zh_CN";
import routers = RouterSpace.routers;

const Application = () => {

	return <ConfigProvider locale={zhCN}>
		<RouterProvider router={createBrowserRouter(routers)} />
	</ConfigProvider>
}

export default Application;
