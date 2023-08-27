import React from "react";

export interface IRouteObject {
    children?: Array<IRouteObject>,
    element?: React.ReactNode,
    redirect?: string,
    path?: string
    meta?: {
        isRoot?: boolean,
        title?: string,
        sort?: number,
        icon?: React.ReactNode,
        permission?: Array<string>
    }
}