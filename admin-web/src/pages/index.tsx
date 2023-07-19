import {BrowserRouter} from "react-router-dom";
import {PersistGate} from 'redux-persist/lib/integration/react';
import {Provider} from "react-redux";
import {persist, store} from "@/redux";
import {SystemRouter} from "@/router";
import {ConfigProvider} from "antd";
import zhCN from "antd/lib/locale/zh_CN";

const Application = () => {
	return <Provider store={store}>
		<PersistGate loading={null} persistor={persist}>
				<ConfigProvider locale={zhCN}>
				{/*<Websocket>*/}
					<BrowserRouter>
					{/*<BrowserRouter basename="/admin">*/}
						<SystemRouter />
					</BrowserRouter>
				{/*</Websocket>*/}
			</ConfigProvider>
		</PersistGate>
	</Provider>
}

export default Application;
