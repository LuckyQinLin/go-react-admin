import React from 'react'
import * as icons from '@ant-design/icons'

const Icon = (props: { icon: string }) => {
    const { icon } = props;
    const antIcon: { [key: string]: any } = icons;
    return React.createElement(antIcon[icon]);
};

export const IconObj = (icon: string) => {
    const antIcon: { [key: string]: any } = icons;
    return antIcon[icon];
}

export default Icon

export const IconKeys = Object.keys(icons).filter(item => item.includes('Outlined'))

