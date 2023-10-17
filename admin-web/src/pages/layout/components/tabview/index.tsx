import styled from "@emotion/styled";
import {Tabs, TabsProps} from "antd";


const LayoutTabview = () => {

    const items: TabsProps['items'] = [
        { key: '1', label: '首页', closable: false },
        { key: '2', label: '用户管理' },
        { key: '3', label: '角色管理' },
        { key: '4', label: '角色管理' },
        { key: '5', label: '角色管理' },
        { key: '6', label: '角色管理' },
        { key: '7', label: '角色管理' },
        { key: '8', label: '角色管理' },
        { key: '9', label: '角色管理' },
        { key: '10', label: '角色管理' },
        { key: '11', label: '角色管理' },
        { key: '12', label: '角色管理' },
        { key: '13', label: '角色管理' },
        { key: '14', label: '角色管理' },
        { key: '15', label: '角色管理' },
        { key: '16', label: '角色管理' },
        { key: '17', label: '角色管理' },
        { key: '18', label: '角色管理' },
        { key: '19', label: '角色管理' },
        { key: '20', label: '角色管理' },
    ];

    const onChange = (key: string) => {
        console.log(key);
    };

    return <Container>
        <Tabs
            type="editable-card"
            hideAdd={true}
            tabPosition={'top'}
            defaultActiveKey="1"
            items={items}
            onChange={onChange} />
    </Container>
}

const Container = styled.div`
  background-color: #ffffff;
  border-bottom: 1px solid #f1f2f3;
  
  .ant-tabs-nav-wrap .ant-tabs-nav-wrap-ping-right {
    position: relative;
    display: flex;
    flex: auto;
    align-self: stretch;
    overflow: hidden;
    white-space: nowrap;
    transform: translate(0);
    height: 48px;
    .ant-tabs-nav-list {
      position: relative;
      display: flex;
      transition: opacity 0.3s;
      background-color: #f5f5f5;
      height: 48px;
      flex-direction: row;
      align-content: center;
      align-items: center;
      padding-left: 6px;
    }
    .ant-tabs-nav-operations {
      display: flex;
      height: 48px;
      align-self: stretch;
      background-color: #f5f5f5;
      align-items: center;
      button {
        position: relative;
        padding: 8px 10px;
        background: #fff;
        border: 0;
        color: rgba(0, 0, 0, 0.88);
        height: 35px;
      }
    }
  }
  
  .ant-tabs-nav {
    margin: 0;
  }
  .ant-tabs-nav .ant-tabs-tab {
    border: 0;
    border-radius: 0;
  }
  .ant-tabs-nav .ant-tabs-tab+.ant-tabs-tab { margin-left: 0}
  
  .ant-tabs-tab .ant-tabs-tab-with-remove {
    margin-left: 0;
    border-radius: 0 0 0 0;
  }
  .ant-tabs-nav .ant-tabs-tab-active {
    background-color: #f0f8ff;
  }
  .ant-tabs-card.ant-tabs-top >.ant-tabs-nav .ant-tabs-tab {
    margin-left: 0;
    border-radius: 0 0 0 0;
    border-bottom: 1px solid #f1f2f3;
  }
  
`

export default LayoutTabview;

