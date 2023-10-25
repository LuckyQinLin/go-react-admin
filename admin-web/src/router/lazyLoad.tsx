import { lazy, Suspense } from "react"
import {Spin} from "antd";

// 自定义懒加载函数
const lazyLoad = (factory: () => Promise<any>) => {
    const Module = lazy(factory)
    return (
        <Suspense fallback={<Spin
            size='large'
            style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}/>}>
            <Module />
        </Suspense>
    )
}

export default lazyLoad;