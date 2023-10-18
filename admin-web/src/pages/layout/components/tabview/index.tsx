import styled from "@emotion/styled";
import React, {ReactElement, useEffect, useRef, useState} from "react";
import {
    CloseOutlined,
    ColumnWidthOutlined,
    LeftOutlined,
    MinusOutlined,
    RedoOutlined,
    RightOutlined
} from "@ant-design/icons";
import {Dropdown} from "antd";
import type { MenuProps } from 'antd';
import {cleanUserStoreActionCreator} from "@/redux/user/action.ts";
import { MenuInfo } from "rc-menu/lib/interface";


interface TabViewProp {
    key: string | number;
    title: string | React.ReactNode;
    closeIcon?: boolean | React.ReactNode
}

interface TabTotailProp {
    tabsWith: number;
    parentWith: number;
}

const LayoutTabview = () => {

    const [items, setItems] = useState<TabViewProp[]>([
        {key: 'home', title: '首页', closeIcon: true},
        {key: 'user', title: '用户管理'},
        {key: 'account', title: '系统管理'},
        {key: 'person', title: '人员管理'},
        {key: 'role', title: '角色管理'},
        {key: 'resource', title: '资源管理'},
        {key: 'file', title: '文件管理'},
        {key: 'logger', title: '日志管理'},
        {key: 'message', title: '消息管理'},
    ]);

    const dropItems: MenuProps['items'] = [
        { label: '刷新当前', key: '1', icon: <RedoOutlined />},
        { label: '关闭当前', key: '2', icon: <CloseOutlined />},
        { label: '关闭其他', key: '3', icon: <ColumnWidthOutlined />},
        { label: '关闭全部', key: '4', icon: <MinusOutlined />},
    ];

    const removeStart = () => {
        if (scrollableRef.current) {
            scrollableRef.current.scrollTo({ left: 0, behavior: 'smooth' });
        }
    }

    const removeEnd = () => {
        if (scrollableRef.current) {
            const maxScroll = scrollableRef.current.clientWidth;
            scrollableRef.current.scrollTo({ left: maxScroll, behavior: 'smooth' });
        }
    }

    const scrollableRef = useRef<HTMLDivElement | null>(null);
    const [tabSize, setTabSize] = useState<TabTotailProp>({tabsWith: 0, parentWith: 0});

    const [selectedKey, setSelectedKey] = useState<string | number>('home');

    const handleResize = () => {
        console.log(tabSize, scrollableRef.current?.clientWidth)
        setInterval(() => {
            setTabSize({...tabSize, parentWith: scrollableRef.current?.clientWidth ?? tabSize.parentWith})
        }, 100)
    };

    /**
     * 关闭tab页面
     * @param key
     */
    const closeTabPage = (key: string | number) => {

    }

    const dropDownHandler = (e: MenuInfo, data: TabViewProp) => {
        switch (e.key) {
            case '1':
            // 刷新当前
            case '2':
            // 关闭当前
            case '3':
            // 关闭其他
            case '4':
            // 关闭全部
            default:
                break
        }
    }



    useEffect(() => {
        if (scrollableRef.current) {
            const width = Array.from(scrollableRef.current?.children)
                .reduce((acc, childNode) => {return acc + childNode.clientWidth;}, 0);
            setTabSize({tabsWith: width, parentWith: scrollableRef.current?.clientWidth})
        }
        window.addEventListener('resize', handleResize);
        return () => {
            window.removeEventListener('resize', handleResize);
        }
    }, []);

    return <Container>
        {tabSize.parentWith < tabSize.tabsWith && <div className="tab-view-left">
            <LeftOutlined onClick={removeStart} />
        </div>}
        <div className="tab-view-items" ref={scrollableRef}>
            {items.length > 0 &&
                items.map(item =>
                    <Dropdown menu={{ items: dropItems, onClick: (e) => dropDownHandler(e, item) }} trigger={['contextMenu']}>
                        <div
                            id={`tab-view-id-${item.key}`}
                            onClick={() => setSelectedKey(item.key)}
                            className={`tab-view-item-btn ${item.key === selectedKey ? 'tab-view-item-active' : null}`}
                        >
                            <span>{item.title}</span>
                            {item.closeIcon ? null : <CloseOutlined className="span-icon" onClick={() => closeTabPage(item.key)} />}
                        </div>
                    </Dropdown>
                )
            }
        </div>
        {tabSize.parentWith < tabSize.tabsWith && <div className="tab-view-right">
            <RightOutlined onClick={removeEnd}/>
        </div>}
    </Container>
}

const Container = styled.div`
  background-color: #f5f5f5;
  border-bottom: 1px solid #f1f2f3;
  height: 42px;
  display: flex;
  flex-direction: row;
  align-items: center;  
  padding-left: 10px;
  .tab-view-left {
    color: rgb(31, 34, 37);
    background: #fff;
    flex-shrink: 0;
    border-radius: 3px;
    margin-right: 6px;
    text-align: center;
    span {
      width: 32px;
      height: 32px;
      svg {
        width: 32px;
      }
    }
  }
  .tab-view-right {
    color: rgb(31, 34, 37);
    background: #fff;
    flex-shrink: 0;
    border-radius: 3px;
    margin-left: 6px;
    margin-right: 10px;
    text-align: center;
    span {
      width: 32px;
      height: 32px;
      svg {
        width: 32px;
      }
    }
  }
  .tab-view-items {
    flex-grow: 1;
    display: flex;
    flex-direction: row;
    white-space: nowrap;
    overflow: hidden;
    .tab-view-item-btn {
      background: #fff;
      color: rgb(31, 34, 37);
      height: 32px;
      padding: 6px 12px 4px;
      border-radius: 3px;
      margin-right: 6px;
      cursor: pointer;
      display: inline-block;
      position: relative;
      flex: 0 0 auto;
      span {
        float: left;
        vertical-align: middle;
        margin-right: 5px;
      }
      .span-icon {
        height: 22px;
        width: 21px;
        margin-right: -6px;
        position: relative;
        vertical-align: middle;
        text-align: center;
        color: #808695;
      }
    }
    .tab-view-item-active {
      color: #2d8cf0;
    }
  }
`

export default LayoutTabview;

