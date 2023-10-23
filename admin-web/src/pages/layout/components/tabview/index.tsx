import styled from "@emotion/styled";
import React, {useEffect, useRef, useState} from "react";
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
import { MenuInfo } from "rc-menu/lib/interface";
import {Menus} from "@/types";
import useStore from "@/store/store.ts";
import {useNavigate} from "react-router-dom";

interface TabTotalProp {
    tabsWith: number;
    parentWith: number;
}

const LayoutTabview: React.FC = () => {

    let navigate = useNavigate();
    const items = useStore(state => state.tabViews)
    const selectedKey = useStore(state => state.tabViewKey)
    const removeTabView = useStore(state => state.removeTabView)
    const setTabViewKey = useStore(state => state.setTabViewKey)
    const closeTabViewAll = useStore(state => state.closeTabViewAll)

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
    const [tabSize, setTabSize] = useState<TabTotalProp>({tabsWith: 0, parentWith: 0});

    /**
     * 关闭tab页面
     * @param event
     * @param key
     */
    const closeTabPage = (event: React.MouseEvent, key: string | number) => {
        event.stopPropagation();
        removeTabView(key as string)
    }

    const dropDownHandler = (e: MenuInfo, data: Menus.TabViewProp) => {
        switch (e.key) {
            case '1':
                // 刷新当前
                break;
            case '2':
                // 关闭当前
                removeTabView(data.key as string);
                break;
            case '3':
                // 关闭其他
                removeTabView(data.key as string, true);
                break;
            case '4':
                // 关闭全部
                closeTabViewAll()
                break;
            default:
                break
        }
    }

    const clickTabView = (key: string) => {
        setTabViewKey(key)
        navigate(key)
    }

    useEffect(() => {
        navigate(selectedKey)
    }, [selectedKey]);

    useEffect(() => {
        getTabViewWith()
    }, [items]);


    const getTabViewWith = () => {
        if (scrollableRef.current) {
            const width = Array.from(scrollableRef.current?.children)
                .reduce((acc, childNode) => {return acc + childNode.clientWidth;}, 0);
            setTimeout(() => setTabSize({
                tabsWith: width,
                parentWith: scrollableRef.current?.clientWidth ?? tabSize.parentWith
            }), 100)
        }
    }


    useEffect(() => {
        getTabViewWith()
        window.addEventListener('resize', getTabViewWith);
        return () => {
            window.removeEventListener('resize', getTabViewWith);
        }
    }, []);

    return <Container>
        {tabSize.parentWith < tabSize.tabsWith && <div className="tab-view-left">
            <LeftOutlined onClick={removeStart} />
        </div>}
        <div className="tab-view-items" ref={scrollableRef}>
            {items.length > 0 &&
                items.map(item =>
                    <Dropdown key={item.key} menu={{ items: dropItems, onClick: (e) => dropDownHandler(e, item) }} trigger={['contextMenu']}>
                        <div
                            id={`tab-view-id-${item.key}`}
                            onClick={() => clickTabView(item.key as string)}
                            className={`tab-view-item-btn ${item.key === selectedKey ? 'tab-view-item-active' : null}`}
                        >
                            <span>{item.title}</span>
                            {item.closeIcon ? null : <CloseOutlined
                                className="span-icon"
                                onClick={(event) => closeTabPage(event, item.key)}
                            />}
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

