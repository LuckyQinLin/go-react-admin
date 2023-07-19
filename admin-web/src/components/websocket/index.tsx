import React, {ReactElement, useEffect, useRef, useState} from "react";
import {useDispatch} from "react-redux";
import {useSelector} from "@/redux/hooks";
import {message, notification} from "antd";
import {downloadEndActionCreator} from "@/redux/software/action";
import {NotificationPlacement} from "antd/es/notification/interface";
import Context from "react-redux/es/components/Context";

// https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fkaygod%2Freact-websocket-demo



const stateArr: WsState[] = [
    {key: 0, value: '正在连接中'},
    {key: 1, value: '已经连接并且可以通讯'},
    {key: 2, value: '连接正在关闭'},
    {key: 3, value: '连接已关闭或者没有连接成功'},
];

interface WsState {
    key: number;
    value: string;
}

interface MessageData {
    modules: string;
    message: string;
    types: string;
}


export const Websocket = (props: {children: ReactElement}) => {

    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const software = useSelector((state) => state.software);
    const ws = useRef<WebSocket | undefined>(undefined);
    //  socket 状态
    const [readyState, setReadyState] = useState<WsState>({ key: 0, value: '正在连接中' })
    const user = useSelector((state) => state.user);

    const distributeMessage = (msg: MessageEvent<any>) => {
        console.log("ws 服务端推送消息 => ", msg);
        const target: MessageData = JSON.parse(msg.data as string);
        switch (target.modules) {
            case "software":
                dispatch(downloadEndActionCreator({id: software.id, is_result: true}));
                openNotification(target.message);
                // message.success(target.message);
                break;
            default:
                break;
        }
    }

    const openNotification = (message: string) => {
        api.success({
            message: '提示',
            description: message,
            placement: 'topRight'
        })
    };


    useEffect(() => {
        console.log("用户状态信息 => ", user);
        // debugger;
        const {token, userId, status} = user;
        if (!status) {
            console.log("退出登录，断开ws链接")
            ws.current?.close()
            ws.current = undefined;
        } else if (status && ws.current === undefined) {
            // 初始化ws
            console.log("登录成功了，开始链接ws")
            try {
                ws.current = new WebSocket(`ws://localhost:8070/api/ws?token=${token}`)
                ws.current.onopen = () => {
                    // 初始化连接的时候发送认证信息
                    ws.current?.send(JSON.stringify({msgType: 1, data: "ok"}))
                    // 设置状态
                    setReadyState(stateArr[ws.current?.readyState ?? 0]);
                }
                ws.current.onclose = () => {
                    setReadyState(stateArr[ws.current?.readyState ?? 0])
                }
                ws.current.onerror = () => {
                    setReadyState(stateArr[ws.current?.readyState ?? 0])
                }
                ws.current.onmessage = (e) => {
                    distributeMessage(e)
                    // console.log("e => ", e)
                    // message.success(e.data)
                }
            } catch (error) {
                console.log(error)
            }
        }
        return () => {
            console.log("销毁组件，断开ws链接")
            ws.current?.close()
        }

    }, [user.status])

    useEffect(() => {
        let timer: NodeJS.Timer | undefined = undefined;
        if (readyState.key === 1) {
            timer = setInterval(() => {
                ws.current?.send(JSON.stringify({msgType: 2, data: "ping"}))
            }, 5000);
        }
        if ((readyState.key === 2 || readyState.key === 3) && timer) {
            clearInterval(timer);
        }
        return () => {
            if (timer) {
                clearInterval(timer);
            }
        }
    }, [readyState])



    return <>
        {contextHolder}
        {props.children}
    </>
}
