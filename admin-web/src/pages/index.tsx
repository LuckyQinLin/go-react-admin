import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {PersistGate} from 'redux-persist/lib/integration/react';
import {Provider} from "react-redux";
import {persist, store} from "@/redux";
import {ConfigProvider} from "antd";
import zhCN from "antd/lib/locale/zh_CN";
import NewRouter from "@/new-router";
import routers = NewRouter.routers;

const Application = () => {

	// return <Provider store={store}>
	// 	<PersistGate loading={null} persistor={persist}>
	// 			<ConfigProvider locale={zhCN}>
	// 				<BrowserRouter>
	// 				{/*<BrowserRouter basename="/admin">*/}
	// 					<Router />
	// 				</BrowserRouter>
	// 		</ConfigProvider>
	// 	</PersistGate>
	// </Provider>

	return <Provider store={store}>
		<PersistGate loading={null} persistor={persist}>
			<ConfigProvider locale={zhCN}>
				<RouterProvider router={createBrowserRouter(routers)} />
			</ConfigProvider>
		</PersistGate>
	</Provider>
}

export default Application;
