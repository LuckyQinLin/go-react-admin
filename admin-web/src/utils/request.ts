import axios, {AxiosInstance, AxiosRequestConfig, AxiosResponse, Canceler, InternalAxiosRequestConfig} from "axios";
import {message as Message, Modal} from 'antd';
import {store} from "@/redux";
import {cleanUserStoreActionCreator} from "@/redux/user/action.ts";

export enum ResultEnum {
	SUCCESS = 100200,
	DATA_NOT_EXIST = 100100,
	TIMEOUT = 10042,
	CODE_EXPIRE = 100499,
	PARAM_ERROR = 100500,
	BUSINESS_ERROR = 100510,
	TYPE = 'success',
}

export interface Result<T> {
	code: number;
	message: string;
	result: T;
}

export interface PageData<T> {
	records: T[];
	page: number;
	size: number;
	total: number;
}

/**
 * @description: 请求方法
 */
export enum RequestEnum {
	GET = 'GET',
	POST = 'POST',
	PATCH = 'PATCH',
	PUT = 'PUT',
	DELETE = 'DELETE',
}


/**
 * @description: 判断值是否未某个类型
 */
function is(val: unknown, type: string) {
	return toString.call(val) === `[object ${type}]`;
}

/**
 * @description:  是否为函数
 */
function isFunction<T = any>(val: unknown): val is T {
	return is(val, 'Function') || is(val, 'AsyncFunction');
}

function isString(val: unknown): val is string {
	return is(val, 'String');
}

function joinTimestamp(join: boolean, restful = false): string | object {
	if (!join) {
		return restful ? '' : {};
	}
	const now = new Date().getTime();
	if (restful) {
		return `?_t=${now}`;
	}
	return { _t: now };
}

/**
 * 判断是否 url
 * */
function isUrl(url: string) {
	return /^(http|https):\/\//g.test(url);
}

const isObject = (val: any): val is Record<any, any> => {
	return val !== null && is(val, 'Object');
};

/**
 * 将对象添加当作参数拼接到URL上面
 * @param baseUrl 需要拼接的url
 * @param obj 参数对象
 * @returns {string} 拼接后的对象
 * 例子:
 *  let obj = {a: '3', b: '4'}
 *  setObjToUrlParams('www.baidu.com', obj)
 *  ==>www.baidu.com?a=3&b=4
 */
export function setObjToUrlParams(baseUrl: string, obj: object): string {
	let parameters = '';
	let url = '';
	for (const key in obj) {
		// eslint-disable-next-line @typescript-eslint/ban-ts-comment
		// @ts-ignore
		parameters += key + '=' + encodeURIComponent(obj[key]) + '&';
	}
	parameters = parameters.replace(/&$/, '');
	if (/\?$/.test(baseUrl)) {
		url = baseUrl + parameters;
	} else {
		url = baseUrl.replace(/\/?$/, '?') + parameters;
	}
	return url;
}

const DATE_TIME_FORMAT = 'YYYY-MM-DD HH:mm';

type Recordable<T = any> = Record<string, T>

function formatRequestDate(params: Recordable) {
	if (Object.prototype.toString.call(params) !== '[object Object]') {
		return;
	}

	for (const key in params) {
		if (params[key] && params[key]._isAMomentObject) {
			params[key] = params[key].format(DATE_TIME_FORMAT);
		}
		if (isString(key)) {
			const value = params[key];
			if (value) {
				try {
					params[key] = isString(value) ? value.trim() : value;
				} catch (error) {
					throw new Error(error as any);
				}
			}
		}
		if (isObject(params[key])) {
			formatRequestDate(params[key]);
		}
	}
}

function checkStatus(status: number, msg: string): void {
	switch (status) {
		case 400:
			Message.error(msg);
			break;
		// 401: 未登录
		// 未登录则跳转登录页面，并携带当前页面的路径
		// 在登录成功后返回当前页面，这一步需要在登录页操作。
		case 401:
			Message.error('认证失败，' + msg);
			store.dispatch(cleanUserStoreActionCreator())
			window.location.href = '/login';
			break;
		case 403:
			Message.error('用户得到授权，但是访问是被禁止的。!');
			break;
		// 404请求不存在
		case 404:
			Message.error('网络请求错误，未找到该资源!');
			break;
		case 405:
			Message.error('网络请求错误，请求方法未允许!');
			break;
		case 408:
			Message.error('网络请求超时');
			break;
		case 500:
			Message.error('服务器错误,请联系管理员!');
			break;
		case 501:
			Message.error('网络未实现');
			break;
		case 502:
			Message.error('网络错误');
			break;
		case 503:
			Message.error('服务不可用，服务器暂时过载或维护!');
			break;
		case 504:
			Message.error('网络超时');
			break;
		case 505:
			Message.error('http版本不支持该请求!');
			break;
		default:
			Message.error(msg);
	}
}


export abstract class AxiosTransform {
	/**
	 * @description: 请求之前处理配置
	 * @description: Process configuration before request
	 */
	beforeRequestHook?: (config: AxiosRequestConfig, options: RequestOptions) => AxiosRequestConfig;

	/**
	 * @description: 请求成功处理
	 */
	transformRequestData?: (res: AxiosResponse<Result<any>>, options: RequestOptions) => any;

	/**
	 * @description: 请求失败处理
	 */
	requestCatch?: (e: Error) => Promise<any>;

	/**
	 * @description: 请求之前的拦截器
	 */
	requestInterceptors?: (config: AxiosRequestConfig, options: CreateAxiosOptions) => AxiosRequestConfig;

	/**
	 * @description: 请求之后的拦截器
	 */
	responseInterceptors?: (res: AxiosResponse<any>) => AxiosResponse<any>;

	/**
	 * @description: 请求之前的拦截器错误处理
	 */
	requestInterceptorsCatch?: (error: Error) => void;

	/**
	 * @description: 请求之后的拦截器错误处理
	 */
	responseInterceptorsCatch?: (error: Error) => void;
}

export interface RequestOptions {
	// 请求参数拼接到url
	joinParamsToUrl?: boolean;
	// 格式化请求参数时间
	formatDate?: boolean;
	// 是否显示提示信息
	isShowMessage?: boolean;
	// 是否解析成JSON
	isParseToJson?: boolean;
	// 成功的文本信息
	successMessageText?: string;
	// 是否显示成功信息
	isShowSuccessMessage?: boolean;
	// 是否显示失败信息
	isShowErrorMessage?: boolean;
	// 错误的文本信息
	errorMessageText?: string;
	// 是否加入url
	joinPrefix?: boolean;
	// 接口地址， 不填则使用默认apiUrl
	apiUrl?: string;
	// 请求拼接路径
	urlPrefix?: string;
	// 错误消息提示类型
	errorMessageMode?: 'none' | 'modal';
	// 是否添加时间戳
	joinTime?: boolean;
	// 不进行任何处理，直接返回
	isTransformResponse?: boolean;
	// 是否返回原生响应头
	isReturnNativeResponse?: boolean;
	//忽略重复请求
	ignoreCancelToken?: boolean;
	// 是否携带token
	withToken?: boolean;
	isLoading?: boolean;
}

export interface CreateAxiosOptions extends AxiosRequestConfig {
	transform?: AxiosTransform;
	requestOptions?: RequestOptions;
	authenticationScheme?: string;
}

let pendingMap = new Map<string, Canceler>();
const getPendingUrl = (config: AxiosRequestConfig) => [config.method, config.url, JSON.stringify(config.data), JSON.stringify(config.params)].join('&');

export class AxiosCanceler {
	/**
	 * 添加请求
	 * @param {Object} config
	 */
	addPending(config: AxiosRequestConfig) {
		this.removePending(config);
		const url = getPendingUrl(config);
		config.cancelToken =
			config.cancelToken ||
			new axios.CancelToken((cancel) => {
				if (!pendingMap.has(url)) {
					// 如果 pending 中不存在当前请求，则添加进去
					pendingMap.set(url, cancel);
				}
			});
	}

	/**
	 * @description: 清空所有pending
	 */
	removeAllPending() {
		pendingMap.forEach((cancel) => {
			cancel && isFunction(cancel) && cancel();
		});
		pendingMap.clear();
	}

	/**
	 * 移除请求
	 * @param {Object} config
	 */
	removePending(config: AxiosRequestConfig) {
		const url = getPendingUrl(config);

		if (pendingMap.has(url)) {
			// 如果在 pending 中存在当前请求标识，需要取消当前请求，并且移除
			const cancel = pendingMap.get(url);
			cancel && cancel(url);
			pendingMap.delete(url);
		}
	}

	/**
	 * @description: 重置
	 */
	reset(): void {
		pendingMap = new Map<string, Canceler>();
	}
}

/**
 * @description: 数据处理，方便区分多种处理方式
 */
const transform: AxiosTransform = {
	/**
	 * @description: 处理请求数据
	 */
	transformRequestData: (res: AxiosResponse<Result<any>>, options: RequestOptions) => {
		const {
			isShowMessage = true,
			isShowErrorMessage,
			isShowSuccessMessage,
			successMessageText,
			errorMessageText,
			isTransformResponse,
			isReturnNativeResponse,
		} = options;

		// 是否返回原生响应头 比如：需要获取响应头时使用该属性
		if (isReturnNativeResponse) {
			return res;
		}
		// 不进行任何处理，直接返回
		// 用于页面代码可能需要直接获取code，data，message这些信息时开启
		if (!isTransformResponse) {
			return res.data;
		}

		const { data } = res;

		if (!data) {
			throw new Error('请求出错，请稍候重试');
		}
		//  这里 code，result，message为 后台统一的字段，需要修改为项目自己的接口返回格式
		const { code, result, message } = data;
		// 请求成功
		const hasSuccess = data && Reflect.has(data, 'code') && code === ResultEnum.SUCCESS;
		// 是否显示提示信息
		if (isShowMessage) {
			if (hasSuccess && (successMessageText || isShowSuccessMessage)) {
				// 是否显示自定义信息提示
				Modal.success({
					title: '成功',
					content: successMessageText || message || '操作成功！',
				})
			} else if (!hasSuccess && (errorMessageText || isShowErrorMessage)) {
				// 是否显示自定义信息提示
				Message.error(message || errorMessageText || '操作失败！');
			} else if (!hasSuccess && options.errorMessageMode === 'modal') {
				// errorMessageMode=‘custom-modal’的时候会显示modal错误弹窗，而不是消息提示，用于一些比较重要的错误
				Modal.success({
					title: '提示',
					content: message,
				})
			}
		}

		// 接口请求成功，直接返回结果
		if (code === ResultEnum.SUCCESS) {
			return result;
		}
		// 接口请求错误，统一提示错误信息 这里逻辑可以根据项目进行修改
		const errorMsg = message;
		switch (code) {
			case ResultEnum.CODE_EXPIRE:
				Message.error(result);
				break;
			case ResultEnum.BUSINESS_ERROR:
				Message.error(errorMsg);
				break;
			// 请求失败
			case ResultEnum.DATA_NOT_EXIST:
				Message.error(errorMsg);
				break;
			case ResultEnum.PARAM_ERROR:
				Message.error(errorMsg);
				break;
			// 登录超时
			case ResultEnum.TIMEOUT:
				// const LoginName = SystemConfig.BASE_LOGIN_NAME;
				// const LoginPath = SystemConfig.BASE_LOGIN;
				// if (router.currentRoute.value?.name === LoginName) return;
				// // 到登录页
				// errorMsg = '登录超时，请重新登录!';
				// Modal.warning({
				// 	title: '警告',
				// 	content: '登录身份已失效，请重新登录!',
				// 	onOk: () => {
				// 		// TODO 清空storage
				// 		window.location.href = LoginPath;
				// 	}
				// })
				break;
		}
		throw new Error(errorMsg);
		// return false;
	},

	// 请求之前处理config
	beforeRequestHook: (config, options) => {
		const { apiUrl, joinPrefix, joinParamsToUrl, formatDate, joinTime = true, urlPrefix } = options;

		const isUrlStr = isUrl(config.url as string);

		if (!isUrlStr && joinPrefix) {
			config.url = `${urlPrefix}${config.url}`;
		}

		if (!isUrlStr && apiUrl && isString(apiUrl)) {
			config.url = `${apiUrl}${config.url}`;
		}
		const params = config.params || {};
		const data = config.data || false;
		if (config.method?.toUpperCase() === RequestEnum.GET) {
			if (!isString(params)) {
				// 给 get 请求加上时间戳参数，避免从缓存中拿数据。
				config.params = Object.assign(params || {}, joinTimestamp(joinTime, false));
			} else {
				// 兼容restful风格
				config.url = config.url + params + `${joinTimestamp(joinTime, true)}`;
				config.params = undefined;
			}
		} else {
			if (!isString(params)) {
				formatDate && formatRequestDate(params);
				if (Reflect.has(config, 'data') && config.data && Object.keys(config.data).length > 0) {
					config.data = data;
					config.params = params;
				} else {
					config.data = params;
					config.params = undefined;
				}
				if (joinParamsToUrl) {
					config.url = setObjToUrlParams(
						config.url as string,
						Object.assign({}, config.params, config.data)
					);
				}
			} else {
				// 兼容restful风格
				config.url = config.url + params;
				config.params = undefined;
			}
		}
		return config;
	},

	/**
	 * @description: 请求拦截器处理
	 */
	requestInterceptors: (config, options) => {
		// 请求之前处理config
		const token = store.getState().user.token;
		if (token && (config as Recordable)?.requestOptions?.withToken !== false) {
			// jwt token
			(config as Recordable).headers.Authorization = options.authenticationScheme
				? `${options.authenticationScheme} ${token}`
				: token;
		}
		return config;
	},

	/**
	 * @description: 响应错误处理
	 */
	responseInterceptorsCatch: (error: any) => {
		const { response, code, message } = error || {};
		// TODO 此处要根据后端接口返回格式修改
		const msg: string =
			response && response.data && response.data.message ? response.data.message : '';
		const err: string = error.toString();
		try {
			if (code === 'ECONNABORTED' && message.indexOf('timeout') !== -1) {
				Message.error('接口请求超时，请刷新页面重试!');
				return;
			}
			if (err && err.includes('Network Error')) {
				Modal.info({
					title: '网络异常',
					content: '请检查您的网络连接是否正常',
				})
				return Promise.reject(error);
			}
		} catch (error) {
			throw new Error(error as any);
		}
		// 请求是否被取消
		const isCancel = axios.isCancel(error);
		if (!isCancel) {
			checkStatus(error.response && error.response.status, msg);
		} else {
			console.warn(error, '请求被取消！');
		}
		//return Promise.reject(error);
		return Promise.reject(response?.data);
	},
};


export class AxiosPlus {
	private axiosInstance: AxiosInstance;
	private readonly options: CreateAxiosOptions;

	constructor(options: CreateAxiosOptions) {
		this.options = options;
		this.axiosInstance = axios.create(options);
		this.setupInterceptors();
	}

	getAxios(): AxiosInstance {
		return this.axiosInstance;
	}

	/**
	 * @description: 重新配置axios
	 */
	configAxios(config: CreateAxiosOptions) {
		if (!this.axiosInstance) {
			return;
		}
		this.createAxios(config);
	}

	/**
	 * @description: 设置通用header
	 */
	setHeader(headers: any): void {
		if (!this.axiosInstance) {
			return;
		}
		Object.assign(this.axiosInstance.defaults.headers, headers);
	}

	/**
	 * @description:  创建axios实例
	 */
	private createAxios(config: CreateAxiosOptions): void {
		this.axiosInstance = axios.create(config);
	}

	private getTransform() {
		const { transform } = this.options;
		return transform;
	}

	/**
	 * @description:   请求方法
	 */
	request<T = any>(config: AxiosRequestConfig, options?: RequestOptions): Promise<T> {
		let conf: AxiosRequestConfig = JSON.parse(JSON.stringify(config));
		const transform = this.getTransform();
		const { requestOptions } = this.options;
		const opt: RequestOptions = Object.assign({}, requestOptions, options);

		const { beforeRequestHook, requestCatch, transformRequestData } = transform || {};
		if (beforeRequestHook && isFunction(beforeRequestHook)) {
			conf = beforeRequestHook(conf, opt);
		}
		// 挂载属性
		// eslint-disable-next-line @typescript-eslint/ban-ts-comment
		// @ts-ignore
		conf.requestOption = opt;

		// 发送请求
		return new Promise((resolve, reject) => {
			this.axiosInstance
				.request<any, AxiosResponse<Result<any>>>(conf)
				.then((res: AxiosResponse<Result<any>>) => {
					// 请求是否被取消
					const isCancel = axios.isCancel(res);
					if (transformRequestData && isFunction(transformRequestData) && !isCancel) {
						try {
							const ret = transformRequestData(res, opt);
							resolve(ret);
						} catch (err) {
							reject(err || new Error('request error!'));
						}
						return;
					}
					resolve(res as unknown as Promise<T>);
				})
				.catch((e: Error) => {
					if (requestCatch && isFunction(requestCatch)) {
						reject(requestCatch(e));
						return;
					}
					reject(e);
				});
		});
	}

	/**
	 * @description: 拦截器配置
	 */
	private setupInterceptors() {
		const transform = this.getTransform();
		if (!transform) {
			return;
		}
		const {
			requestInterceptors,
			requestInterceptorsCatch,
			responseInterceptors,
			responseInterceptorsCatch,
		} = transform;

		// 取消请求
		const axiosCanceler = new AxiosCanceler();

		// 请求拦截器配置处理
		this.axiosInstance.interceptors.request.use((config) => {
			const { headers: { ignoreCancelToken } } = config;
			const ignoreCancel =
				ignoreCancelToken !== undefined
					? ignoreCancelToken
					: this.options.requestOptions?.ignoreCancelToken;

			!ignoreCancel && axiosCanceler.addPending(config);
			if (requestInterceptors && isFunction(requestInterceptors)) {
				config = requestInterceptors(config, this.options) as InternalAxiosRequestConfig;
			}
			return config;
		}, undefined);

		// 请求拦截器错误捕获
		requestInterceptorsCatch &&
		isFunction(requestInterceptorsCatch) &&
		this.axiosInstance.interceptors.request.use(undefined, requestInterceptorsCatch);

		// 响应结果拦截器处理
		this.axiosInstance.interceptors.response.use((res: AxiosResponse<any>) => {
			res && axiosCanceler.removePending(res.config);
			if (responseInterceptors && isFunction(responseInterceptors)) {
				res = responseInterceptors(res);
			}
			return res;
		}, undefined);

		// 响应结果拦截器错误捕获
		responseInterceptorsCatch &&
		isFunction(responseInterceptorsCatch) &&
		this.axiosInstance.interceptors.response.use(undefined, responseInterceptorsCatch);
	}
}

function deepMerge<T = any>(src: any = {}, target: any = {}): T {
	let key: string;
	for (key in target) {
		src[key] = isObject(src[key]) ? deepMerge(src[key], target[key]) : (src[key] = target[key]);
	}
	return src;
}

export enum ContentTypeEnum {
	// json
	JSON = 'application/json;charset=UTF-8',
	// json
	TEXT = 'text/plain;charset=UTF-8',
	// form-data 一般配合qs
	FORM_URLENCODED = 'application/x-www-form-urlencoded;charset=UTF-8',
	// form-data  上传
	FORM_DATA = 'multipart/form-data;charset=UTF-8',
}

function createAxios(opt?: Partial<CreateAxiosOptions>) {
	return new AxiosPlus(
		deepMerge(
			{
				timeout: 10 * 1000,
				authenticationScheme: '',
				// 接口前缀
				prefixUrl: '/api',
				headers: { 'Content-Type': ContentTypeEnum.JSON },
				// 数据处理方式
				transform,
				// 配置项，下面的选项都可以在独立的接口请求中覆盖
				requestOptions: {
					// 默认将prefix 添加到url
					joinPrefix: true,
					// 是否返回原生响应头 比如：需要获取响应头时使用该属性
					isReturnNativeResponse: false,
					// 需要对返回数据进行处理
					isTransformResponse: true,
					// post请求的时候添加参数到url
					joinParamsToUrl: false,
					// 格式化提交参数时间
					formatDate: true,
					// 消息提示类型
					errorMessageMode: 'none',
					// 接口地址
					apiUrl: '',
					// 接口拼接地址
					urlPrefix: '/api',
					//  是否加入时间戳
					joinTime: true,
					// 忽略重复请求
					ignoreCancelToken: true,
					// 是否携带token
					withToken: true,
				},
				withCredentials: false,
			},
			opt || {}
		)
	);
}


export const https = createAxios();