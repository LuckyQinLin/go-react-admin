import React from "react";
import {DataNode} from "rc-tree/lib/interface";

/**
 * 对象属性赋值
 * @param data
 * @param val
 * @param key
 */
export function setProps<T>(data: T, val: T[keyof T], key: keyof T): T {
    (<typeof val>data[key]) = val;
    return data;
}

/**
 * addTreeData 递归增加节点内容
 * @param list
 * @param key
 * @param item
 */
export function addTreeData<T extends DataNode>(list: T[], key: React.Key, item: T[]) : T[] {
    return list.map(node => {
        if (node.key === key) {
            if (node.children === undefined) {
                node.children = [];
            }
            node.children.push(...item);
            return node;
        }
        if (node.children) {
            addTreeData(node.children, key, item)
        }
        return node;
    });
}

/**
 * 加载节点
 * @param list
 * @param key
 * @param item
 */
export function loadChildrenTreeData<T extends DataNode>(list: T[], key: React.Key, item: T[]): T[] {
    return list.map(node => {
        // debugger;
        if (node.key === key) {
            if (node.children === undefined) {
                node.children = [];
            }
            for (const inner of item) {
                if (node.children.findIndex(k => k.key === inner.key) === -1) {
                    node.children.push(inner)
                }
            }
            return node;
        }
        if (node.children) {
            loadChildrenTreeData(node.children, key, item)
        }
        return node;
    });
}

/**
 * 更新节点
 * @param list
 * @param key
 * @param item
 * @param field 指定更新的属性
 */
export function updateTreeData<T extends DataNode>(list: T[], key: React.Key, item: T, ...field: ("title" )[]): T[] {
    return list.map((node, index) => {
        if (node.key === key) {
            if (node.children !== undefined) {
                let number = node.children.findIndex(inner => inner.key === item.key);
                if (number !== undefined && number !== -1) {
                    if (field.length > 0) {
                        node.children[number].title = item.title
                    } else {
                        node.children.splice(number, 1, item);
                    }
                }
            } else {
                node.children = []
            }
            return node;
        }
        if (node.children) {
            updateTreeData(node.children, key, item)
        }
        return node;
    });
}


/**
 * 删除节点
 * @param list
 * @param parentKey
 * @param key   支持删除多个
 */
export function deleteTreeData<T extends DataNode>(list: T[], parentKey: React.Key, key: React.Key | React.Key[]): T[] {
    return list.map(node => {
        if (node.key === parentKey) {
            node.children = Array.isArray(key) ?
                node.children?.filter(item => !key.includes(item.key)) :
                node.children?.filter(item => item.key !== key)
            return node;
        }
        if (node.children) {
            deleteTreeData(node.children, parentKey, key)
        }
        return node;
    })
}

