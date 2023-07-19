import {combineReducers, createStore} from 'redux';
import { persistStore, persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
import userReducer from "./user/reducer";
import systemReducer from "@/redux/system/reducer";
import softwareReducer from "@/redux/software/reducer";

const allReducer = combineReducers({
	user: userReducer,
	system: systemReducer,
	software: softwareReducer,
})

const persistConfig = {
	key: 'root',
	storage,
}

export const store = createStore(persistReducer(persistConfig, allReducer))

export const persist = persistStore(store);

export type RootState = ReturnType<typeof store.getState>
