import { lazy, Suspense } from "react"

// 自定义懒加载函数
const lazyLoad = (factory: () => Promise<any>) => {
    const Module = lazy(factory)
    return (
        <Suspense fallback={<div>加载中......</div>}>
            <Module />
        </Suspense>
    )
}

export default lazyLoad;