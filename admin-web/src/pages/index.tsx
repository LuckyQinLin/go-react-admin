import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {ConfigProvider} from "antd";
import zhCN from "antd/lib/locale/zh_CN";
import NewRouter from "src/router";
import routers = NewRouter.routers;

const Application = () => {

	return <ConfigProvider locale={zhCN}>
		<RouterProvider router={createBrowserRouter(routers)} />
	</ConfigProvider>
}

export default Application;
