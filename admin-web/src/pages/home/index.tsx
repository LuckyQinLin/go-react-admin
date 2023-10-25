import {Calendar, CalendarProps, Card, Col, Row, Statistic} from "antd";
import UserImg from '@/pages/home/imgs/user.png';
import RoleImg from '@/pages/home/imgs/role.png';
import ResourceImg from '@/pages/home/imgs/resource.png';
import ServerImg from '@/pages/home/imgs/server.png';
import PersonImg from '@/pages/home/imgs/person.png';
import {Dayjs} from "dayjs";
import React from "react";
import styled from "@emotion/styled";
import VisitCharts from "@/pages/home/VisitCharts.tsx";

const HomePage = () => {
    const onPanelChange = (value: Dayjs, mode: CalendarProps<Dayjs>['mode']) => {
        console.log(value.format('YYYY-MM-DD'), mode);
    };

    const wrapperStyle: React.CSSProperties = {
        borderRadius: 5,
        marginBottom: 5,
    };

    return <>
        {/* 统计 */}
        {/* 快捷入口 */}
        {/* 日历 */}
        {/* 访问统计 左 */}
        {/* 更新日志 右 */}
        <Row gutter={[5, 5]}>
            <Col span={18}>
                <StatisticCard title="统计信息" style={wrapperStyle}>
                    <Row>
                        <Col span={4}>
                            <Statistic title="在线用户" value={112} precision={0} valueStyle={{ color: '#3f8600' }} suffix="人"/>
                        </Col>
                        <Col span={4}>
                            <Statistic title="新增用户" value={11} precision={0} valueStyle={{ color: '#faad14' }} suffix="人"/>
                        </Col>
                        <Col span={4}>
                            <Statistic title="今日访问" value={1120} precision={0} valueStyle={{ color: '#a0d911' }} suffix="次"/>
                        </Col>
                        <Col span={4}>
                            <Statistic title="流量统计" value={1120678} precision={0} valueStyle={{ color: '#4096ff' }} suffix="M"/>
                        </Col>
                        <Col span={4}>
                            <Statistic title="系统消息" value={11} precision={0} valueStyle={{ color: '#f5222d' }} suffix="条"/>
                        </Col>
                    </Row>
                </StatisticCard>
                <QuickEntrance title="快捷入口" style={{}}>
                    <div className="quick-item">
                        <img src={UserImg} alt="" style={{width: 32, height: 32}}/>
                        <span>用户管理</span>
                    </div>
                    <div className="quick-item">
                        <img src={RoleImg} alt="" style={{width: 32, height: 32}}/>
                        <span>角色管理</span>
                    </div>
                    <div className="quick-item">
                        <img src={ResourceImg} alt="" style={{width: 32, height: 32}}/>
                        <span>资源管理</span>
                    </div>
                    <div className="quick-item">
                        <img src={ServerImg} alt="" style={{width: 32, height: 32}}/>
                        <span>服务器</span>
                    </div>
                    <div className="quick-item">
                        <img src={PersonImg} alt="" style={{width: 32, height: 32}}/>
                        <span>个人中心</span>
                    </div>
                </QuickEntrance>
                <Card title="访问统计" style={{borderRadius: 5}}>
                    <VisitCharts />
                </Card>
            </Col>
            <Col span={6}>
                <Card title="日历" style={{borderRadius: 5, marginBottom: 5}}>
                    <Calendar fullscreen={false} onPanelChange={onPanelChange} />
                </Card>
                <Card title="更新日志" style={{borderRadius: 5, marginBottom: 5}}>

                </Card>
            </Col>
        </Row>


    </>
}

const StatisticCard = styled(Card)`
    border-radius: 5px;
    margin-bottom: 5px;
    .ant-card-head {
        min-height: 40px;
    }
`

const QuickEntrance = styled(Card)`
    border-radius: 5px; 
    margin-bottom: 5px;
    .ant-card-head {
        min-height: 40px;
    }
    .ant-card-body {
        display: flex;
        flex-direction: row;

        .quick-item {
            display: flex;
            flex-direction: row;
            align-items: center;
            flex-wrap: nowrap;
            align-content: center;
            justify-content: center;
            margin-right: 30px;
            cursor: pointer;
            border: 1px solid #f1efef;
            padding: 10px;
            border-radius: 5px;

            img {
                width: 32px;
                height: 32px;
            }

            span {
                display: inline-block;
                height: 32px;
                line-height: 32px;
                margin-left: 5px;
                color: #2b85e4;
            }
        }
    }
`

export default HomePage;